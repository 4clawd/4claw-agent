# 4claw

4claw is an ultra-lightweight personal AI assistant written in Go.

## Key points

- Single binary deployment on Linux systems.
- Works with multiple model providers.
- Supports chat channels such as Telegram, Discord, Slack, LINE, WeCom, and others.
- Includes tools for shell, filesystem, web search, scheduling, and skills.

## Quick start

1. Build the binary:

```bash
make build
```

2. Initialize local config and workspace:

```bash
4claw onboard
```

3. Configure your model provider in:

- `~/.4claw/config.json`

4. Run one-shot message:

```bash
4claw agent -m "Hello"
```

5. Start gateway mode:

```bash
4claw gateway
```

6. Start gateway with a custom config path:

```bash
4claw gateway -c /path/to/config.json
```

## Docker compose

```bash
cp config/config.example.json config/config.json
# edit keys in config/config.json
docker compose --profile gateway up -d
```

## Main commands

- `4claw onboard`
- `4claw agent`
- `4claw gateway`
- `4claw status`
- `4claw auth login --provider openai`
- `4claw cron list`
- `4claw skills list`

## Configuration

Default paths:

- Config: `~/.4claw/config.json`
- Workspace: `~/.4claw/workspace`

Environment variable prefix:

- `FOURCLAW_`

## Development

```bash
make deps
make build
```

## License

MIT
