package svc

import (
	"context"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	rediswatcher "github.com/casbin/redis-watcher/v2"
	commoncasbin "github.com/coder-lulu/newbee-common/v2/casbin"
	commonadapter "github.com/coder-lulu/newbee-common/v2/casbin/adapter"
	"github.com/coder-lulu/newbee-common/v2/i18n"
	"github.com/coder-lulu/newbee-common/v2/middleware/audit"
	"github.com/coder-lulu/newbee-common/v2/middleware/integration"
	"github.com/coder-lulu/newbee-common/v2/middleware/keys"
	"github.com/coder-lulu/newbee-common/v2/orm/ent/hooks"
	"github.com/coder-lulu/newbee-core/rpc/coreclient"
	apicasbin "github.com/coder-lulu/newbee-io-api/internal/casbin"
	"github.com/coder-lulu/newbee-io-api/internal/config"
	i18n2 "github.com/coder-lulu/newbee-io-api/internal/i18n"
	"github.com/coder-lulu/newbee-io-rpc/ent"
	ioclient "github.com/coder-lulu/newbee-io-rpc/ioclient"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"strings"
)

type ServiceContext struct {
	Config         config.Config
	ContextManager *keys.ContextManager
	CoreRpc        coreclient.Core
	IoRpc          ioclient.Io
	DB             *ent.Client
	Trans          *i18n.Translator
	Casbin         *casbin.Enforcer
	// 🎯 统一中间件集成结果
	IntegrationResult *integration.Result
	// 统一中间件链
	ManagedMiddlewareChain []rest.Middleware
}

// RpcApiResourceProvider 通过RPC获取API资源名称的提供器
type RpcApiResourceProvider struct {
	coreRpc coreclient.Core
}

func NewRpcApiResourceProvider(coreRpc coreclient.Core) *RpcApiResourceProvider {
	return &RpcApiResourceProvider{coreRpc: coreRpc}
}

func (p *RpcApiResourceProvider) GetApiResourceName(ctx context.Context, method, path string) (string, error) {
	// 使用Core RPC服务获取API资源名称
	// 这里可以实现更复杂的资源名称解析逻辑
	return "io:" + method + ":" + path, nil
}

// GetCoreRpcClient 实现audit.AuditSvcProvider接口，为审计插件提供Core RPC客户端
func (svc *ServiceContext) GetCoreRpcClient() interface{} {
	return svc.CoreRpc
}

// GetCasbinEnforcer 实现permission.EnforcerProvider接口，为RBAC插件提供Casbin执行器
func (svc *ServiceContext) GetCasbinEnforcer() interface{} {
	return svc.Casbin
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 安全默认值，避免空环境变量导致连接失败
	if strings.TrimSpace(c.RedisConf.Host) == "" {
		c.RedisConf.Host = "127.0.0.1:6380"
	}
	// ===========================================
	// 🎉 统一中间件框架集成 - IO服务
	// ===========================================

	// 1. 初始化基础服务
	rds := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{c.RedisConf.Host},
		Password: c.RedisConf.Pass,
		DB:       c.RedisConf.Db,
	})

	trans := i18n.NewTranslator(c.I18nConf, i18n2.LocaleFS)

	// 2. 初始化数据库连接
	db := ent.NewClient(
		ent.Log(func(v ...interface{}) {}),
		ent.Driver(c.DatabaseConf.NewNoCacheDriver()),
		ent.Debug(),
	)

	// Match the RPC service exclusions for shared CMDB reference tables.
	hooks.AddExcludedTable("cmdb_asset_types")
	hooks.AddExcludedTable("cmdb_templates")
	hooks.AddExcludedTable("cmdb_attribute_definitions")
	if err := hooks.QuickSetup(db); err != nil {
		panic("IO API tenant hook initialization failed: " + err.Error())
	}

	// 3. 初始化RPC客户端 - 使用SystemContext拦截器支持系统级操作
	coreRpcClient, err := zrpc.NewClient(c.CoreRpc, zrpc.WithUnaryClientInterceptor(hooks.SystemContextClientInterceptor()))
	if err != nil {
		panic("Failed to create Core RPC client: " + err.Error())
	}
	coreRpc := coreclient.NewCore(coreRpcClient)

	// Io RPC 客户端（用于InitDatabase等动作）
	var ioRpc ioclient.Io
	if cli, err := zrpc.NewClient(c.IoRpc, zrpc.WithUnaryClientInterceptor(hooks.SystemContextClientInterceptor())); err == nil {
		ioRpc = ioclient.NewIo(cli)
	}

	// 3. 获取JWT密钥 - 使用新的Middleware配置
	jwtSecret := ""
	if c.Middleware.Auth != nil && c.Middleware.Auth.AccessSecret != "" {
		jwtSecret = c.Middleware.Auth.AccessSecret
	}

	// 4. 🔥 初始化Casbin - 使用EntAdapter通过RPC查询规则
	logx.Info("Initializing Casbin with EntAdapter and RPC querier for all tenants")

	// 4.1 创建RPC查询器
	rpcQuerier := apicasbin.NewRpcCasbinRuleQuerier(coreRpc)

	// 4.2 创建EntAdapter
	// 🔥 使用SystemContext绕过租户隔离Hook，加载所有租户的Casbin规则
	systemCtx := hooks.NewSystemContext(context.Background())
	adapter := commonadapter.NewEntAdapter(rpcQuerier, systemCtx)

	// 4.3 创建Casbin模型
	modelText := commoncasbin.GetDefaultRBACWithDomainsModel()
	m, err := model.NewModelFromString(modelText)
	if err != nil {
		logx.Errorf("Failed to create Casbin model: %v", err)
		panic("Casbin model creation failed: " + err.Error())
	}

	// 4.4 创建Enforcer
	cbn, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		logx.Errorf("Failed to create Casbin enforcer: %v", err)
		panic("Casbin enforcer creation failed: " + err.Error())
	}

	// 4.5 加载策略
	err = cbn.LoadPolicy()
	if err != nil {
		logx.Errorf("Failed to load Casbin policy: %v", err)
		panic("Casbin policy loading failed: " + err.Error())
	}

	// 4.6 添加Redis Watcher（用于策略同步）
	w := c.CasbinConf.MustNewOriginalRedisWatcher(c.RedisConf, func(data string) {
		rediswatcher.DefaultUpdateCallback(cbn)(data)
	})
	err = cbn.SetWatcher(w)
	if err != nil {
		logx.Errorf("Failed to set Casbin watcher: %v", err)
		panic("Casbin watcher setup failed: " + err.Error())
	}

	logx.Info("✅ IO API: Casbin initialized with EntAdapter, connected to sys_casbin_rules table via RPC")

	// 5. 创建服务上下文实例（需要先创建以便传递给审计插件）
	svcCtx := &ServiceContext{
		Config:  c,
		CoreRpc: coreRpc,
		IoRpc:   ioRpc,
		DB:      db,
		Trans:   trans,
		Casbin:  cbn,
	}

	// 6. 🎯 使用统一中间件集成API，创建审计写入器
	ioAuditWriter := audit.NewBuiltinAuditWriter(svcCtx)

	result, err := integration.Setup(&integration.Config{
		Redis:               rds,
		JWTSecret:           jwtSecret,
		Mode:                integration.Production,
		ApiResourceProvider: NewRpcApiResourceProvider(coreRpc),
		// 使用common包的BuiltinAuditWriter通过Core RPC写入审计日志
		AuditWriter:            ioAuditWriter,
		TenantInfoProvider:     NewRpcTenantInfoProvider(coreRpc),
		RbacProvider:           svcCtx,     // 🔥 提供 Casbin enforcer（通过GetCasbinEnforcer接口）
		DataPermCasbinProvider: rpcQuerier, // 🔥 提供数据权限Casbin查询器（UnifiedDataPermPlugin必需）
		Middleware:             &c.Middleware,
	})
	if err != nil {
		panic("IO统一中间件集成失败: " + err.Error())
	}

	// 7. 完善服务上下文
	svcCtx.ContextManager = result.ContextManager
	svcCtx.IntegrationResult = result
	svcCtx.ManagedMiddlewareChain = result.Middlewares

	return svcCtx
}
