<div align="center">
  <a href="./README.md"><kbd>English (Default)</kbd></a>
  <a href="./README.zh.md"><kbd>简体中文</kbd></a>
  <a href="./README.ja.md"><kbd>日本語</kbd></a>
  <a href="./README.fr.md"><kbd>Français</kbd></a>
  <a href="./README.pt-br.md"><kbd>Português (Brasil)</kbd></a>
  <a href="./README.vi.md"><kbd>Tiếng Việt</kbd></a>
</div>

<br />

<div align="center">
  <img src="./docs/images/logo.png" alt="4claw logo" width="128" />
</div>

<div align="center">
  <img src="./docs/images/banner.png" alt="4claw banner" width="100%" />
</div>

# 4claw Agent Core

`agent/` 是 `cli/` 桌面端背后的**核心运行进程仓库**。

可以这样理解：

- `agent/`：核心引擎（Go 二进制、配置解析、工具系统、gateway 运行）
- `cli/`：桌面控制台（Electron 可视化管理、窗口与托盘交互）

## 与 `cli/` 的关系

`cli/` 不替代 `agent/`，而是管理它：

1. 写入和修改 `config.json`
2. 拉起/停止 `4claw gateway --config ...`
3. 读取运行日志
4. 处理备份导入导出与恢复

因此，`agent/` 是真正执行任务的核心进程，`cli/` 是操作面板。

## `agent/` 实现的核心能力

- Go 编写的单二进制运行时
- Gateway 模式与多渠道接入
- 基于配置的模型/Provider 路由
- 工具系统（文件、Shell、Web、定时、Skills）
- Agent 执行循环与任务编排
- 可移植的配置驱动部署模式

## CLI 控制台截图（控制层）

虽然本仓库是核心运行时，以下是控制它的 `cli/` 界面截图：

### 主面板

![CLI Main](./docs/images/main.png)

### 设置面板

![CLI Settings](./docs/images/setting.png)

## 仓库结构

- `cmd/`：程序入口与命令
- `config/`：配置样例与默认项
- `internal/`：核心业务逻辑
- `docs/`：设计与路线图文档
- `docs/images/`：README 使用的 logo/banner/截图资源

## 快速开始

```bash
make deps
make build
```

构建后运行：

```bash
./4claw gateway
```

使用自定义配置运行：

```bash
./4claw gateway -c /path/to/config.json
```

## 与桌面端配合的典型流程

1. 在 `agent/` 构建或下载运行时二进制。
2. 放入 `cli/resources/bin/`。
3. 启动 `cli/` Electron 程序。
4. 在桌面端管理多个 Agent 实例。

## 这种分层带来的好处

- 核心运行时更轻、更易移植
- GUI 与核心解耦，迭代更快
- 支持无界面服务器部署
- 支持桌面端多实例安全运维

## License

MIT
