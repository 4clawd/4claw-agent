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

This `agent/` repository is the **core runtime process** behind the `cli/` desktop application.

In simple terms:

- `agent/` = core engine (Go binary, config parsing, tools, gateway runtime)
- `cli/` = desktop UI shell (Electron app, visual management, tray/window UX)

## Relationship with `cli/`

The desktop app in `cli/` does not replace this runtime.  
It manages this runtime:

1. writes `config.json`
2. launches/stops the binary (`4claw ... gateway --config ...`)
3. reads logs
4. manages backups/import/export

So the `agent/` directory is the actual execution core, while `cli/` is the control plane.

## Core Capabilities in `agent/`

- Go-based single-binary runtime
- Gateway mode for channel integrations
- Model/provider routing through config
- Tool system (filesystem, shell, web, scheduling, skills)
- Agent execution loop and task orchestration
- Config-driven behavior with portable deployment

## CLI Desktop Screenshots (Control Plane)

Although this repo is runtime-core, here are the `cli/` interface screenshots that operate it:

### Main Panel

![CLI Main](./docs/images/main.png)

### Settings Panel

![CLI Settings](./docs/images/setting.png)

## Repository Structure

- `cmd/` and entrypoints for agent runtime
- `config/` examples and defaults
- `internal/` core logic
- `docs/` design and roadmap materials
- `docs/images/` branding and screenshot assets used in README files

## Quick Start

```bash
make deps
make build
```

Then run:

```bash
./4claw gateway
```

Or with custom config:

```bash
./4claw gateway -c /path/to/config.json
```

## Typical Workflow with Desktop

1. Build/download this runtime binary from `agent/`.
2. Put the binary in `cli/resources/bin/`.
3. Start the Electron app (`cli/`).
4. Manage multiple runtime instances visually.

## Why This Split Is Useful

- Keeps the runtime lean and portable
- Keeps desktop UX separate and faster to iterate
- Allows server/headless usage without UI
- Allows desktop users to operate multiple agents safely

## License

MIT
