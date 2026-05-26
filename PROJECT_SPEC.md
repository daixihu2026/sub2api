# Sub2API 项目详细规格说明文档

## 1. 项目概述

Sub2API 是一个 **AI API 网关平台**，用于分发和管理 AI 产品订阅的 API 配额。它本质上是一个多租户、多平台的 API 中转/代理系统，接受客户端以多种 API 格式（Anthropic Messages、OpenAI Responses/Chat Completions、Gemini GenerateContent）发起的请求，并将其路由到上游 AI 服务提供商的账号。

**核心技术栈：** Go 后端 + Vue 3 前端 + PostgreSQL + Redis

**官方域名：** `sub2api.org` 与 `pincc.ai`

---

## 2. 技术架构总览

```
┌─────────────────────────────────────────────────────────────┐
│                        客户端 (Client)                       │
│    Claude Code CLI / Codex CLI / Gemini CLI / SDK / Web     │
└──────────────────────┬──────────────────────────────────────┘
                       │ Anthropic API / OpenAI API / Gemini API
                       ▼
┌─────────────────────────────────────────────────────────────┐
│                    Sub2API 网关 (Gateway)                     │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  API Key 鉴权 → 用户并发控制 → 负载均衡 → 账号选择   │   │
│  │  → 请求改写 → 上游转发 → 故障转移 → 用量计费         │   │
│  └──────────────────────────────────────────────────────┘   │
└──────────────────────┬──────────────────────────────────────┘
                       │ 多种上游账号
                       ▼
┌──────────┬──────────┬──────────┬──────────────────────────────┐
│ Anthropic│ OpenAI   │ Gemini   │ Antigravity (自定义)         │
│ (OAuth/  │ (API Key/│ (OAuth/  │                              │
│ SetupToken)│ OAuth) │ API Key) │                              │
└──────────┴──────────┴──────────┴──────────────────────────────┘
```

### 运行时模式

- **standard 模式：** 完整功能，包括计费、配额管理
- **simple 模式：** 简化模式，禁用计费和配额检查，适合个人/小团队自用
- **backend 模式：** 仅管理员可访问，普通用户无法登录/注册

---

## 3. 技术栈详情

| 层级 | 技术 | 版本 |
|------|------|------|
| **后端语言** | Go | 1.25.7 |
| **HTTP 框架** | Gin | - |
| **ORM** | Ent (entgo.io) | - |
| **数据库** | PostgreSQL | 15+ |
| **缓存** | Redis | 7+ |
| **依赖注入** | Google Wire | 编译时 |
| **配置管理** | Viper (YAML) | - |
| **前端框架** | Vue | 3.4+ |
| **构建工具** | Vite | - |
| **CSS 框架** | Tailwind CSS | - |
| **状态管理** | Pinia | - |
| **路由** | Vue Router | 4.x |
| **HTTP 客户端** | Axios | - |
| **包管理** | pnpm | - |
| **容器化** | Docker | - |
| **CI/CD** | GitHub Actions + GoReleaser | - |

---

## 4. 目录结构

```
sub2api/
├── backend/                    # Go 后端
│   ├── cmd/                    # 入口程序
│   │   ├── server/             # 主服务 (main.go, wire.go)
│   │   └── jwtgen/             # JWT 生成工具
│   ├── ent/                    # Ent ORM 生成代码 + Schema 定义
│   │   └── schema/             # 数据库表定义
│   ├── internal/
│   │   ├── config/             # 配置系统 (Viper)
│   │   ├── domain/             # 领域常量、消息分发
│   │   ├── handler/            # HTTP 处理器 (Controller 层)
│   │   │   ├── admin/          # 管理后台处理器
│   │   │   └── dto/            # 数据传输对象
│   │   ├── integration/        # 外部集成 (iframely 等)
│   │   ├── middleware/         # 中间件
│   │   ├── model/              # 领域模型
│   │   ├── payment/            # 支付系统
│   │   │   └── provider/       # 支付提供商实现
│   │   ├── pkg/                # 通用工具包
│   │   ├── repository/         # 数据访问层
│   │   ├── server/             # HTTP 路由注册
│   │   │   └── routes/         # 各模块路由定义
│   │   ├── service/            # 业务逻辑层 (核心)
│   │   ├── setup/              # 系统初始化
│   │   ├── testutil/           # 测试工具
│   │   ├── util/               # 工具函数
│   │   └── web/                # 前端相关
│   ├── migrations/             # 数据库迁移 SQL
│   └── resources/              # 静态资源
├── frontend/                   # Vue 3 前端
│   ├── src/
│   │   ├── api/                # API 请求层 (Axios)
│   │   │   └── admin/          # 管理后台 API
│   │   ├── components/         # 公共组件
│   │   ├── composables/        # Vue Composables
│   │   ├── constants/          # 常量
│   │   ├── i18n/               # 国际化
│   │   ├── router/             # Vue Router 路由
│   │   ├── stores/             # Pinia 状态管理
│   │   ├── styles/             # 样式
│   │   ├── types/              # TypeScript 类型定义
│   │   ├── utils/              # 工具函数
│   │   └── views/              # 页面视图
│   │       ├── admin/          # 管理后台页面
│   │       ├── auth/           # 认证页面
│   │       ├── public/         # 公开页面
│   │       ├── setup/          # 系统设置向导
│   │       └── user/           # 用户页面
│   └── public/                 # 静态资源
├── deploy/                     # 部署配置 (Docker Compose)
├── docs/                       # 文档
├── tools/                      # 工具脚本
├── assets/                     # 资源文件
├── Dockerfile                  # 主 Docker 镜像
├── Makefile                    # 构建命令
└── config.yaml                 # 配置文件示例
```

---

## 5. 后端架构详解

### 5.1 入口与启动流程

**主入口：** `backend/cmd/server/main.go`

启动流程分为以下阶段：

1. 解析命令行参数：`--setup` (CLI安装向导)、`--version` (版本信息)
2. 检查是否需要初始化设置 (`setup.NeedsSetup()`)
   - Docker 部署：自动从环境变量完成设置
   - 手动部署：启动 Web 设置向导 (Gin 最小服务器)
3. 执行 `runMainServer()`：
   - 加载配置 → 初始化结构化日志
   - 通过 **Google Wire** 编译时依赖注入构建完整应用对象图
   - 启动 HTTP 服务器 + 优雅关闭 (SIGINT/SIGTERM, 5秒超时)
   - 清理时并行停止 25 个后台服务，关闭 Redis 和数据库连接

**依赖注入层级：**
```
config → DB/Redis → Repository → Service → Middleware → Handler → Router → HTTP Server
```

### 5.2 配置系统

**文件：** `backend/internal/config/config.go`

使用 **Viper** 从 `config.yaml` 加载。配置路径优先级：`DATA_DIR` 环境变量 > 当前目录 > `./config/`。

**主要配置域（共 40+ 个配置节）：**

| 配置域 | 说明 |
|--------|------|
| `server` | 服务地址、模式、超时、CORS、H2C (HTTP/2 Cleartext) |
| `database` | PostgreSQL 连接参数 |
| `redis` | Redis 连接参数、连接池、TLS |
| `jwt` | JWT 密钥、过期时间、刷新策略 |
| `security` | CSP 策略、帧来源、安全头 |
| `billing` | 计费模式、Pricing 配置 |
| `payment` | 支付开关、支付提供商密钥 |
| `gateway` | 网关核心配置 (超时、体大小限制、连接池隔离、TLS 指纹等) |
| `rate_limit` | 速率限制策略 |
| `concurrency` | 用户级和账号级并发限制 |
| `gemini` | Gemini 平台专用配置 (CLI 会话管理等) |
| `totp` | 双因素认证配置 |
| `oidc_connect` / `linuxdo_connect` / `wechat_connect` / `github_oauth` / `google_oauth` | 各 OAuth 登录配置 |
| `subscription_cache` / `subscription_maintenance` | 订阅缓存和维护 |
| `dashboard_cache` / `dashboard_aggregation` | 仪表盘缓存和聚合 |
| `idempotency` | 幂等性请求去重 |
| `usage_cleanup` | 用量记录清理策略 |
| `token_refresh` | Token 刷新调度 |

### 5.3 路由结构

**全局中间件链：**
1. `RequestLogger` → 注入请求级日志 (request_id)
2. `Logger` → Gin 风格请求日志
3. `CORS` → 可配置跨域
4. `SecurityHeaders` → CSP、X-Frame-Options、X-Content-Type-Options
5. `Recovery` → 自定义 Panic 恢复

**路由分组：**

#### 网关路由（API Key 鉴权）
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/v1/messages` | Anthropic Messages API (按分组平台自动路由) |
| POST | `/v1/messages/count_tokens` | Token 计数 |
| GET  | `/v1/models` | 模型列表 |
| POST | `/v1/responses`, `/v1/responses/*subpath` | OpenAI Responses API |
| GET  | `/v1/responses` | Responses WebSocket 升级 |
| POST | `/v1/chat/completions` | OpenAI Chat Completions |
| POST | `/v1/images/generations`, `/v1/images/edits` | 图片生成 |
| GET  | `/v1beta/models` | Gemini 模型列表 |
| POST | `/v1beta/models/*modelAction` | Gemini generateContent/streamGenerateContent |
| *    | `/antigravity/v1/*` | Antigravity 专有路由 (Anthropic 格式) |
| *    | `/antigravity/v1beta/*` | Antigravity 专有路由 (Gemini 格式) |
| POST | `/backend-api/codex/*` | Codex CLI 直接路由 |
| POST | `/responses`, `/chat/completions`, `/images/*` | 无版本前缀别名 |

#### 用户 API (`/api/v1`)

| 路径 | 鉴权 | 说明 |
|------|------|------|
| `/auth/*` | 无/速率限制 | 注册、登录、2FA、刷新Token、OAuth、密码重置 |
| `/settings/public` | 无 | 公开站点设置 |
| `/user/profile` | JWT | 用户信息 |
| `/user/aff` | JWT | 分销/推荐信息 |
| `/keys` | JWT | API Key 管理 CRUD |
| `/groups/available`, `/groups/rates` | JWT | 分组浏览 |
| `/channels/available` | JWT | 通道浏览 |
| `/usage` | JWT | 用量记录与统计 |
| `/subscriptions` | JWT | 订阅管理 |
| `/redeem` | JWT | 兑换码 |
| `/announcements` | JWT | 公告 |
| `/channel-monitors` | JWT | 通道状态监控 |

#### 支付路由 (`/api/v1/payment`)
| 路径 | 鉴权 | 说明 |
|------|------|------|
| `/config`, `/plans`, `/checkout-info` | JWT | 支付配置/套餐 |
| `/orders` | JWT | 创建/查询订单 |
| `/channels` | JWT | 可用支付渠道 |
| `/webhook/easypay`, `/webhook/alipay`, `/webhook/wxpay`, `/webhook/stripe`, `/webhook/airwallex` | 无 | 支付回调 |
| `/public/orders/verify`, `/public/orders/resolve` | 无 | 公开订单验证 |

#### 管理后台 (`/api/v1/admin`)
覆盖以下功能模块：仪表盘、用户管理、分组管理、账号管理、公告、代理、兑换码、优惠码、支付管理、用量记录、系统设置、运维监控、订阅管理、分销管理、风控、通道监控、数据管理、备份、TLS指纹、错误透传规则、用户属性等 30+ 个管理域。

### 5.4 数据库 Schema（核心实体）

#### User (`users`) - 软删除
| 字段 | 类型 | 说明 |
|------|------|------|
| email | string(255) | 邮箱（部分唯一） |
| password_hash | string(255) | 密码哈希 |
| role | string(20) | 角色 (user/admin) |
| balance | decimal(20,8) | 账户余额 (USD) |
| concurrency | int | API 并发数 (默认5) |
| status | string(20) | 状态 (active/disabled) |
| totp_enabled | bool | TOTP 2FA 开关 |
| signup_source | string | 注册来源 (email/linuxdo/wechat/oidc/github/google) |
| total_recharged | decimal(20,8) | 累计充值金额 |
| rpm_limit | int | 每分钟请求限制 |
| aff_code | string | 分销邀请码 |
| aff_quota | decimal | 可提现佣金 |
| aff_frozen_quota | decimal | 冻结中佣金 |
| aff_history_quota | decimal | 历史佣金总额 |
| aff_count | int | 邀请人数 |
| aff_rebate_rate_percent | decimal | 个人返佣比例 |

#### APIKey (`api_keys`) - 软删除
| 字段 | 类型 | 说明 |
|------|------|------|
| user_id | int64 | 所属用户 |
| key | string(128) | API Key 哈希 (唯一) |
| name | string(100) | 名称 |
| group_id | int64 | 所属分组 |
| status | string(20) | 状态 (active/disabled) |
| ip_whitelist | JSONB | IP 白名单 |
| ip_blacklist | JSONB | IP 黑名单 |
| quota | decimal | 配额上限 |
| quota_used | decimal | 已用配额 |
| expires_at | timestamptz | 过期时间 |
| rate_limit_5h/1d/7d | decimal | 各时间窗口速率限制 |
| usage_5h/1d/7d | decimal | 各时间窗口累计用量 |

#### Account (`accounts`) - 软删除 (AI上游账号)
| 字段 | 类型 | 说明 |
|------|------|------|
| name | string(100) | 名称 |
| platform | string(50) | 平台 (anthropic/openai/gemini/antigravity) |
| type | string(20) | 类型 (api_key/oauth/cookie/bedrock等) |
| credentials | JSONB | 凭据加密存储 |
| extra | JSONB | 额外配置 |
| proxy_id | int64 | 关联代理 |
| concurrency | int | 账号并发数 (默认3) |
| priority | int | 调度优先级 (默认50, 越低越高) |
| load_factor | int | 负载因子覆盖 |
| rate_multiplier | decimal | 成本倍率 |
| status | string(20) | 状态 |
| overload_until | timestamptz | 过载冷却截止 |
| temp_unschedulable_until | timestamptz | 临时不可调度截止 |
| session_window_start/end | timestamptz | 5小时会话窗口 |
| session_window_status | string(20) | 窗口状态 |

#### Group (`groups`) - 软删除 (API分组/套餐)
| 字段 | 类型 | 说明 |
|------|------|------|
| name | string(100) | 分组名 |
| platform | string(50) | 平台类型 |
| rate_multiplier | decimal | 价格倍率 |
| is_exclusive | bool | 是否专属分组 |
| daily_limit_usd / weekly_limit_usd / monthly_limit_usd | decimal | 用量上限 |
| claude_code_only | bool | 仅限 Claude Code |
| fallback_group_id | int64 | 故障转移目标分组 |
| model_routing | JSONB | 模型路由配置 |
| supported_model_scopes | JSONB | 支持的模型范围 |
| allow_image_generation | bool | 允许图片生成 |
| require_oauth_only | bool | 仅限 OAuth 账号 |
| rpm_limit | int | 每分钟请求限制 |

#### PaymentOrder (`payment_orders`) - 硬删除
记录充值/订阅订单的完整生命周期：金额、支付类型、交易号、状态流转 (PENDING→PAID→RECHARGING→COMPLETED)、退款信息。

#### RedeemCode (`redeem_codes`) - 硬删除
兑换码：支持 `balance` (余额)、`concurrency` (并发)、`subscription` (订阅)、`invitation` (邀请) 四种类型。

#### SubscriptionPlan (`subscription_plans`) - 硬删除
订阅套餐：价格、有效期、关联分组、特性描述。

#### UserSubscription (`user_subscriptions`) - 软删除
用户订阅：激活/过期时间、各时间窗口用量、日/周/月使用限额追踪。

#### UsageLog (`usage_logs`) - 追加不删
用量日志：Token 计数 (输入/输出/缓存创建/缓存读取)、成本、计费类型、模型、响应时间、User-Agent、IP 等。

其他核心实体：
- **Proxy** - HTTP/HTTPS/SOCKS5 代理节点
- **Setting** - 键值对系统设置
- **Announcement** - 系统公告 (支持定向投放)
- **PromoCode** - 注册优惠码 (赠送余额)
- **AuthIdentity** - OAuth 身份绑定
- **ChannelMonitor** - 通道健康监控
- **IdempotencyRecord** - 幂等请求去重
- **ErrorPassthroughRule** - 错误透传规则
- **TLSFingerprintProfile** - TLS 指纹模拟配置

---

## 6. 核心业务逻辑

### 6.1 API 请求代理流程

以 `/v1/messages` 请求（Anthropic 协议）为例，完整处理管线如下：

```
Client Request
  │
  ├─ 1. ApiKeyAuth 中间件 ─── 提取 API Key、验证、加载用户/分组/订阅
  ├─ 2. 读取请求体 ────────── 解析 JSON、验证格式
  ├─ 3. 客户端检测 ────────── 识别 Claude Code CLI / Codex CLI
  ├─ 4. 内容审核 ──────────── 检查请求内容
  ├─ 5. 用户并发控制 ──────── 获取用户级并发槽位 (等待时发送 SSE ping)
  ├─ 6. 计费资格重检 ──────── 再次检查余额/订阅状态
  ├─ 7. 粘性会话查询 ──────── 查找缓存中的会话→账号绑定
  ├─ 8. 账号选择 (负载均衡) ── 多层调度算法选择最优账号
  ├─ 9. 请求拦截检测 ──────── 识别预热/探测请求并返回模拟响应
  ├─ 10. 账号并发控制 ─────── 获取账号级并发槽位
  ├─ 11. 用户消息序列化 ───── 对OAuth账号强制执行消息顺序
  ├─ 12. 请求改写 ─────────── 系统提示词注入、模型名映射、缓存控制
  ├─ 13. Token 获取 ───────── 获取 OAuth Access Token 或 API Key
  ├─ 14. 代理/TLS指纹 ─────── 配置 HTTP/SOCKS5 代理和 TLS JA3 指纹
  ├─ 15. 上游转发 ─────────── HTTP 请求发送 (带重试循环)
  ├─ 16. 故障转移 ─────────── 失败时切换账号重试
  ├─ 17. 粘性会话绑定 ─────── 成功后绑定会话→账号
  └─ 18. 用量记录 ─────────── 异步提交用量到 Worker Pool
```

### 6.2 支持的 AI 平台与 API 格式

| 平台 | 常量值 | 支持的 API 格式 | 账号类型 |
|------|--------|----------------|---------|
| Anthropic | `anthropic` | `/v1/messages`, `/v1/models` | OAuth, SetupToken, API Key |
| OpenAI | `openai` | `/v1/responses`, `/v1/chat/completions`, `/v1/images/*` | API Key, OAuth |
| Gemini | `gemini` | `/v1beta/models/*:generateContent` | OAuth, API Key, Service Account |
| Antigravity | `antigravity` | Anthropic 格式 + Gemini 格式 | 自定义 (类似 OAuth) |

### 6.3 Token 获取机制

上游 AI 服务账号的凭据管理方式因账号类型而异：

**OAuth 类型账号 (Anthropic/Gemini/Antigravity)：**
1. 管理员在后台通过 OAuth 流程授权：系统生成授权 URL → 管理员在浏览器中授权 → 系统交换授权码获取 Access Token + Refresh Token
2. Access Token 加密存储在 `accounts.credentials` JSONB 字段中
3. `TokenRefreshService` 后台服务定期检查即将过期的 Token 并使用 Refresh Token 自动续期
4. 如果 Token 刷新失败，账号会被标记为 `rate_limited` 或 `overloaded` 状态
5. 对于 Antigravity 平台，使用 `SetupToken` 作为初始凭据交换 API Key

**API Key 类型账号 (OpenAI/部分 Gemini)：**
1. 管理员直接输入 API Key
2. API Key 加密存储在 `accounts.credentials` 中
3. 调用时直接从存储中取出使用 (无需刷新)

**凭据安全：**
- 所有凭据在数据库中使用 AES-256-GCM 加密存储
- 加密密钥存储在 `security_secrets` 表中
- API Key 哈希后存储 (`api_keys.key` 为哈希值)

### 6.4 负载均衡与账号选择算法

**文件：** `backend/internal/service/gateway_service.go`

账号选择采用**多层调度算法**，优先级从高到低：

```
Layer 0: 预检查
  ├── 通道定价检查
  ├── 粘性会话预取
  └── 简单模式回退

Layer 1: 模型路由 (仅 Anthropic 分组开启 model_routing)
  ├── 按请求模型筛选路由账号
  ├── 过滤：可调度性、平台兼容、模型支持、配额、窗口成本、RPM
  └── 负载感知选择：LoadRate < 100% → 按 Priority > LoadRate > LRU 排序

Layer 1.5: 粘性会话
  ├── 检查已绑定的会话→账号
  ├── 验证门控条件（RPM 豁免策略）
  └── 优先使用绑定账号

Layer 2: 负载感知选择
  ├── 筛选所有可用候选账号
  ├── 批量获取负载信息
  ├── 优先级筛选 → 负载率筛选 → LRU 排序
  └── 对于 Gemini 优先选择 OAuth 账号

Layer 3: 回退等待
  └── 所有账号槽位已满 → 返回等待计划
```

**调度配置默认值：**

| 参数 | 默认值 |
|------|--------|
| StickySessionMaxWaiting | 3 |
| StickySessionWaitTimeout | 45秒 |
| FallbackWaitTimeout | 30秒 |
| FallbackMaxWaiting | 100 |
| LoadBatchEnabled | true |

### 6.5 故障转移 (Failover)

当上游请求失败时，系统会尝试故障转移：

1. **相同账号重试：** 如果错误标记为 `RetryableOnSameAccount`，最多重试 3 次（间隔 500ms）
2. **切换账号重试：** 同账号重试耗尽后标记该账号临时不可调度，切换到其他可用账号（最多 `MaxSwitches` 次，Anthropic 默认 10，Gemini/OpenAI 默认 3）
3. **AntiGravity 平台延迟：** 账号切换间增加线性延迟 `(switchCount-1) * 1s`
4. **错误映射：** 上游错误码映射为网关标准响应
   - 401/403 → 502 (upstream_error)
   - 429 → 429 (rate_limit_error)
   - 529 → 503 (overloaded_error)
   - 500/502/503/504 → 502 (upstream_error)
5. **SSE 流式错误：** 已开始的流式请求使用 SSE error 事件而非 JSON 响应

### 6.6 计费系统

**文件：** `backend/internal/service/billing_service.go`

三种计费模式：

1. **Token 计费（默认）：**
   - 按 token 数量计费：input_tokens、output_tokens、cache_creation_tokens、cache_read_tokens
   - 支持缓存分类计费：5分钟缓存 (cache_creation_5m)、1小时缓存 (cache_creation_1h)
   - 支持分层定价：按 token 使用量区间不同价格
   - 支持服务等级倍率：Priority 2x、Flex 0.5x
   - 支持长上下文倍率（如 GPT-5.4: >272K tokens 输入2x/输出1.5x）
   - 图片输出 Token 独立计费

2. **按次计费：** 固定每请求价格，支持按图片尺寸分层（1K/2K/4K）

3. **图片计费：** 按图片数量和尺寸计费

**价格来源优先级：**
1. `PricingService` 动态定价（LiteLLM 兼容格式，USD/token）
2. 代码内硬编码回退价格（覆盖 Claude、GPT、Gemini、Codex 等约 20 个模型系列）

**计费扣费项：**
- `BalanceCost` - 余额扣费
- `SubscriptionCost` - 订阅用量扣费
- `APIKeyQuotaCost` - API Key 配额扣费
- `APIKeyRateLimitCost` - API Key 速率限制统计
- `AccountQuotaCost` - 账号配额扣费

**去重机制：** 使用 SHA-256 对计费命令所有字段进行指纹去重，防止重复扣费。

---

## 7. 支付系统

### 7.1 支付提供商

支付采用**策略模式 + 注册表模式**，支持以下提供商：

| 提供商 | Provider Key | 支持类型 | 说明 |
|--------|-------------|---------|------|
| EasyPay | `easypay` | alipay, wxpay | 易支付聚合支付，MD5签名，支持扫码和跳转两种模式 |
| 支付宝官方 | `alipay` | alipay_direct | smartwalle/alipay/v3 SDK，支持桌面扫码和手机WAP |
| 微信官方 | `wxpay` | wxpay_direct | wechatpay-apiv3 SDK，支持扫码/H5/JSAPI三种模式 |
| Stripe | `stripe` | stripe, card, alipay, wechat_pay, link | stripe-go/v85 SDK，PaymentIntents API |
| Airwallex | `airwallex` | airwallex | REST API + OAuth2 Bearer Token (HS256)，支持多币种 |

**提供商接口：**
```go
type Provider interface {
    Name() string
    ProviderKey() string
    SupportedTypes() []Type
    CreatePayment(ctx, order, instance, req) (*CreatePaymentResult, error)
    QueryOrder(ctx, order, instance) (*QueryOrderResult, error)
    VerifyNotification(ctx, body, instance) (*VerifyResult, error)
    Refund(ctx, order, instance, amount, reason) (*RefundResult, error)
}
```

**提供商负载均衡：**
- Round-robin（默认）：原子计数器轮转
- Least-amount：选择当日订单金额最低的实例
- 过滤条件：检查单笔限额、单笔上限、每日限额

### 7.2 订单生命周期

```
PENDING ──支付成功──▶ PAID ──处理回调──▶ RECHARGING ──充值完成──▶ COMPLETED
   │                    │                      │
   │                    │                      └──充值失败──▶ FAILED
   │                    │
   └──超时──▶ EXPIRED   └──主动取消──▶ CANCELLED

COMPLETED ──申请退款──▶ REFUND_REQUESTED ──处理中──▶ REFUNDING ──完成──▶ REFUNDED
                                                               └──部分──▶ PARTIALLY_REFUNDED
```

**订单类型：**
- `balance` - 余额充值
- `subscription` - 订阅购买

### 7.3 Token/额度购买流程

```
用户前端 (PaymentView)
  │
  ├─ 选择充值金额/套餐
  ├─ 选择支付方式 (支付宝/微信/Stripe/Airwallex)
  ├─ 如有优惠码输入
  │
  ▼
POST /api/v1/payment/orders  →  PaymentHandler.CreateOrder()
  │
  ├─ 验证：支付开关、金额范围、日限额、待支付订单数
  ├─ 计算手续费 (fee.go)
  ├─ 负载均衡选择提供商实例
  ├─ 创建 PaymentOrder (PENDING 状态)
  ├─ 调用 Provider.CreatePayment()
  │   ├─ EasyPay: 返回扫码/跳转页面 URL
  │   ├─ 支付宝: 返回 QR Code URL
  │   ├─ 微信: 返回扫码URL / H5 URL / JSAPI参数
  │   ├─ Stripe: 返回 clientSecret
  │   └─ Airwallex: 返回 clientSecret
  └─ 返回支付信息给前端
  │
  ▼
用户在支付页面完成支付 → 支付网关回调
  │
  ▼
POST /api/v1/payment/webhook/{provider}  →  PaymentWebhookHandler
  │
  ├─ 提取 out_trade_no → 匹配订单
  ├─ Provider.VerifyNotification() → 签名验证
  ├─ PaymentService.HandlePaymentNotification()
  │   ├─ 验证金额容差 (0.01 CNY 以内)
  │   ├─ 订单状态 PENDING → PAID
  │   ├─ executeFulfillment():
  │   │   ├─ Balance订单: 创建内部兑换码 → 自动兑换 → 用户余额增加
  │   │   └─ Subscription订单: 分配/延长用户订阅
  │   └─ 触发分销佣金 (applyAffiliateRebateForOrder)
  └─ 返回成功响应
```

**手续费计算：**
- 充值金额 + 手续费比例 (fee_rate) = 实际支付金额
- 按照币种最小单位向上取整 (如 CNY 到分)

**余额充值倍率：**
- 配置项 `BALANCE_RECHARGE_MULTIPLIER`：支付 $X，到账 `$X * multiplier`
- 默认 1.0（1:1）

### 7.4 订阅模型

**订阅套餐 (SubscriptionPlan)：**
- 关联到 API Group
- 属性：价格、原价、有效期 (天)、特性列表
- 可标记 `for_sale` 控制是否上架

**用户订阅 (UserSubscription)：**
- 关联用户 + 分组
- 日/周/月用量限额追踪
- 状态：active / expired / suspended

**订阅购买流程：**
1. 用户在套餐页选择套餐
2. 通过支付系统完成支付
3. 回调处理调用 `SubscriptionService.AssignOrExtendSubscription()`
4. 自动创建或延长用户订阅
5. 订阅期内 API 调用将消耗订阅用量限额

### 7.5 兑换码系统

**兑换码类型：**

| 类型 | 说明 | 效果 |
|------|------|------|
| `balance` | 余额兑换码 | 增加/减少用户余额（正数为增加，负数为退款） |
| `concurrency` | 并发兑换码 | 调整用户 API 并发限制 |
| `subscription` | 订阅兑换码 | 赠送/延长指定分组的订阅时间 |
| `invitation` | 邀请兑换码 | 用于分销邀请追踪（不可由终端用户兑换） |

**兑换码格式：** 128位随机十六进制，格式 `XXXX-XXXX-XXXX-XXXX`（大写，短横线分隔）

**兑换流程：**
1. 速率限制：每用户每小时最多 20 次失败尝试
2. 分布式锁：每个兑换码 10 秒 TTL
3. 原子操作：乐观锁 (`WHERE status='unused'`)
4. 余额充值兑换码会触发分销佣金计算

---

## 8. 分销/推荐系统

### 8.1 是否支持分销

**是的，Sub2API 内置了完整的分销/推荐（Affiliate）系统。**

### 8.2 分销架构

```
┌─────────────────────────────────────────────┐
│                 系统管理员                     │
│  设置全局返佣比例 / 管理分销记录              │
└──────────────────┬──────────────────────────┘
                   │
      ┌────────────┴────────────┐
      ▼                         ▼
┌──────────┐            ┌──────────┐
│  用户 A   │            │  用户 B   │
│ (邀请人)  │──邀请码──▶│ (被邀请人) │
│          │            │          │
│ 邀请码:  │            │ 注册时填写│
│ ABC123   │            │ aff=ABC123│
└──────────┘            └────┬─────┘
      ▲                      │
      │         充值 $100     │
      │                      ▼
      │              ┌──────────────┐
      │              │  充值到账 $100 │
      ├──返佣 $X─────┤  触发返佣计算  │
      │              └──────────────┘
      │
┌─────┴────────┐
│  用户 A 获得返佣│
│  aff_quota += X │
│  (冻结状态)     │
└────────────────┘
      │ 冻结期过后自动解冻
      ▼
┌──────────────┐
│ 用户 A 提现   │
│ 转到余额      │
│ balance += X  │
└──────────────┘
```

### 8.3 分销机制详解

**核心文件：** `backend/internal/service/affiliate_service.go`

**1. 邀请码生成：**
- 每个用户注册时自动生成分销邀请码
- 默认 12 字符随机码
- 管理员可为用户设置 4-32 字符的自定义码 (A-Z, 0-9, _, -)

**2. 邀请绑定：**
- 新用户注册时在 URL 参数中携带 `aff=<邀请码>`
- 也支持 `invitation_code`（邀请兑换码）绑定
- 系统调用 `AffiliateService.BindInviterByCode()` 建立绑定关系

**3. 返佣触发时机：**
返佣在两个场景触发：

a) **用户充值后（支付回调处理）：**
   - 文件：`backend/internal/service/payment_fulfillment.go` 中 `applyAffiliateRebateForOrder()`
   - 仅对 `balance` 类型订单触发（订阅订单不触发）
   - 返佣金额 = 充值金额 × 返佣比例

b) **余额兑换码被兑换时：**
   - 文件：`backend/internal/service/redeem_service.go` 中 `tryAccrueAffiliateRebateForRedeem()`
   - 仅对 `balance` 类型且正数金额的兑换码触发

**4. 返佣比例：**
- 全局默认比例（由管理员在系统设置中配置）
- 可为特定用户设置个性化比例 (`aff_rebate_rate_percent`)
- 例如：全局 10%，用户充值 $100 → 邀请人获得 $10 佣金

**5. 佣金冻结期：**
- 返佣到账后进入**冻结状态** (`aff_frozen_quota`)
- 冻结时长由管理员配置（小时）
- 后台定时任务自动解冻：`aff_frozen_quota` → `aff_quota`
- 防止被邀请人退款导致的佣金套利

**6. 佣金提现：**
- 用户在分销页面点击 "Transfer to Balance"
- POST `/api/v1/user/aff/transfer`
- 系统将 `aff_quota` 转入用户 `balance`
- 生成审计快照记录

**7. 分销数据模型 (User 表相关字段)：**
| 字段 | 说明 |
|------|------|
| aff_code | 用户的分销邀请码 |
| aff_quota | 可提现佣金余额 |
| aff_frozen_quota | 冻结中佣金 |
| aff_history_quota | 历史累计佣金 |
| aff_count | 邀请的用户数 |
| aff_rebate_rate_percent | 个人返佣比例 (可自定义) |

### 8.4 管理员分销管理

**管理后台页面：**
- `AdminAffiliateInvitesView` - 所有邀请记录
- `AdminAffiliateRebatesView` - 所有返佣记录 (按订单)
- `AdminAffiliateTransfersView` - 所有提现记录 (含快照审计)
- API: `/api/v1/admin/affiliates/*`

### 8.5 分销功能开关

分销功能默认**关闭**，需要管理员在系统设置中启用。功能开关使用 feature flag 控制：
- `affiliate_enabled` - 控制用户端分销页面和管理端分销菜单的显示

---

## 9. 前端架构

### 9.1 路由结构

前端共 40+ 条路由，分为四类：

**公开路由（无需登录）：**
`/home`、`/login`、`/register`、各种 OAuth 回调、`/forgot-password`、`/reset-password`、`/key-usage`、`/setup`

**用户路由（需登录）：**
`/dashboard`、`/keys`、`/usage`、`/redeem`、`/affiliate`、`/available-channels`、`/profile`、`/subscriptions`、`/purchase`、`/orders`、`/payment/*`、`/monitor`

**管理后台路由（需管理员）：**
`/admin/dashboard`、`/admin/ops`、`/admin/users`、`/admin/groups`、`/admin/channels/*`、`/admin/subscriptions`、`/admin/accounts`、`/admin/announcements`、`/admin/proxies`、`/admin/redeem`、`/admin/promo-codes`、`/admin/settings`、`/admin/risk-control`、`/admin/usage`、`/admin/orders/*`、`/admin/affiliates/*`

### 9.2 状态管理 (Pinia)

| Store | 职责 |
|-------|------|
| `auth` | 用户认证状态、JWT Token 管理、自动刷新 |
| `app` | 全局 UI 状态、站点配置、Toast 通知 |
| `adminSettings` | 管理员功能开关缓存 |
| `subscriptions` | 用户订阅状态（自动轮询） |
| `payment` | 支付配置、订单状态 |
| `announcements` | 公告管理和弹窗队列 |
| `onboarding` | 新手引导 (driver.js) |

### 9.3 用户端功能页面

| 页面 | 功能 |
|------|------|
| Dashboard | 用量统计、趋势图表、模型分布、最近使用 |
| API Keys | 创建/管理 API Key，设置 IP 白名单、配额、速率限制 |
| Usage Records | 分页、可筛选的用量日志 |
| Buy Subscription | 套餐选择、多支付方式 |
| My Orders | 订单历史和状态跟踪 |
| Redeem Code | 兑换码激活 |
| Affiliate | 分销邀请码、受邀统计、返佣余额、提现 |
| Profile | 用户信息、密码修改、OAuth 身份绑定、通知邮箱 |
| Subscriptions | 订阅管理和订阅转换 |
| Available Channels | 浏览可用的模型通道和定价 |
| Channel Status | 实时通道健康状态 |

### 9.4 管理后台功能概览

管理后台包含 30+ 个功能模块，涵盖：
- **仪表盘：** API Keys、账号、请求数、成本、RPM/TPM 趋势
- **运维监控：** 实时吞吐量、延迟、错误分布、并发数
- **用户管理：** CRUD、余额调整、分组分配、并发管理
- **分组管理：** 费率倍率、模型路由、RPM限制、用量上限
- **账号管理：** OAuth/API Key 账号管理、代理绑定、TLS指纹、缓存覆盖
- **通道管理：** 模型定价、通道健康监控 (含定时测试)
- **支付管理：** 仪表盘(收入图表)、订单管理(含退款)、套餐管理
- **分销管理：** 邀请记录、返佣记录、提现记录
- **风控管理：** 风险规则配置、API Key 检测

---

## 10. 安全特性

| 特性 | 实现方式 |
|------|---------|
| **API Key 鉴权** | Bearer Token / x-api-key / x-goog-api-key 多种格式 |
| **IP 白名单/黑名单** | API Key 级别的 IP 访问控制 |
| **JWT 认证** | 双 Token 机制 (Access + Refresh)，自动刷新 |
| **TOTP 2FA** | AES-256 加密的 TOTP 密钥，支持双因素认证 |
| **OAuth 2.0** | 支持 6 种 OAuth 提供商 (LinuxDo/WeChat/GitHub/Google/OIDC/Email) |
| **凭据加密** | 数据库中的 API Key/OAuth Token 采用 AES-256-GCM 加密 |
| **CSP 策略** | 可配置的 Content-Security-Policy |
| **CORS** | 可配置的跨域资源共享 |
| **Turnstile** | Cloudflare Turnstile 人机验证 |
| **速率限制** | 基于 Redis 的多级速率限制（认证/支付/API 调用） |
| **并发控制** | 用户级 + 账号级双层并发限制 |
| **幂等性** | SHA-256 请求去重，防止重复扣费 |
| **内容审核** | 请求内容审核拦截 |
| **风控系统** | IP/API Key 风险检测规则、自动封禁和解封 |
| **软删除** | User/APIKey/Account/Group/Subscription 等核心实体使用软删除 |
| **TLS 指纹** | 可配置的 JA3 TLS 指纹模拟，防止被上游识别为代理 |

---

## 11. 部署模型

### Docker Compose 部署 (推荐)
```yaml
# 服务组件:
- sub2api       # Go 后端 + 前端静态文件 (单容器)
- postgres      # PostgreSQL 15
- redis         # Redis 7
```

### 自动初始化
- Docker 环境：自动从环境变量读取配置完成初始化
- 手动部署：通过 Web 设置向导 `/setup` 完成配置

### 配置文件
- 主配置：`config.yaml`（通过 Viper 加载）
- 环境变量支持：`DATA_DIR` 指定配置目录

### CI/CD
- GitHub Actions + GoReleaser
- 多架构 Docker 镜像构建
- 自动化版本发布

---

## 12. 外部集成

| 集成 | 说明 |
|------|------|
| **iframely** | 支持通过 iframe 嵌入外部管理工具 |
| **Claude Code CLI** | 完整支持 Claude Code CLI 代理（系统提示词注入、版本检查） |
| **Codex CLI** | 支持 Codex CLI 直接路由 |
| **Gemini CLI** | 支持 Gemini CLI 会话管理 |
| **LiteLLM** | 定价数据兼容 LiteLLM 格式 |
| **WebSocket** | 支持 OpenAI Responses 的 WebSocket 协议双向代理 |

---

## 13. 后台服务

系统运行时启动 **25+ 个后台服务**，包括：

- **TokenRefreshService** - OAuth Token 自动续期
- **SubscriptionMaintenanceService** - 订阅过期检查和用量窗口重置
- **AccountExpiryService** - 账号过期自动暂停
- **AccountQuotaResetService** - 账号配额定时重置
- **BillingCacheService** - 计费缓存刷新
- **UsageRecordWorkerPool** - 用量记录异步批量写入
- **ChannelMonitorScheduler** - 通道健康监控调度
- **AffiliateThawService** - 分销佣金冻结→解冻
- **DashboardAggregationService** - 仪表盘数据聚合
- **IdempotencyCleanupService** - 幂等记录清理
- **UsageCleanupService** - 用量记录清理任务
- **PaymentOrderExpiryService** - 支付订单超时处理

---

## 14. 关键设计模式总结

| 模式 | 应用场景 |
|------|---------|
| **策略模式** | 支付提供商、计费模式、负载均衡策略 |
| **注册表模式** | 支付提供商注册 |
| **责任链** | 中间件链、账号选择层级 |
| **观察者/回调** | Webhook 通知处理 |
| **模板方法** | 支付提供商接口、API 请求处理管线 |
| **工厂模式** | 支付提供商实例化 |
| **对象池** | HTTP 连接池、Worker Pool |
| **乐观锁** | 兑换码兑换 (`WHERE status='unused'`) |
| **分布式锁** | 兑换码兑换 (Redis 10s TTL) |
| **编译时 DI** | Google Wire 依赖注入 |
