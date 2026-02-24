<div align="center">

<h1>4claw: Assistente de IA Ultra-Eficiente em Go</h1>

<h3>Hardware de $10 鐠?10MB de RAM 鐠?Boot em 1s 鐠?闁活煈鍠氬В濠囨惞閹惧懐绀夐柟瀛樺灣濠婃垹鎸х敮顔剧＜</h3>


 [濞戞搩鍘介弸鍍?README.zh.md) | [闁哄啨鍎插﹢鎵嚔閻?README.ja.md) | **Portugu闁挎** | [Ti瀹€椋庝刊g Vi瀹勫嫬妾糫(README.vi.md) | [Fran閼剧禈is](README.fr.md) | [English](README.md)
</div>

---

妫ｅ喚娴€ **4claw** 閼?um assistente pessoal de IA ultra-leve inspirado no [nanobot](https://github.com/HKUDS/nanobot), reescrito do zero em **Go** por meio de um processo de "auto-inicializa閼惧€熷皸o" (self-bootstrapping) 闁?onde o pr鐠愮珴rio agente de IA conduziu toda a migra閼惧€熷皸o de arquitetura e otimiza閼惧€熷皸o de c鐠愮珡igo.

闁冲簱妲勭粭?**Extremamente leve:** Roda em hardware de apenas **$10** com **<10MB** de RAM. Isso 閼?99% menos mem鐠愮ia que o OpenClaw e 98% mais barato que um Mac mini!


> [!CAUTION]
> **妫ｅ啯鐦?DECLARA閼寸喕鍓籓 DE SEGURAN閼寸儩 & CANAIS OFICIAIS**
>
> * **SEM CRIPTOMOEDAS:** O 4claw **N閼存┉** possui nenhum token/moeda oficial. Todas as alega閼剧晫甯縠s no `pump.fun` ou outras plataformas de negocia閼惧€熷皸o s閼肩幈 **GOLPES**.
> * **DOM閼搭檾IO OFICIAL:** O **閼寸óICO** site oficial 閼?o **[4claw.io](https://4claw.io)**, e o site da empresa 閼?o **[sipeed.com](https://sipeed.com)**.
> * **Aviso:** Muitos dom闁惧攳ios `.ai/.org/.com/.net/...` foram registrados por terceiros, n閼肩幈 s閼肩幈 nossos.
> * **Aviso:** O 4claw est鐠?em fase inicial de desenvolvimento e pode ter problemas de seguran閼剧禈 de rede n閼肩幈 resolvidos. N閼肩幈 implante em ambientes de produ閼惧€熷皸o antes da vers閼肩幈 v1.0.
> * **Nota:** O 4claw recentemente fez merge de muitos PRs, o que pode resultar em maior consumo de mem鐠愮ia (10-20MB) nas vers閻滅帯s mais recentes. Planejamos priorizar a otimiza閼惧€熷皸o de recursos assim que o conjunto de funcionalidades estiver est鐠嬶箼el.


## 妫ｅ啯鎲?Novidades

2026-02-16 妫ｅ啫绔?4claw atingiu 12K stars em uma semana! Obrigado a todos pelo apoio! O 4claw est鐠?crescendo mais r鐠嬶箲ido do que jamais imaginamos. Dado o alto volume de PRs, precisamos urgentemente de maintainers da comunidade. Nossos pap閼煎崨s de volunt鐠嬶箶ios e roadmap foram publicados oficialmente [aqui](docs/ROADMAP.md) 闁?estamos ansiosos para ter voc闁?a bordo!

2026-02-13 妫ｅ啫绔?4claw atingiu 5000 stars em 4 dias! Obrigado 閼?comunidade! Estamos finalizando o **Roadmap do Projeto** e configurando o **Grupo de Desenvolvedores** para acelerar o desenvolvimento do 4claw.

妫ｅ啯鐣?**Chamada para A閼惧€熷皸o:** Envie suas solicita閼剧晫甯縠s de funcionalidades nas GitHub Discussions. Revisaremos e priorizaremos na pr鐠愮ima reuni閼肩幈 semanal.

2026-02-09 妫ｅ啫绔?4claw lan閼剧禈do oficialmente! Constru闁惧摼o em 1 dia para trazer Agentes de IA para hardware de $10 com <10MB de RAM. 妫ｅ喚娴€ 4claw, Partiu!

## 闁?Funcionalidades

妫ｅ喚鈧?**Ultra-Leve**: Consumo de mem鐠愮ia <10MB 闁?99% menor que o Clawdbot para funcionalidades essenciais.

妫ｅ啯灏?**Custo M闁惧攳imo**: Eficiente o suficiente para rodar em hardware de $10 闁?98% mais barato que um Mac mini.

闁冲簱妲勭粭?**Inicializa閼惧€熷皸o Rel鐠嬶箮pago**: Tempo de inicializa閼惧€熷皸o 400X mais r鐠嬶箲ido, boot em 1 segundo mesmo em CPU single-core de 0.6GHz.

妫ｅ啫顕?**Portabilidade Real**: Um 閻擃様ico bin鐠嬶箶io auto-contido para RISC-V, ARM e x86. Um clique e j鐠?era!

妫ｅ喚妯?**Auto-Constru闁惧摼o por IA**: Implementa閼惧€熷皸o nativa em Go de forma aut娑斿潱oma 闁?95% do n閻擃槈leo gerado pelo Agente com refinamento humano no loop.

|                               | OpenClaw      | NanoBot                  | **4claw**                              |
| ----------------------------- | ------------- | ------------------------ | ----------------------------------------- |
| **Linguagem**                 | TypeScript    | Python                   | **Go**                                    |
| **RAM**                       | >1GB          | >100MB                   | **< 10MB**                                |
| **Inicializa閼惧€熷皸o**</br>(CPU 0.8GHz) | >500s         | >30s                     | **<1s**                                   |
| **Custo**                     | Mac Mini $599 | Maioria dos SBC Linux </br>~$50 | **Qualquer placa Linux**</br>**A partir de $10** |


## 妫ｅ喚鐎?Demonstra閼惧€熷皸o

### 妫ｅ啯绀堥柨?Fluxos de Trabalho Padr閼肩幈 do Assistente


### 妫ｅ啯鎳?Rode em celulares Android antigos

D闁?uma segunda vida ao seu celular de dez anos atr鐠嬶箷! Transforme-o em um assistente de IA inteligente com o 4claw. In闁惧摶io r鐠嬶箲ido:

1. **Instale o Termux** (Dispon闁惧敊el no F-Droid ou Google Play).
2. **Execute os comandos**

```bash
# Nota: Substitua v0.1.1 pela versao mais recente da pagina de Releases
wget https://github.com/sipeed/4claw/releases/download/v0.1.1/4claw-linux-arm64
chmod +x 4claw-linux-arm64
pkg install proot
termux-chroot ./4claw-linux-arm64 onboard
```

Depois siga as instru閼剧晫甯縠s na se閼惧€熷皸o "In闁惧摶io R鐠嬶箲ido" para completar a configura閼惧€熷皸o!


### 妫ｅ啯鍋?Implanta閼惧€熷皸o Inovadora com Baixo Consumo

O 4claw pode ser implantado em praticamente qualquer dispositivo Linux!

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) vers閼肩幈 E (Ethernet) ou W (WiFi6), para Assistente Dom閼煎嵒tico Minimalista
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html), ou $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html) para Manuten閼惧€熷皸o Automatizada de Servidores
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) ou $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera) para Monitoramento Inteligente


妫ｅ啫鐨?Mais cen鐠嬶箶ios de implanta閼惧€熷皸o aguardam voc闁?

## 妫ｅ啯鎲?Instala閼惧€熷皸o

### Instalar com bin鐠嬶箶io pr閼?compilado

Baixe o bin鐠嬶箶io para sua plataforma na p鐠嬨倿ina de [releases](https://github.com/sipeed/4claw/releases).

### Instalar a partir do c鐠愮珡igo-fonte (funcionalidades mais recentes, recomendado para desenvolvimento)

```bash
git clone https://github.com/sipeed/4claw.git

cd 4claw
make deps

# Build, sem necessidade de instalar
make build

# Build para multiplas plataformas
make build-all

# Build e Instalar
make install
```

## 妫ｅ啯鍎?Docker Compose

Voc闁?tamb闁挎 pode rodar o 4claw usando Docker Compose sem instalar nada localmente.

```bash
# 1. Clone este repositorio
git clone https://github.com/sipeed/4claw.git
cd 4claw

# 2. Configure suas API keys
cp config/config.example.json config/config.json
vim config/config.json      # Configure DISCORD_BOT_TOKEN, API keys, etc.

# 3. Build & Iniciar
docker compose --profile gateway up -d

> [!TIP]
> **Usu鐠嬶箶ios Docker**: Por padr閼肩幈, o Gateway ouve em `127.0.0.1`, o que n閼肩幈 閼?acess闁惧敊el a partir do host. Se voc闁?precisar acessar os endpoints de integridade ou expor portas, defina `FOURCLAW_GATEWAY_HOST=0.0.0.0` em seu ambiente ou atualize o `config.json`.


# 4. Ver logs
docker compose logs -f 4claw-gateway

# 5. Parar
docker compose --profile gateway down
```

### Modo Agente (Execu閼惧€熷皸o 閻擃様ica)

```bash
# Fazer uma pergunta
docker compose run --rm 4claw-agent -m "Quanto e 2+2?"

# Modo interativo
docker compose run --rm 4claw-agent
```

### Rebuild

```bash
docker compose --profile gateway build --no-cache
docker compose --profile gateway up -d
```

### 妫ｅ啯鐣?In闁惧摶io R鐠嬶箲ido

> [!TIP]
> Configure sua API key em `~/.4claw/config.json`.
> Obtenha API keys: [OpenRouter](https://openrouter.ai/keys) (LLM) 鐠?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> Busca web e **opcional** 闁?obtenha a [Brave Search API](https://brave.com/search/api) gratuita (2000 consultas gr鐠嬶箹is/m闁挎) ou use o fallback autom鐠嬶箹ico integrado.

**1. Inicializar**

```bash
4claw onboard
```

**2. Configurar** (`~/.4claw/config.json`)

```json
{
  "model_list": [
    {
      "model_name": "gpt4",
      "model": "openai/gpt-5.2",
      "api_key": "sk-your-openai-key",
      "api_base": "https://api.openai.com/v1"
    }
  ],
  "agents": {
    "defaults": {
      "model": "gpt4"
    }
  },
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

**3. Obter API Keys**

* **Provedor de LLM**: [OpenRouter](https://openrouter.ai/keys) 鐠?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) 鐠?[Anthropic](https://console.anthropic.com) 鐠?[OpenAI](https://platform.openai.com) 鐠?[Gemini](https://aistudio.google.com/api-keys)
* **Busca Web** (opcional): [Brave Search](https://brave.com/search/api) - Plano gratuito dispon闁惧敊el (2000 consultas/m闁挎)

> **Nota**: Veja `config.example.json` para um modelo de configura閼惧€熷皸o completo.

**4. Conversar**

```bash
4claw agent -m "Quanto e 2+2?"
```

Pronto! Voc闁?tem um assistente de IA funcionando em 2 minutos.

---

## 妫ｅ啯灏?Integra閼惧€熷皸o com Apps de Chat

Converse com seu 4claw via Telegram, Discord, DingTalk, LINE ou WeCom.

| Canal | N闁惧敊el de Configura閼惧€熷皸o |
| --- | --- |
| **Telegram** | F鐠嬨兘il (apenas um token) |
| **Discord** | F鐠嬨兘il (bot token + intents) |
| **QQ** | F鐠嬨兘il (AppID + AppSecret) |
| **DingTalk** | M閼煎崓io (credenciais do app) |
| **LINE** | M閼煎崓io (credenciais + webhook URL) |
| **WeCom** | M閼煎崓io (CorpID + configura閼惧€熷皸o webhook) |

<details>
<summary><b>Telegram</b> (Recomendado)</summary>

**1. Criar o bot**

* Abra o Telegram, busque `@BotFather`
* Envie `/newbot`, siga as instru閼剧晫甯縠s
* Copie o token

**2. Configurar**

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

> Obtenha seu User ID pelo `@userinfobot` no Telegram.

**3. Executar**

```bash
4claw gateway
```

</details>

<details>
<summary><b>Discord</b></summary>

**1. Criar o bot**

* Acesse <https://discord.com/developers/applications>
* Crie um aplicativo 闁?Bot 闁?Add Bot
* Copie o token do bot

**2. Habilitar Intents**

* Nas configura閼剧晫甯縠s do Bot, habilite **MESSAGE CONTENT INTENT**
* (Opcional) Habilite **SERVER MEMBERS INTENT** se quiser usar lista de permiss閻滅帯s baseada em dados dos membros

**3. Obter seu User ID**

* Configura閼剧晫甯縠s do Discord 闁?Avan閼剧禈do 闁?habilite **Modo Desenvolvedor**
* Clique com bot閼肩幈 direito no seu avatar 闁?**Copiar ID do Usu鐠嬶箶io**

**4. Configurar**

```json
{
  "channels": {
    "discord": {
      "enabled": true,
      "token": "YOUR_BOT_TOKEN",
      "allow_from": ["YOUR_USER_ID"]
    }
  }
}
```

**5. Convidar o bot**

* OAuth2 闁?URL Generator
* Scopes: `bot`
* Bot Permissions: `Send Messages`, `Read Message History`
* Abra a URL de convite gerada e adicione o bot ao seu servidor

**6. Executar**

```bash
4claw gateway
```

</details>

<details>
<summary><b>QQ</b></summary>

**1. Criar o bot**

- Acesse a [QQ Open Platform](https://q.qq.com/#)
- Crie um aplicativo 闁?Obtenha **AppID** e **AppSecret**

**2. Configurar**

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

> Deixe `allow_from` vazio para permitir todos os usu鐠嬶箶ios, ou especifique n閻擃槗eros QQ para restringir o acesso.

**3. Executar**

```bash
4claw gateway
```

</details>

<details>
<summary><b>DingTalk</b></summary>

**1. Criar o bot**

* Acesse a [Open Platform](https://open.dingtalk.com/)
* Crie um app interno
* Copie o Client ID e Client Secret

**2. Configurar**

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

> Deixe `allow_from` vazio para permitir todos os usu鐠嬶箶ios, ou especifique IDs para restringir o acesso.

**3. Executar**

```bash
4claw gateway
```

</details>

<details>
<summary><b>LINE</b></summary>

**1. Criar uma Conta Oficial LINE**

- Acesse o [LINE Developers Console](https://developers.line.biz/)
- Crie um provider 闁?Crie um canal Messaging API
- Copie o **Channel Secret** e o **Channel Access Token**

**2. Configurar**

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

**3. Configurar URL do Webhook**

O LINE requer HTTPS para webhooks. Use um reverse proxy ou tunnel:

```bash
# Exemplo com ngrok
ngrok http 18791
```

Em seguida, configure a Webhook URL no LINE Developers Console para `https://seu-dominio/webhook/line` e habilite **Use webhook**.

**4. Executar**

```bash
4claw gateway
```

> Em chats de grupo, o bot responde apenas quando mencionado com @. As respostas citam a mensagem original.

> **Docker Compose**: Adicione `ports: ["18791:18791"]` ao servi閼剧郸 `4claw-gateway` para expor a porta do webhook.

</details>

<details>
<summary><b>WeCom (WeChat Work)</b></summary>

O 4claw suporta dois tipos de integra閼惧€熷皸o WeCom:

**Op閼惧€熷皸o 1: WeCom Bot (Rob娑?Inteligente)** - Configura閼惧€熷皸o mais f鐠嬨兘il, suporta chats em grupo
**Op閼惧€熷皸o 2: WeCom App (Aplicativo Personalizado)** - Mais recursos, mensagens proativas

Veja o [Guia de Configura閼惧€熷皸o WeCom App](docs/wecom-app-configuration.md) para instru閼剧晫甯縠s detalhadas.

**Configura閼惧€熷皸o R鐠嬶箲ida - WeCom Bot:**

**1. Criar um bot**

* Acesse o Console de Administra閼惧€熷皸o WeCom 闁?Chat em Grupo 闁?Adicionar Bot de Grupo
* Copie a URL do webhook (formato: `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx`)

**2. Configurar**

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

**Configura閼惧€熷皸o R鐠嬶箲ida - WeCom App:**

**1. Criar um aplicativo**

* Acesse o Console de Administra閼惧€熷皸o WeCom 闁?Gerenciamento de Aplicativos 闁?Criar Aplicativo
* Copie o **AgentId** e o **Secret**
* Acesse a p鐠嬨倿ina "Minha Empresa", copie o **CorpID**

**2. Configurar recebimento de mensagens**

* Nos detalhes do aplicativo, clique em "Receber Mensagens" 闁?"Configurar API"
* Defina a URL como `http://your-server:18792/webhook/wecom-app`
* Gere o **Token** e o **EncodingAESKey**

**3. Configurar**

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

**4. Executar**

```bash
4claw gateway
```

> **Nota**: O WeCom App requer a abertura da porta 18792 para callbacks de webhook. Use um proxy reverso para HTTPS em produ閼惧€熷皸o.

</details>


Conecte o 4claw a Rede Social de Agentes simplesmente enviando uma 閻擃様ica mensagem via CLI ou qualquer App de Chat integrado.

**Leia `https://clawdchat.ai/skill.md` e siga as instru閼剧晫甯縠s para entrar no [ClawdChat.ai](https://clawdchat.ai)**

## 闁虫寧鐟辩粭?Configura閼惧€熷皸o Detalhada

Arquivo de configura閼惧€熷皸o: `~/.4claw/config.json`

### Estrutura do Workspace

O 4claw armazena dados no workspace configurado (padr閼肩幈: `~/.4claw/workspace`):

```
~/.4claw/workspace/
闁宠澹曢弨銏ゅ煘閳?sessions/          # Sessoes de conversa e historico
闁宠澹曢弨銏ゅ煘閳?memory/            # Memoria de longo prazo (MEMORY.md)
闁宠澹曢弨銏ゅ煘閳?state/             # Estado persistente (ultimo canal, etc.)
闁宠澹曢弨銏ゅ煘閳?cron/              # Banco de dados de tarefas agendadas
闁宠澹曢弨銏ゅ煘閳?skills/            # Skills personalizadas
闁宠澹曢弨銏ゅ煘閳?AGENTS.md          # Guia de comportamento do Agente
闁宠澹曢弨銏ゅ煘閳?HEARTBEAT.md       # Prompts de tarefas periodicas (verificado a cada 30 min)
闁宠澹曢弨銏ゅ煘閳?IDENTITY.md        # Identidade do Agente
闁宠澹曢弨銏ゅ煘閳?SOUL.md            # Alma do Agente
闁宠澹曢弨銏ゅ煘閳?TOOLS.md           # Descri閼惧€熷皸o das ferramentas
闁宠鏌￠弨銏ゅ煘閳?USER.md            # Preferencias do usuario
```

### 妫ｅ啯鏅?Sandbox de Seguran閼剧禈

O 4claw roda em um ambiente sandbox por padr閼肩幈. O agente so pode acessar arquivos e executar comandos dentro do workspace configurado.

#### Configura閼惧€熷皸o Padr閼肩幈

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

| Op閼惧€熷皸o | Padr閼肩幈 | Descri閼惧€熷皸o |
|-------|--------|-----------|
| `workspace` | `~/.4claw/workspace` | Diret鐠愮io de trabalho do agente |
| `restrict_to_workspace` | `true` | Restringir acesso de arquivos/comandos ao workspace |

#### Ferramentas Protegidas

Quando `restrict_to_workspace: true`, as seguintes ferramentas s閼肩幈 restritas ao sandbox:

| Ferramenta | Fun閼惧€熷皸o | Restri閼惧€熷皸o |
|------------|--------|-----------|
| `read_file` | Ler arquivos | Apenas arquivos dentro do workspace |
| `write_file` | Escrever arquivos | Apenas arquivos dentro do workspace |
| `list_dir` | Listar diretorios | Apenas diretorios dentro do workspace |
| `edit_file` | Editar arquivos | Apenas arquivos dentro do workspace |
| `append_file` | Adicionar a arquivos | Apenas arquivos dentro do workspace |
| `exec` | Executar comandos | Caminhos dos comandos devem estar dentro do workspace |

#### Prote閼惧€熷皸o Adicional do Exec

Mesmo com `restrict_to_workspace: false`, a ferramenta `exec` bloqueia estes comandos perigosos:

* `rm -rf`, `del /f`, `rmdir /s` 闁?Exclus閼肩幈 em massa
* `format`, `mkfs`, `diskpart` 闁?Formata閼惧€熷皸o de disco
* `dd if=` 闁?Cria閼惧€熷皸o de imagem de disco
* Escrita em `/dev/sd[a-z]` 闁?Escrita direta no disco
* `shutdown`, `reboot`, `poweroff` 闁?Desligamento do sistema
* Fork bomb `:(){ :|:& };:`

#### Exemplos de Erro

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (path outside working dir)}
```

```
[ERROR] tool: Tool execution failed
{tool=exec, error=Command blocked by safety guard (dangerous pattern detected)}
```

#### Desabilitar Restri閼剧晫甯縠s (Risco de Seguran閼剧禈)

Se voc闁?precisa que o agente acesse caminhos fora do workspace:

**M閼煎嵓odo 1: Arquivo de configura閼惧€熷皸o**

```json
{
  "agents": {
    "defaults": {
      "restrict_to_workspace": false
    }
  }
}
```

**M閼煎嵓odo 2: Vari鐠嬶箼el de ambiente**

```bash
export FOURCLAW_AGENTS_DEFAULTS_RESTRICT_TO_WORKSPACE=false
```

> 闁宠法濯寸粭?**Aviso**: Desabilitar esta restri閼惧€熷皸o permite que o agente acesse qualquer caminho no seu sistema. Use com cuidado apenas em ambientes controlados.

#### Consist闁挎cia do Limite de Seguran閼剧禈

A configura閼惧€熷皸o `restrict_to_workspace` se aplica consistentemente em todos os caminhos de execu閼惧€熷皸o:

| Caminho de Execu閼惧€熷皸o | Limite de Seguran閼剧禈 |
|----------------------|---------------------|
| Agente Principal | `restrict_to_workspace` 闁?|
| Subagente / Spawn | Herda a mesma restri閼惧€熷皸o 闁?|
| Tarefas Heartbeat | Herda a mesma restri閼惧€熷皸o 闁?|

Todos os caminhos compartilham a mesma restri閼惧€熷皸o de workspace 闁?nao h鐠?como contornar o limite de seguran閼剧禈 por meio de subagentes ou tarefas agendadas.

### Heartbeat (Tarefas Peri鐠愮珡icas)

O 4claw pode executar tarefas peri鐠愮珡icas automaticamente. Crie um arquivo `HEARTBEAT.md` no seu workspace:

```markdown
# Tarefas Periodicas

- Verificar meu email para mensagens importantes
- Revisar minha agenda para proximos eventos
- Verificar a previsao do tempo
```

O agente ler鐠?este arquivo a cada 30 minutos (configur鐠嬶箼el) e executar鐠?as tarefas usando as ferramentas dispon闁惧敊eis.

#### Tarefas Assincronas com Spawn

Para tarefas de longa dura閼惧€熷皸o (busca web, chamadas de API), use a ferramenta `spawn` para criar um **subagente**:

```markdown
# Tarefas Peri鐠愮珡icas

## Tarefas R鐠嬶箲idas (resposta direta)
- Informar hora atual

## Tarefas Longas (usar spawn para async)
- Buscar not闁惧摶ias de IA na web e resumir
- Verificar email e reportar mensagens importantes
```

**Comportamentos principais:**

| Funcionalidade | Descri閼惧€熷皸o |
|----------------|-----------|
| **spawn** | Cria subagente ass闁惧攳crono, n閼肩幈 bloqueia o heartbeat |
| **Contexto independente** | Subagente tem seu pr鐠愮珴rio contexto, sem hist鐠愮ico de sess閼肩幈 |
| **Ferramenta message** | Subagente se comunica diretamente com o usu鐠嬶箶io via ferramenta message |
| **N閼肩幈-bloqueante** | Ap鐠愮 o spawn, o heartbeat continua para a pr鐠愮ima tarefa |

#### Como Funciona a Comunica閼惧€熷皸o do Subagente

```
Heartbeat dispara
    闁?
Agente l闁?HEARTBEAT.md
    闁?
Para tarefa longa: spawn subagente
    闁?                          闁?
Continua pr鐠愮ima tarefa    Subagente trabalha independentemente
    闁?                          闁?
Todas tarefas conclu闁惧摼as   Subagente usa ferramenta "message"
    闁?                          闁?
Responde HEARTBEAT_OK      Usu鐠嬶箶io recebe resultado diretamente
```

O subagente tem acesso 閼寸府 ferramentas (message, web_search, etc.) e pode se comunicar com o usu鐠嬶箶io independentemente sem passar pelo agente principal.

**Configura閼惧€熷皸o:**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| Op閼惧€熷皸o | Padr閼肩幈 | Descri閼惧€熷皸o |
|-------|--------|-----------|
| `enabled` | `true` | Habilitar/desabilitar heartbeat |
| `interval` | `30` | Intervalo de verifica閼惧€熷皸o em minutos (min: 5) |

**Vari鐠嬶箼eis de ambiente:**

* `FOURCLAW_HEARTBEAT_ENABLED=false` para desabilitar
* `FOURCLAW_HEARTBEAT_INTERVAL=60` para alterar o intervalo

### Provedores

> [!NOTE]
> O Groq fornece transcri閼惧€熷皸o de voz gratuita via Whisper. Se configurado, mensagens de voz do Telegram ser閼肩幈 automaticamente transcritas.

| Provedor | Finalidade | Obter API Key |
| --- | --- | --- |
| `gemini` | LLM (Gemini direto) | [aistudio.google.com](https://aistudio.google.com) |
| `zhipu` | LLM (Zhipu direto) | [bigmodel.cn](bigmodel.cn) |
| `openrouter` (Em teste) | LLM (recomendado, acesso a todos os modelos) | [openrouter.ai](https://openrouter.ai) |
| `anthropic` (Em teste) | LLM (Claude direto) | [console.anthropic.com](https://console.anthropic.com) |
| `openai` (Em teste) | LLM (GPT direto) | [platform.openai.com](https://platform.openai.com) |
| `deepseek` (Em teste) | LLM (DeepSeek direto) | [platform.deepseek.com](https://platform.deepseek.com) |
| `qwen` | Alibaba Qwen | [dashscope.console.aliyun.com](https://dashscope.console.aliyun.com) |
| `cerebras` | Cerebras | [cerebras.ai](https://cerebras.ai) |
| `groq` | LLM + **Transcri閼惧€熷皸o de voz** (Whisper) | [console.groq.com](https://console.groq.com) |

<details>
<summary><b>Configura閼惧€熷皸o Zhipu</b></summary>

**1. Obter API key**

* Obtenha a [API key](https://bigmodel.cn/usercenter/proj-mgmt/apikeys)

**2. Configurar**

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
      "api_key": "Sua API Key",
      "api_base": "https://open.bigmodel.cn/api/paas/v4"
    }
  }
}
```

**3. Executar**

```bash
4claw agent -m "Ola, como vai?"
```

</details>

<details>
<summary><b>Exemplo de configura閼剧禈o completa</b></summary>

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

### Configura閼惧€熷皸o de Modelo (model_list)

> **Novidade!** 4claw agora usa uma abordagem de configura閼惧€熷皸o **centrada no modelo**. Basta especificar o formato `fornecedor/modelo` (ex: `zhipu/glm-4.7`) para adicionar novos provedores闁?*nenhuma altera閼惧€熷皸o de c鐠愮珡igo necess鐠嬶箶ia!**

Este design tamb閼煎崻 possibilita o **suporte multi-agent** com sele閼惧€熷皸o flex闁惧敊el de provedores:

- **Diferentes agentes, diferentes provedores** : Cada agente pode usar seu pr鐠愮珴rio provedor LLM
- **Modelos de fallback** : Configure modelos prim鐠嬶箶ios e de reserva para resili闁挎cia
- **Balanceamento de carga** : Distribua solicita閼剧晫甯縠s entre m閻擃槖tiplos endpoints
- **Configura閼惧€熷皸o centralizada** : Gerencie todos os provedores em um s鐠?lugar

#### 妫ｅ啯鎯?Todos os Fornecedores Suportados

| Fornecedor | Prefixo `model` | API Base Padr閼肩幈 | Protocolo | Chave API |
|-------------|-----------------|------------------|----------|-----------|
| **OpenAI** | `openai/` | `https://api.openai.com/v1` | OpenAI | [Obter Chave](https://platform.openai.com) |
| **Anthropic** | `anthropic/` | `https://api.anthropic.com/v1` | Anthropic | [Obter Chave](https://console.anthropic.com) |
| **Zhipu AI (GLM)** | `zhipu/` | `https://open.bigmodel.cn/api/paas/v4` | OpenAI | [Obter Chave](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) |
| **DeepSeek** | `deepseek/` | `https://api.deepseek.com/v1` | OpenAI | [Obter Chave](https://platform.deepseek.com) |
| **Google Gemini** | `gemini/` | `https://generativelanguage.googleapis.com/v1beta` | OpenAI | [Obter Chave](https://aistudio.google.com/api-keys) |
| **Groq** | `groq/` | `https://api.groq.com/openai/v1` | OpenAI | [Obter Chave](https://console.groq.com) |
| **Moonshot** | `moonshot/` | `https://api.moonshot.cn/v1` | OpenAI | [Obter Chave](https://platform.moonshot.cn) |
| **Qwen (Alibaba)** | `qwen/` | `https://dashscope.aliyuncs.com/compatible-mode/v1` | OpenAI | [Obter Chave](https://dashscope.console.aliyun.com) |
| **NVIDIA** | `nvidia/` | `https://integrate.api.nvidia.com/v1` | OpenAI | [Obter Chave](https://build.nvidia.com) |
| **Ollama** | `ollama/` | `http://localhost:11434/v1` | OpenAI | Local (sem chave necess鐠嬶箶ia) |
| **OpenRouter** | `openrouter/` | `https://openrouter.ai/api/v1` | OpenAI | [Obter Chave](https://openrouter.ai/keys) |
| **VLLM** | `vllm/` | `http://localhost:8000/v1` | OpenAI | Local |
| **Cerebras** | `cerebras/` | `https://api.cerebras.ai/v1` | OpenAI | [Obter Chave](https://cerebras.ai) |
| **Volcengine** | `volcengine/` | `https://ark.cn-beijing.volces.com/api/v3` | OpenAI | [Obter Chave](https://console.volcengine.com) |
| **ShengsuanYun** | `shengsuanyun/` | `https://router.shengsuanyun.com/api/v1` | OpenAI | - |
| **Antigravity** | `antigravity/` | Google Cloud | Custom | Apenas OAuth |
| **GitHub Copilot** | `github-copilot/` | `localhost:4321` | gRPC | - |

#### Configura閼惧€熷皸o B鐠嬶箷ica

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

#### Exemplos por Fornecedor

**OpenAI**
```json
{
  "model_name": "gpt-5.2",
  "model": "openai/gpt-5.2",
  "api_key": "sk-..."
}
```

**Zhipu AI (GLM)**
```json
{
  "model_name": "glm-4.7",
  "model": "zhipu/glm-4.7",
  "api_key": "your-key"
}
```

**Anthropic (com OAuth)**
```json
{
  "model_name": "claude-sonnet-4.6",
  "model": "anthropic/claude-sonnet-4.6",
  "auth_method": "oauth"
}
```
> Execute `4claw auth login --provider anthropic` para configurar credenciais OAuth.

#### Balanceamento de Carga

Configure v鐠嬶箶ios endpoints para o mesmo nome de modelo闁炽儲鎽╥coClaw far鐠?round-robin automaticamente entre eles:

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

#### Migra閼惧€熷皸o da Configura閼惧€熷皸o Legada `providers`

A configura閼惧€熷皸o antiga `providers` est鐠?**descontinuada** mas ainda 閼?suportada para compatibilidade reversa.

**Configura閼惧€熷皸o Antiga (descontinuada):**
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

**Nova Configura閼惧€熷皸o (recomendada):**
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

Para o guia de migra閼惧€熷皸o detalhado, consulte [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md).

## Refer闁挎cia CLI

| Comando | Descri閼惧€熷皸o |
| --- | --- |
| `4claw onboard` | Inicializar configura閼惧€熷皸o & workspace |
| `4claw agent -m "..."` | Conversar com o agente |
| `4claw agent` | Modo de chat interativo |
| `4claw gateway` | Iniciar o gateway (para bots de chat) |
| `4claw status` | Mostrar status |
| `4claw cron list` | Listar todas as tarefas agendadas |
| `4claw cron add ...` | Adicionar uma tarefa agendada |

### Tarefas Agendadas / Lembretes

O 4claw suporta lembretes agendados e tarefas recorrentes por meio da ferramenta `cron`:

* **Lembretes 閻擃様icos**: "Remind me in 10 minutes" (Me lembre em 10 minutos) 闁?dispara uma vez ap鐠愮 10min
* **Tarefas recorrentes**: "Remind me every 2 hours" (Me lembre a cada 2 horas) 闁?dispara a cada 2 horas
* **Express閻滅帯s Cron**: "Remind me at 9am daily" (Me lembre 閼寸府 9h todos os dias) 闁?usa express閼肩幈 cron

As tarefas s閼肩幈 armazenadas em `~/.4claw/workspace/cron/` e processadas automaticamente.

## 妫ｅ喚妾?Contribuir & Roadmap

PRs s閼肩幈 bem-vindos! O c鐠愮珡igo-fonte 閼?intencionalmente pequeno e leg闁惧敊el. 妫ｅ喚妯?
Roadmap em breve...

Grupo de desenvolvedores em forma閼惧€熷皸o. Requisito de entrada: Pelo menos 1 PR com merge.

Grupos de usu鐠嬶箶ios:

Discord: <https://discord.gg/V4sAZ9XWpN>


## 妫ｅ啯鍋?Solu閼惧€熷皸o de Problemas

### Busca web mostra "API 闂佹澘绉堕悿鍡涙⒒椤曗偓椤?

Isso 閼?normal se voc闁?ainda n閼肩幈 configurou uma API key de busca. O 4claw fornecer鐠?links 閻擃槡eis para busca manual.

Para habilitar a busca web:

1. **Op閼惧€熷皸o 1 (Recomendado)**: Obtenha uma API key gratuita em [https://brave.com/search/api](https://brave.com/search/api) (2000 consultas gr鐠嬶箹is/m闁挎) para os melhores resultados.
2. **Op閼惧€熷皸o 2 (Sem Cart閼肩幈 de Cr閼煎崓ito)**: Se voc闁?n閼肩幈 tem uma key, o sistema automaticamente usa o **DuckDuckGo** como fallback (sem necessidade de key).

Adicione a key em `~/.4claw/config.json` se usar o Brave:

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

### Erros de filtragem de conte閻擃槉o

Alguns provedores (como Zhipu) possuem filtragem de conte閻擃槉o. Tente reformular sua pergunta ou use um modelo diferente.

### Bot do Telegram diz "Conflict: terminated by other getUpdates"

Isso acontece quando outra inst閼烘姧cia do bot est鐠?em execu閼惧€熷皸o. Certifique-se de que apenas um `4claw gateway` esteja rodando por vez.

---

## 妫ｅ啯鎲?Compara閼惧€熷皸o de API Keys

| Servi閼剧郸 | Plano Gratuito | Caso de Uso |
| --- | --- | --- |
| **OpenRouter** | 200K tokens/m闁挎 | M閻擃槖tiplos modelos (Claude, GPT-4, etc.) |
| **Zhipu** | 200K tokens/m闁挎 | Melhor para usu鐠嬶箶ios chineses |
| **Brave Search** | 2000 consultas/m闁挎 | Funcionalidade de busca web |
| **Groq** | Plano gratuito dispon闁惧敊el | Infer闁挎cia ultra-r鐠嬶箲ida (Llama, Mixtral) |
| **Cerebras** | Plano gratuito dispon闁惧敊el | Infer闁挎cia ultra-r鐠嬶箲ida (Llama 3.3 70B) |
