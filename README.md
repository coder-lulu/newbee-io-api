# 新蜂资产管理平台 — 统一 I/O API 服务

为平台统一数据接入、发现及任务管理提供 HTTP 接口，将请求交给统一 I/O RPC 服务处理，并接入 Core RPC 的认证与权限能力。

仓库：[coder-lulu/newbee-io-api](https://github.com/coder-lulu/newbee-io-api) · [平台工作区](https://github.com/coder-lulu/newbee)

## 获取代码

推荐通过完整工作区开发，保留兄弟模块目录及本地 `replace` 依赖。以下命令使用 Bash；Go 工作区要求 Go 1.25.1 或更高版本。

```bash
git clone --recurse-submodules https://github.com/coder-lulu/newbee.git
cd newbee/unified-io/api
```

已有工作区执行 `git submodule update --init --recursive`。单独克隆模块时，需要自行补齐 `go.mod` 中的本地依赖路径。

## 目录导航

| 路径 | 用途 |
| --- | --- |
| `io.go` | 服务入口 |
| `desc/` | HTTP API 定义 |
| `internal/handler/`、`internal/logic/` | 路由处理与业务逻辑 |
| `internal/svc/`、`internal/middleware/` | 服务依赖与权限接入 |
| `etc/io.yaml.example` | 公开配置模板 |

## 配置与本地运行

首次配置时执行下列复制命令；已有配置不要覆盖：

```bash
cp etc/io.yaml.example etc/io.yaml
```

按环境修改 `DatabaseConf、RedisConf、CoreRpc、IoRpc、CROSConf 和 Middleware` 等配置项，以及监听地址。示例 HTTP 端口为 `9501`，以实际配置为准。入口使用 `conf.UseEnv()`，可为示例中的 `${...}` 占位符设置对应环境变量，也可在本地配置中填写值。模板中的内网地址不是可直接使用的公共服务。

先准备数据库、Redis 和配置引用的 RPC 服务，再启动本服务；API 应在对应 RPC 就绪后启动。真实配置和凭据不要提交。

```bash
go run . -f etc/io.yaml
```

## 构建与验证

在当前模块目录执行：

```bash
go build .
go test ./...
go vet ./...
```

测试中的集成用例需要对应基础服务。修改协议或 Ent Schema 后，应使用当前 Makefile 中对应生成目标并审查生成差异；不要直接编辑生成文件。上述命令是验证入口，不表示所有测试已通过。

## 文档

[接口定义](desc/all.api) · [OpenAPI 文件](io.json)

## 许可证与来源

本仓库采用 [Apache-2.0](LICENSE)。沿用现有服务框架的上游许可，保留文件中的原作者版权。第三方依赖遵循各自许可证，保留原有版权与许可声明。
