package config

import (
	"github.com/coder-lulu/newbee-common/v2/config"
	"github.com/coder-lulu/newbee-common/v2/i18n"
	"github.com/coder-lulu/newbee-common/v2/middleware/framework"
	"github.com/coder-lulu/newbee-common/v2/plugins/casbin"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
    rest.RestConf
    // Auth field removed - using unified middleware framework instead
    Middleware   framework.UnifiedConfig `json:",optional"`
    RedisConf    config.RedisConf
    CoreRpc      zrpc.RpcClientConf
    // Io RPC，用于调用RPC层初始化等动作
    IoRpc        zrpc.RpcClientConf `json:",optional"`
    DatabaseConf config.DatabaseConf
    CasbinConf   casbin.CasbinConf
    I18nConf     i18n.Conf
    CROSConf     config.CROSConf
}
