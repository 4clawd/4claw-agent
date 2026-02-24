<div align="center">

  <h1>4claw: Ultra-Efficient AI Assistant in Go</h1>

  <h3>$10 Hardware 鐠?10MB RAM 鐠?1s Boot 鐠?闁活煈鍠氬В濠囨惞閹惧懐绀夐柟瀛樺灣濠婃垹鎸х敮顔剧＜</h3>


[濞戞搩鍘介弸鍍?README.zh.md) | [闁哄啨鍎插﹢鎵嚔閻?README.ja.md) | [Portugu闁挎](README.pt-br.md) | [Ti瀹€椋庝刊g Vi瀹勫嫬妾糫(README.vi.md) | [Fran閼剧禈is](README.fr.md) | **English**

</div>

---

妫ｅ喚娴€ 4claw is an ultra-lightweight personal AI Assistant inspired by [nanobot](https://github.com/HKUDS/nanobot), refactored from the ground up in Go through a self-bootstrapping process, where the AI agent itself drove the entire architectural migration and code optimization.

闁冲簱妲勭粭?Runs on $10 hardware with <10MB RAM: That's 99% less memory than OpenClaw and 98% cheaper than a Mac mini!


> [!CAUTION]
> **妫ｅ啯鐦?SECURITY & OFFICIAL CHANNELS / 閻庣懓顦崣蹇旂珶閻楀牊顫?*
>
> * **NO CRYPTO:** 4claw has **NO** official token/coin. All claims on `pump.fun` or other trading platforms are **SCAMS**.
>
> * **OFFICIAL DOMAIN:** The **ONLY** official website is **[4claw.io](https://4claw.io)**, and company website is **[sipeed.com](https://sipeed.com)**
> * **Warning:** Many `.ai/.org/.com/.net/...` domains are registered by third parties.
> * **Warning:** 4claw is in early development now and may have unresolved network security issues. Do not deploy to production environments before the v1.0 release.
> * **Note:** 4claw has recently merged a lot of PRs, which may result in a larger memory footprint (10闁?0MB) in the latest versions. We plan to prioritize resource optimization as soon as the current feature set reaches a stable state.

## 妫ｅ啯鎲?News

2026-02-16 妫ｅ啫绔?4claw hit 12K stars in one week! Thank you all for your support! 4claw is growing faster than we ever imagined. Given the high volume of PRs, we urgently need community maintainers. Our volunteer roles and roadmap are officially posted [here](docs/ROADMAP.md) 闁炽儲鏁媏 can闁炽儲鐛?wait to have you on board!

2026-02-13 妫ｅ啫绔?4claw hit 5000 stars in 4days! Thank you for the community! There are so many PRs & issues coming in (during Chinese New Year holidays), we are finalizing the Project Roadmap and setting up the Developer Group to accelerate 4claw's development.  
妫ｅ啯鐣?Call to Action: Please submit your feature requests in GitHub Discussions. We will review and prioritize them during our upcoming weekly meeting.

2026-02-09 妫ｅ啫绔?4claw Launched! Built in 1 day to bring AI Agents to $10 hardware with <10MB RAM. 妫ｅ喚娴€ 4claw闁挎稑顒歟t's Go闁?

## 闁?Features

妫ｅ喚鈧?**Ultra-Lightweight**: <10MB Memory footprint 闁?99% smaller than Clawdbot - core functionality.

妫ｅ啯灏?**Minimal Cost**: Efficient enough to run on $10 Hardware 闁?98% cheaper than a Mac mini.

闁冲簱妲勭粭?**Lightning Fast**: 400X Faster startup time, boot in 1 second even in 0.6GHz single core.

妫ｅ啫顕?**True Portability**: Single self-contained binary across RISC-V, ARM, and x86, One-click to Go!

妫ｅ喚妯?**AI-Bootstrapped**: Autonomous Go-native implementation 闁?95% Agent-generated core with human-in-the-loop refinement.

|                               | OpenClaw      | NanoBot                  | **4claw**                              |
| ----------------------------- | ------------- | ------------------------ | ----------------------------------------- |
| **Language**                  | TypeScript    | Python                   | **Go**                                    |
| **RAM**                       | >1GB          | >100MB                   | **< 10MB**                                |
| **Startup**</br>(0.8GHz core) | >500s         | >30s                     | **<1s**                                   |
| **Cost**                      | Mac Mini 599$ | Most Linux SBC </br>~50$ | **Any Linux Board**</br>**As low as 10$** |


## 妫ｅ喚鐎?Demonstration

### 妫ｅ啯绀堥柨?Standard Assistant Workflows


### 妫ｅ啯鎳?Run on old Android Phones

Give your decade-old phone a second life! Turn it into a smart AI Assistant with 4claw. Quick Start:

1. **Install Termux** (Available on F-Droid or Google Play).
2. **Execute cmds**

```bash
# Note: Replace v0.1.1 with the latest version from the Releases page
wget https://github.com/sipeed/4claw/releases/download/v0.1.1/4claw-linux-arm64
chmod +x 4claw-linux-arm64
pkg install proot
termux-chroot ./4claw-linux-arm64 onboard
```

And then follow the instructions in the "Quick Start" section to complete the configuration!

### 妫ｅ啯鍋?Innovative Low-Footprint Deploy

4claw can be deployed on almost any Linux device!

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) E(Ethernet) or W(WiFi6) version, for Minimal Home Assistant
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html), or $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html) for Automated Server Maintenance
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) or $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera) for Smart Monitoring


妫ｅ啫鐨?More Deployment Cases Await闁?

## 妫ｅ啯鎲?Install

### Install with precompiled binary

Download the firmware for your platform from the [release](https://github.com/sipeed/4claw/releases) page.

### Install from source (latest features, recommended for development)

```bash
git clone https://github.com/sipeed/4claw.git

cd 4claw
make deps

# Build, no need to install
make build

# Build for multiple platforms
make build-all

# Build And Install
make install
```

## 妫ｅ啯鍎?Docker Compose

You can also run 4claw using Docker Compose without installing anything locally.

```bash
# 1. Clone this repo
git clone https://github.com/sipeed/4claw.git
cd 4claw

# 2. Set your API keys
cp config/config.example.json config/config.json
vim config/config.json      # Set DISCORD_BOT_TOKEN, API keys, etc.

# 3. Build & Start
docker compose --profile gateway up -d

> [!TIP]
> **Docker Users**: By default, the Gateway listens on `127.0.0.1` which is not accessible from the host. If you need to access the health endpoints or expose ports, set `FOURCLAW_GATEWAY_HOST=0.0.0.0` in your environment or update `config.json`.


# 4. Check logs
docker compose logs -f 4claw-gateway

# 5. Stop
docker compose --profile gateway down
```

### Agent Mode (One-shot)

```bash
# Ask a question
docker compose run --rm 4claw-agent -m "What is 2+2?"

# Interactive mode
docker compose run --rm 4claw-agent
```

### Rebuild

```bash
docker compose --profile gateway build --no-cache
docker compose --profile gateway up -d
```

### 妫ｅ啯鐣?Quick Start

> [!TIP]
> Set your API key in `~/.4claw/config.json`.
> Get API keys: [OpenRouter](https://openrouter.ai/keys) (LLM) 鐠?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> Web Search is **optional** - get free [Tavily API](https://tavily.com) (1000 free queries/month) or [Brave Search API](https://brave.com/search/api) (2000 free queries/month) or use built-in auto fallback.

**1. Initialize**

```bash
4claw onboard
```

**2. Configure** (`~/.4claw/config.json`)

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.4claw/workspace",
      "model": "gpt4",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "model_list": [
    {
      "model_name": "gpt4",
      "model": "openai/gpt-5.2",
      "api_key": "your-api-key"
    },
    {
      "model_name": "claude-sonnet-4.6",
      "model": "anthropic/claude-sonnet-4.6",
      "api_key": "your-anthropic-key"
    }
  ],
  "tools": {
    "web": {
      "brave": {
        "enabled": false,
        "api_key": "YOUR_BRAVE_API_KEY",
        "max_results": 5
      },
      "tavily": {
        "enabled": false,
        "api_key": "YOUR_TAVILY_API_KEY",
        "max_results": 5
      },
      "duckduckgo": {
        "enabled": true,
        "max_results": 5
      }
    }
  }
}
```

> **New**: The `model_list` configuration format allows zero-code provider addition. See [Model Configuration](#model-configuration-model_list) for details.

**3. Get API Keys**

* **LLM Provider**: [OpenRouter](https://openrouter.ai/keys) 鐠?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) 鐠?[Anthropic](https://console.anthropic.com) 鐠?[OpenAI](https://platform.openai.com) 鐠?[Gemini](https://aistudio.google.com/api-keys)
* **Web Search** (optional): [Tavily](https://tavily.com) - Optimized for AI Agents (1000 requests/month) 鐠?[Brave Search](https://brave.com/search/api) - Free tier available (2000 requests/month)

> **Note**: See `config.example.json` for a complete configuration template.

**4. Chat**

```bash
4claw agent -m "What is 2+2?"
```

That's it! You have a working AI assistant in 2 minutes.

---

## 妫ｅ啯灏?Chat Apps

Talk to your 4claw through Telegram, Discord, DingTalk, LINE, or WeCom

| Channel      | Setup                              |
| ------------ | ---------------------------------- |
| **Telegram** | Easy (just a token)                |
| **Discord**  | Easy (bot token + intents)         |
| **QQ**       | Easy (AppID + AppSecret)           |
| **DingTalk** | Medium (app credentials)           |
| **LINE**     | Medium (credentials + webhook URL) |
| **WeCom**    | Medium (CorpID + webhook setup)    |

<details>
<summary><b>Telegram</b> (Recommended)</summary>

**1. Create a bot**

* Open Telegram, search `@BotFather`
* Send `/newbot`, follow prompts
* Copy the token

**2. Configure**

```json
{
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "YOUR_BOT_TOKEN",
      "allow_from": ["YOUR_USER_ID"]
    }
  }
}
```

> Get your user ID from `@userinfobot` on Telegram.

**3. Run**

```bash
4claw gateway
```

</details>

<details>
<summary><b>Discord</b></summary>

**1. Create a bot**

* Go to <https://discord.com/developers/applications>
* Create an application 闁?Bot 闁?Add Bot
* Copy the bot token

**2. Enable intents**

* In the Bot settings, enable **MESSAGE CONTENT INTENT**
* (Optional) Enable **SERVER MEMBERS INTENT** if you plan to use allow lists based on member data

**3. Get your User ID**
* Discord Settings 闁?Advanced 闁?enable **Developer Mode**
* Right-click your avatar 闁?**Copy User ID**

**4. Configure**

```json
{
  "channels": {
    "discord": {
      "enabled": true,
      "token": "YOUR_BOT_TOKEN",
      "allow_from": ["YOUR_USER_ID"],
      "mention_only": false
    }
  }
}
```

**5. Invite the bot**

* OAuth2 闁?URL Generator
* Scopes: `bot`
* Bot Permissions: `Send Messages`, `Read Message History`
* Open the generated invite URL and add the bot to your server

**Optional: Mention-only mode**

Set `"mention_only": true` to make the bot respond only when @-mentioned. Useful for shared servers where you want the bot to respond only when explicitly called.

**6. Run**

```bash
4claw gateway
```

</details>

<details>
<summary><b>QQ</b></summary>

**1. Create a bot**

- Go to [QQ Open Platform](https://q.qq.com/#)
- Create an application 闁?Get **AppID** and **AppSecret**

**2. Configure**

```json
{
  "channels": {
    "qq": {
      "enabled": true,
      "app_id": "YOUR_APP_ID",
      "app_secret": "YOUR_APP_SECRET",
      "allow_from": []
    }
  }
}
```

> Set `allow_from` to empty to allow all users, or specify QQ numbers to restrict access.

**3. Run**

```bash
4claw gateway
```

</details>

<details>
<summary><b>DingTalk</b></summary>

**1. Create a bot**

* Go to [Open Platform](https://open.dingtalk.com/)
* Create an internal app
* Copy Client ID and Client Secret

**2. Configure**

```json
{
  "channels": {
    "dingtalk": {
      "enabled": true,
      "client_id": "YOUR_CLIENT_ID",
      "client_secret": "YOUR_CLIENT_SECRET",
      "allow_from": []
    }
  }
}
```

> Set `allow_from` to empty to allow all users, or specify DingTalk user IDs to restrict access.

**3. Run**

```bash
4claw gateway
```
</details>

<details>
<summary><b>LINE</b></summary>

**1. Create a LINE Official Account**

- Go to [LINE Developers Console](https://developers.line.biz/)
- Create a provider 闁?Create a Messaging API channel
- Copy **Channel Secret** and **Channel Access Token**

**2. Configure**

```json
{
  "channels": {
    "line": {
      "enabled": true,
      "channel_secret": "YOUR_CHANNEL_SECRET",
      "channel_access_token": "YOUR_CHANNEL_ACCESS_TOKEN",
      "webhook_host": "0.0.0.0",
      "webhook_port": 18791,
      "webhook_path": "/webhook/line",
      "allow_from": []
    }
  }
}
```

**3. Set up Webhook URL**

LINE requires HTTPS for webhooks. Use a reverse proxy or tunnel:

```bash
# Example with ngrok
ngrok http 18791
```

Then set the Webhook URL in LINE Developers Console to `https://your-domain/webhook/line` and enable **Use webhook**.

**4. Run**

```bash
4claw gateway
```

> In group chats, the bot responds only when @mentioned. Replies quote the original message.

> **Docker Compose**: Add `ports: ["18791:18791"]` to the `4claw-gateway` service to expose the webhook port.

</details>

<details>
<summary><b>WeCom (濞撮棿妞掔粭鐔奉嚗椤旇绻?</b></summary>

4claw supports two types of WeCom integration:

**Option 1: WeCom Bot (闁哄懘缂氶崗姗€寮甸崫鍕彜濞?** - Easier setup, supports group chats
**Option 2: WeCom App (闁煎浜滅紓鎾存償閺冨倹鏆?** - More features, proactive messaging

See [WeCom App Configuration Guide](docs/wecom-app-configuration.md) for detailed setup instructions.

**Quick Setup - WeCom Bot:**

**1. Create a bot**

* Go to WeCom Admin Console 闁?Group Chat 闁?Add Group Bot
* Copy the webhook URL (format: `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx`)

**2. Configure**

```json
{
  "channels": {
    "wecom": {
      "enabled": true,
      "token": "YOUR_TOKEN",
      "encoding_aes_key": "YOUR_ENCODING_AES_KEY",
      "webhook_url": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=YOUR_KEY",
      "webhook_host": "0.0.0.0",
      "webhook_port": 18793,
      "webhook_path": "/webhook/wecom",
      "allow_from": []
    }
  }
}
```

**Quick Setup - WeCom App:**

**1. Create an app**

* Go to WeCom Admin Console 闁?App Management 闁?Create App
* Copy **AgentId** and **Secret**
* Go to "My Company" page, copy **CorpID**
**2. Configure receive message**

* In App details, click "Receive Message" 闁?"Set API"
* Set URL to `http://your-server:18792/webhook/wecom-app`
* Generate **Token** and **EncodingAESKey**

**3. Configure**

```json
{
  "channels": {
    "wecom_app": {
      "enabled": true,
      "corp_id": "wwxxxxxxxxxxxxxxxx",
      "corp_secret": "YOUR_CORP_SECRET",
      "agent_id": 1000002,
      "token": "YOUR_TOKEN",
      "encoding_aes_key": "YOUR_ENCODING_AES_KEY",
      "webhook_host": "0.0.0.0",
      "webhook_port": 18792,
      "webhook_path": "/webhook/wecom-app",
      "allow_from": []
    }
  }
}
```

**4. Run**

```bash
4claw gateway
```

> **Note**: WeCom App requires opening port 18792 for webhook callbacks. Use a reverse proxy for HTTPS.

</details>


Connect 4claw to the Agent Social Network simply by sending a single message via the CLI or any integrated Chat App.

**Read `https://clawdchat.ai/skill.md` and follow the instructions to join [ClawdChat.ai](https://clawdchat.ai)**

## 闁虫寧鐟辩粭?Configuration

Config file: `~/.4claw/config.json`

### Workspace Layout

4claw stores data in your configured workspace (default: `~/.4claw/workspace`):

```
~/.4claw/workspace/
闁宠澹曢弨銏ゅ煘閳?sessions/          # Conversation sessions and history
闁宠澹曢弨銏ゅ煘閳?memory/           # Long-term memory (MEMORY.md)
闁宠澹曢弨銏ゅ煘閳?state/            # Persistent state (last channel, etc.)
闁宠澹曢弨銏ゅ煘閳?cron/             # Scheduled jobs database
闁宠澹曢弨銏ゅ煘閳?skills/           # Custom skills
闁宠澹曢弨銏ゅ煘閳?AGENTS.md         # Agent behavior guide
闁宠澹曢弨銏ゅ煘閳?HEARTBEAT.md      # Periodic task prompts (checked every 30 min)
闁宠澹曢弨銏ゅ煘閳?IDENTITY.md       # Agent identity
闁宠澹曢弨銏ゅ煘閳?SOUL.md           # Agent soul
闁宠澹曢弨銏ゅ煘閳?TOOLS.md          # Tool descriptions
闁宠鏌￠弨銏ゅ煘閳?USER.md           # User preferences
```

### 妫ｅ啯鏅?Security Sandbox

4claw runs in a sandboxed environment by default. The agent can only access files and execute commands within the configured workspace.

#### Default Configuration

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.4claw/workspace",
      "restrict_to_workspace": true
    }
  }
}
```

| Option                  | Default                 | Description                               |
| ----------------------- | ----------------------- | ----------------------------------------- |
| `workspace`             | `~/.4claw/workspace` | Working directory for the agent           |
| `restrict_to_workspace` | `true`                  | Restrict file/command access to workspace |

#### Protected Tools

When `restrict_to_workspace: true`, the following tools are sandboxed:

| Tool          | Function         | Restriction                            |
| ------------- | ---------------- | -------------------------------------- |
| `read_file`   | Read files       | Only files within workspace            |
| `write_file`  | Write files      | Only files within workspace            |
| `list_dir`    | List directories | Only directories within workspace      |
| `edit_file`   | Edit files       | Only files within workspace            |
| `append_file` | Append to files  | Only files within workspace            |
| `exec`        | Execute commands | Command paths must be within workspace |

#### Additional Exec Protection

Even with `restrict_to_workspace: false`, the `exec` tool blocks these dangerous commands:

* `rm -rf`, `del /f`, `rmdir /s` 闁?Bulk deletion
* `format`, `mkfs`, `diskpart` 闁?Disk formatting
* `dd if=` 闁?Disk imaging
* Writing to `/dev/sd[a-z]` 闁?Direct disk writes
* `shutdown`, `reboot`, `poweroff` 闁?System shutdown
* Fork bomb `:(){ :|:& };:`

#### Error Examples

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (path outside working dir)}
```

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (dangerous pattern detected)}
```

#### Disabling Restrictions (Security Risk)

If you need the agent to access paths outside the workspace:

**Method 1: Config file**

```json
{
  "agents": {
    "defaults": {
      "restrict_to_workspace": false
    }
  }
}
```

**Method 2: Environment variable**

```bash
export FOURCLAW_AGENTS_DEFAULTS_RESTRICT_TO_WORKSPACE=false
```

> 闁宠法濯寸粭?**Warning**: Disabling this restriction allows the agent to access any path on your system. Use with caution in controlled environments only.

#### Security Boundary Consistency

The `restrict_to_workspace` setting applies consistently across all execution paths:

| Execution Path   | Security Boundary            |
| ---------------- | ---------------------------- |
| Main Agent       | `restrict_to_workspace` 闁?  |
| Subagent / Spawn | Inherits same restriction 闁?|
| Heartbeat tasks  | Inherits same restriction 闁?|

All paths share the same workspace restriction 闁?there's no way to bypass the security boundary through subagents or scheduled tasks.

### Heartbeat (Periodic Tasks)

4claw can perform periodic tasks automatically. Create a `HEARTBEAT.md` file in your workspace:

```markdown
# Periodic Tasks

- Check my email for important messages
- Review my calendar for upcoming events
- Check the weather forecast
```

The agent will read this file every 30 minutes (configurable) and execute any tasks using available tools.

#### Async Tasks with Spawn

For long-running tasks (web search, API calls), use the `spawn` tool to create a **subagent**:

```markdown
# Periodic Tasks

## Quick Tasks (respond directly)

- Report current time

## Long Tasks (use spawn for async)

- Search the web for AI news and summarize
- Check email and report important messages
```

**Key behaviors:**

| Feature                 | Description                                               |
| ----------------------- | --------------------------------------------------------- |
| **spawn**               | Creates async subagent, doesn't block heartbeat           |
| **Independent context** | Subagent has its own context, no session history          |
| **message tool**        | Subagent communicates with user directly via message tool |
| **Non-blocking**        | After spawning, heartbeat continues to next task          |

#### How Subagent Communication Works

```
Heartbeat triggers
    闁?
Agent reads HEARTBEAT.md
    闁?
For long task: spawn subagent
    闁?                          闁?
Continue to next task      Subagent works independently
    闁?                          闁?
All tasks done            Subagent uses "message" tool
    闁?                          闁?
Respond HEARTBEAT_OK      User receives result directly
```

The subagent has access to tools (message, web_search, etc.) and can communicate with the user independently without going through the main agent.

**Configuration:**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| Option     | Default | Description                        |
| ---------- | ------- | ---------------------------------- |
| `enabled`  | `true`  | Enable/disable heartbeat           |
| `interval` | `30`    | Check interval in minutes (min: 5) |

**Environment variables:**

* `FOURCLAW_HEARTBEAT_ENABLED=false` to disable
* `FOURCLAW_HEARTBEAT_INTERVAL=60` to change interval

### Providers

> [!NOTE]
> Groq provides free voice transcription via Whisper. If configured, Telegram voice messages will be automatically transcribed.

| Provider                   | Purpose                                 | Get API Key                                                          |
| -------------------------- | --------------------------------------- | -------------------------------------------------------------------- |
| `gemini`                   | LLM (Gemini direct)                     | [aistudio.google.com](https://aistudio.google.com)                   |
| `zhipu`                    | LLM (Zhipu direct)                      | [bigmodel.cn](https://bigmodel.cn)                                   |
| `openrouter(To be tested)` | LLM (recommended, access to all models) | [openrouter.ai](https://openrouter.ai)                               |
| `anthropic(To be tested)`  | LLM (Claude direct)                     | [console.anthropic.com](https://console.anthropic.com)               |
| `openai(To be tested)`     | LLM (GPT direct)                        | [platform.openai.com](https://platform.openai.com)                   |
| `deepseek(To be tested)`   | LLM (DeepSeek direct)                   | [platform.deepseek.com](https://platform.deepseek.com)               |
| `qwen`                     | LLM (Qwen direct)                       | [dashscope.console.aliyun.com](https://dashscope.console.aliyun.com) |
| `groq`                     | LLM + **Voice transcription** (Whisper) | [console.groq.com](https://console.groq.com)                         |
| `cerebras`                 | LLM (Cerebras direct)                   | [cerebras.ai](https://cerebras.ai)                                   |

### Model Configuration (model_list)

> **What's New?** 4claw now uses a **model-centric** configuration approach. Simply specify `vendor/model` format (e.g., `zhipu/glm-4.7`) to add new providers闁?*zero code changes required!**

This design also enables **multi-agent support** with flexible provider selection:

- **Different agents, different providers**: Each agent can use its own LLM provider
- **Model fallbacks**: Configure primary and fallback models for resilience
- **Load balancing**: Distribute requests across multiple endpoints
- **Centralized configuration**: Manage all providers in one place

#### 妫ｅ啯鎯?All Supported Vendors

| Vendor              | `model` Prefix    | Default API Base                                    | Protocol  | API Key                                                          |
| ------------------- | ----------------- | --------------------------------------------------- | --------- | ---------------------------------------------------------------- |
| **OpenAI**          | `openai/`         | `https://api.openai.com/v1`                         | OpenAI    | [Get Key](https://platform.openai.com)                           |
| **Anthropic**       | `anthropic/`      | `https://api.anthropic.com/v1`                      | Anthropic | [Get Key](https://console.anthropic.com)                         |
| **闁哄懘缂氬?AI (GLM)**   | `zhipu/`          | `https://open.bigmodel.cn/api/paas/v4`              | OpenAI    | [Get Key](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek**        | `deepseek/`       | `https://api.deepseek.com/v1`                       | OpenAI    | [Get Key](https://platform.deepseek.com)                         |
| **Google Gemini**   | `gemini/`         | `https://generativelanguage.googleapis.com/v1beta`  | OpenAI    | [Get Key](https://aistudio.google.com/api-keys)                  |
| **Groq**            | `groq/`           | `https://api.groq.com/openai/v1`                    | OpenAI    | [Get Key](https://console.groq.com)                              |
| **Moonshot**        | `moonshot/`       | `https://api.moonshot.cn/v1`                        | OpenAI    | [Get Key](https://platform.moonshot.cn)                          |
| **闂侇偅鐭粻鐔煎础閸愵喗锛?(Qwen)** | `qwen/`           | `https://dashscope.aliyuncs.com/compatible-mode/v1` | OpenAI    | [Get Key](https://dashscope.console.aliyun.com)                  |
| **NVIDIA**          | `nvidia/`         | `https://integrate.api.nvidia.com/v1`               | OpenAI    | [Get Key](https://build.nvidia.com)                              |
| **Ollama**          | `ollama/`         | `http://localhost:11434/v1`                         | OpenAI    | Local (no key needed)                                            |
| **OpenRouter**      | `openrouter/`     | `https://openrouter.ai/api/v1`                      | OpenAI    | [Get Key](https://openrouter.ai/keys)                            |
| **VLLM**            | `vllm/`           | `http://localhost:8000/v1`                          | OpenAI    | Local                                                            |
| **Cerebras**        | `cerebras/`       | `https://api.cerebras.ai/v1`                        | OpenAI    | [Get Key](https://cerebras.ai)                                   |
| **闁诲浚鍋勯崠妤€顕ｉ弴鐔告儧**        | `volcengine/`     | `https://ark.cn-beijing.volces.com/api/v3`          | OpenAI    | [Get Key](https://console.volcengine.com)                        |
| **缂佷胶鍋熼悾缁樼?*          | `shengsuanyun/`   | `https://router.shengsuanyun.com/api/v1`            | OpenAI    | -                                                                |
| **Antigravity**     | `antigravity/`    | Google Cloud                                        | Custom    | OAuth only                                                       |
| **GitHub Copilot**  | `github-copilot/` | `localhost:4321`                                    | gRPC      | -                                                                |

#### Basic Configuration

```json
{
  "model_list": [
    {
      "model_name": "gpt-5.2",
      "model": "openai/gpt-5.2",
      "api_key": "sk-your-openai-key"
    },
    {
      "model_name": "claude-sonnet-4.6",
      "model": "anthropic/claude-sonnet-4.6",
      "api_key": "sk-ant-your-key"
    },
    {
      "model_name": "glm-4.7",
      "model": "zhipu/glm-4.7",
      "api_key": "your-zhipu-key"
    }
  ],
  "agents": {
    "defaults": {
      "model": "gpt-5.2"
    }
  }
}
```

#### Vendor-Specific Examples

**OpenAI**

```json
{
  "model_name": "gpt-5.2",
  "model": "openai/gpt-5.2",
  "api_key": "sk-..."
}
```

**闁哄懘缂氬?AI (GLM)**

```json
{
  "model_name": "glm-4.7",
  "model": "zhipu/glm-4.7",
  "api_key": "your-key"
}
```

**DeepSeek**

```json
{
  "model_name": "deepseek-chat",
  "model": "deepseek/deepseek-chat",
  "api_key": "sk-..."
}
```

**Anthropic (with API key)**

```json
{
  "model_name": "claude-sonnet-4.6",
  "model": "anthropic/claude-sonnet-4.6",
  "api_key": "sk-ant-your-key"
}
```

> Run `4claw auth login --provider anthropic` to paste your API token.

**Ollama (local)**

```json
{
  "model_name": "llama3",
  "model": "ollama/llama3"
}
```

**Custom Proxy/API**

```json
{
  "model_name": "my-custom-model",
  "model": "openai/custom-model",
  "api_base": "https://my-proxy.com/v1",
  "api_key": "sk-..."
}
```

#### Load Balancing

Configure multiple endpoints for the same model name闁炽儲鎽╥coClaw will automatically round-robin between them:

```json
{
  "model_list": [
    {
      "model_name": "gpt-5.2",
      "model": "openai/gpt-5.2",
      "api_base": "https://api1.example.com/v1",
      "api_key": "sk-key1"
    },
    {
      "model_name": "gpt-5.2",
      "model": "openai/gpt-5.2",
      "api_base": "https://api2.example.com/v1",
      "api_key": "sk-key2"
    }
  ]
}
```

#### Migration from Legacy `providers` Config

The old `providers` configuration is **deprecated** but still supported for backward compatibility.

**Old Config (deprecated):**

```json
{
  "providers": {
    "zhipu": {
      "api_key": "your-key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  },
  "agents": {
    "defaults": {
      "provider": "zhipu",
      "model": "glm-4.7"
    }
  }
}
```

**New Config (recommended):**

```json
{
  "model_list": [
    {
      "model_name": "glm-4.7",
      "model": "zhipu/glm-4.7",
      "api_key": "your-key"
    }
  ],
  "agents": {
    "defaults": {
      "model": "glm-4.7"
    }
  }
}
```

For detailed migration guide, see [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md).

### Provider Architecture

4claw routes providers by protocol family:

- OpenAI-compatible protocol: OpenRouter, OpenAI-compatible gateways, Groq, Zhipu, and vLLM-style endpoints.
- Anthropic protocol: Claude-native API behavior.
- Codex/OAuth path: OpenAI OAuth/token authentication route.

This keeps the runtime lightweight while making new OpenAI-compatible backends mostly a config operation (`api_base` + `api_key`).

<details>
<summary><b>Zhipu</b></summary>

**1. Get API key and base URL**

* Get [API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. Configure**

```json
{
  "agents": {
    "defaults": {
      "workspace": "~/.4claw/workspace",
      "model": "glm-4.7",
      "max_tokens": 8192,
      "temperature": 0.7,
      "max_tool_iterations": 20
    }
  },
  "providers": {
    "zhipu": {
      "api_key": "Your API Key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  }
}
```

**3. Run**

```bash
4claw agent -m "Hello"
```

</details>

<details>
<summary><b>Full config example</b></summary>

```json
{
  "agents": {
    "defaults": {
      "model": "anthropic/claude-opus-4-5"
    }
  },
  "providers": {
    "openrouter": {
      "api_key": "sk-or-v1-xxx"
    },
    "groq": {
      "api_key": "gsk_xxx"
    }
  },
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "123456:ABC...",
      "allow_from": ["123456789"]
    },
    "discord": {
      "enabled": true,
      "token": "",
      "allow_from": [""]
    },
    "whatsapp": {
      "enabled": false
    },
    "feishu": {
      "enabled": false,
      "app_id": "cli_xxx",
      "app_secret": "xxx",
      "encrypt_key": "",
      "verification_token": "",
      "allow_from": []
    },
    "qq": {
      "enabled": false,
      "app_id": "",
      "app_secret": "",
      "allow_from": []
    }
  },
  "tools": {
    "web": {
      "brave": {
        "enabled": false,
        "api_key": "BSA...",
        "max_results": 5
      },
      "duckduckgo": {
        "enabled": true,
        "max_results": 5
      }
    },
    "cron": {
      "exec_timeout_minutes": 5
    }
  },
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

</details>

## CLI Reference

| Command                   | Description                   |
| ------------------------- | ----------------------------- |
| `4claw onboard`        | Initialize config & workspace |
| `4claw agent -m "..."` | Chat with the agent           |
| `4claw agent`          | Interactive chat mode         |
| `4claw gateway`        | Start the gateway             |
| `4claw status`         | Show status                   |
| `4claw cron list`      | List all scheduled jobs       |
| `4claw cron add ...`   | Add a scheduled job           |

### Scheduled Tasks / Reminders

4claw supports scheduled reminders and recurring tasks through the `cron` tool:

* **One-time reminders**: "Remind me in 10 minutes" 闁?triggers once after 10min
* **Recurring tasks**: "Remind me every 2 hours" 闁?triggers every 2 hours
* **Cron expressions**: "Remind me at 9am daily" 闁?uses cron expression

Jobs are stored in `~/.4claw/workspace/cron/` and processed automatically.

## 妫ｅ喚妾?Contribute & Roadmap

PRs welcome! The codebase is intentionally small and readable. 妫ｅ喚妯?
See our full [Community Roadmap](https://github.com/sipeed/4claw/blob/main/ROADMAP.md).

Developer group building, join after your first merged PR!

User Groups:

discord: <https://discord.gg/V4sAZ9XWpN>


## 妫ｅ啯鍋?Troubleshooting

### Web search says "API 闂佹澘绉堕悿鍡涙⒒椤曗偓椤?

This is normal if you haven't configured a search API key yet. 4claw will provide helpful links for manual searching.

To enable web search:

1. **Option 1 (Recommended)**: Get a free API key at [https://brave.com/search/api](https://brave.com/search/api) (2000 free queries/month) for the best results.
2. **Option 2 (No Credit Card)**: If you don't have a key, we automatically fall back to **DuckDuckGo** (no key required).

Add the key to `~/.4claw/config.json` if using Brave:

```json
{
  "tools": {
    "web": {
      "brave": {
        "enabled": false,
        "api_key": "YOUR_BRAVE_API_KEY",
        "max_results": 5
      },
      "duckduckgo": {
        "enabled": true,
        "max_results": 5
      }
    }
  }
}
```

### Getting content filtering errors

Some providers (like Zhipu) have content filtering. Try rephrasing your query or use a different model.

### Telegram bot says "Conflict: terminated by other getUpdates"

This happens when another instance of the bot is running. Make sure only one `4claw gateway` is running at a time.

---

## 妫ｅ啯鎲?API Key Comparison

| Service          | Free Tier           | Use Case                              |
| ---------------- | ------------------- | ------------------------------------- |
| **OpenRouter**   | 200K tokens/month   | Multiple models (Claude, GPT-4, etc.) |
| **Zhipu**        | 200K tokens/month   | Best for Chinese users                |
| **Brave Search** | 2000 queries/month  | Web search functionality              |
| **Groq**         | Free tier available | Fast inference (Llama, Mixtral)       |
| **Cerebras**     | Free tier available | Fast inference (Llama, Qwen, etc.)    |
