<div align="center">

<h1>4claw: Assistente de IA Ultra-Eficiente em Go</h1>

<h3>Hardware de $10 闁?10MB de RAM 闁?Boot em 1s 闁?闂備焦妞块悡鍫ュ窗濮橆儷鎺撶節閸ャ劍鍎梺瑙勫劤閹虫劗绮堟径鎰厵閻庢稒锚閻忥絾绻濇繝鍐ㄧ仼闁圭鍛殰妞ゆ柨澧介敍?/h3>


 [濠电偞鍨堕幖鈺呭储娴犲鍑犻柛?README.zh.md) | [闂備礁鎼崯銊╁磿閹绘帪鑰块柟娈垮枟閸ゆ棃鏌?README.ja.md) | **Portugu闂備焦瀵ч?* | [Ti閻庡厜鍋撳瀣凹閸掑ゼ Vi閻庣懓瀚€氼剙顬婄化?README.vi.md) | [Fran闂佺厧澧界粋鍧晄](README.fr.md) | [English](README.md)
</div>

---

婵☆偓绲介崰姘归埀?**4claw** 闂?um assistente pessoal de IA ultra-leve inspirado no 4claw, reescrito do zero em **Go** por meio de um processo de "auto-inicializa闂佸吋鍎抽埀顒傚枎閻ㄧ珳" (self-bootstrapping) 闂?onde o pr闁荤姵鍔楅悵纾杋o agente de IA conduziu toda a migra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de arquitetura e otimiza闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de c闁荤姵鍔楅悵顡痝o.

闂備礁鍟跨花鍗炍ｉ崟顓犵當?**Extremamente leve:** Roda em hardware de apenas **$10** com **<10MB** de RAM. Isso 闂?99% menos mem闁荤姵鍔楅鐒沘 que o OpenClaw e 98% mais barato que um Mac mini!


> [!CAUTION]
> **婵☆偓绲介崯顖炴儊?DECLARA闂佺厧顕崰鏇㈠礈缁?DE SEGURAN闂佺厧顕崕?& CANAIS OFICIAIS**
>
> * **SEM CRIPTOMOEDAS:** O 4claw **N闂佺厧鐡ㄩ埞?* possui nenhum token/moeda oficial. Todas as alega闂佺厧澧介弲顐ゆ暜缁虹垙 no `pump.fun` ou outras plataformas de negocia闂佸吋鍎抽埀顒傚枎閻ㄧ珳 s闂佽壈鍋愰獮?**GOLPES**.
> * **DOM闂佸吋鎯屽缍 OFICIAL:** O **闂佺厧顕锤ICO** site oficial 闂?o **[4claw.io](https://4claw.io)**, e o site da empresa 闂?o **[4claw.io](https://4claw.io)**.
> * **Aviso:** Muitos dom闂備焦鍎抽弨鐮硂s `.ai/.org/.com/.net/...` foram registrados por terceiros, n闂佽壈鍋愰獮?s闂佽壈鍋愰獮?nossos.
> * **Aviso:** O 4claw est闁?em fase inicial de desenvolvimento e pode ter problemas de seguran闂佺厧澧界粋?de rede n闂佽壈鍋愰獮?resolvidos. N闂佽壈鍋愰獮?implante em ambientes de produ闂佸吋鍎抽埀顒傚枎閻ㄧ珳 antes da vers闂佽壈鍋愰獮?v1.0.
> * **Nota:** O 4claw recentemente fez merge de muitos PRs, o que pode resultar em maior consumo de mem闁荤姵鍔楅鐒沘 (10-20MB) nas vers闂佺粯绮庣敮鐥?mais recentes. Planejamos priorizar a otimiza闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de recursos assim que o conjunto de funcionalidades estiver est闁荤姴顑戠粻绯磍.


## 婵☆偓绲介崯顖炲箟?Novidades

2026-02-16 婵☆偓绲介崯顐ょ博?4claw atingiu 12K stars em uma semana! Obrigado a todos pelo apoio! O 4claw est闁?crescendo mais r闁荤姴顑戠粻鐬歞o do que jamais imaginamos. Dado o alto volume de PRs, precisamos urgentemente de maintainers da comunidade. Nossos pap闂佽偐鍘у畷鈺?de volunt闁荤姴顑戠粻绉坥s e roadmap foram publicados oficialmente [aqui](docs/ROADMAP.md) 闂?estamos ansiosos para ter voc闂?a bordo!

2026-02-13 婵☆偓绲介崯顐ょ博?4claw atingiu 5000 stars em 4 dias! Obrigado 闂?comunidade! Estamos finalizando o **Roadmap do Projeto** e configurando o **Grupo de Desenvolvedores** para acelerar o desenvolvimento do 4claw.

婵☆偓绲介崯顖炴偩?**Chamada para A闂佸吋鍎抽埀顒傚枎閻ㄧ珳:** Envie suas solicita闂佺厧澧介弲顐ゆ暜缁虹垙 de funcionalidades nas GitHub Discussions. Revisaremos e priorizaremos na pr闁荤姵鍔楅鈺a reuni闂佽壈鍋愰獮?semanal.

2026-02-09 婵☆偓绲介崯顐ょ博?4claw lan闂佺厧澧界粋鍧塷 oficialmente! Constru闂備焦鍎抽幗绱€ em 1 dia para trazer Agentes de IA para hardware de $10 com <10MB de RAM. 婵☆偓绲介崰姘归埀?4claw, Partiu!

## 闂?Funcionalidades

婵☆偓绲介崰姘跺焵?**Ultra-Leve**: Consumo de mem闁荤姵鍔楅鐒沘 <10MB 闂?99% menor que o Clawdbot para funcionalidades essenciais.

婵☆偓绲介崯顖滀焊?**Custo M闂備焦鍎抽弨鐮砿o**: Eficiente o suficiente para rodar em hardware de $10 闂?98% mais barato que um Mac mini.

闂備礁鍟跨花鍗炍ｉ崟顓犵當?**Inicializa闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Rel闁荤姴顑戠粻鐣僡go**: Tempo de inicializa闂佸吋鍎抽埀顒傚枎閻ㄧ珳 400X mais r闁荤姴顑戠粻鐬歞o, boot em 1 segundo mesmo em CPU single-core de 0.6GHz.

婵☆偓绲介崯顐︻敋?**Portabilidade Real**: Um 闂佺粯鎼╁Σ姒燾o bin闁荤姴顑戠粻绉坥 auto-contido para RISC-V, ARM e x86. Um clique e j闁?era!

婵☆偓绲介崰姘?**Auto-Constru闂備焦鍎抽幗绱€ por IA**: Implementa闂佸吋鍎抽埀顒傚枎閻ㄧ珳 nativa em Go de forma aut婵炴垶鏌ㄥ鐪攎a 闂?95% do n闂佺粯鎼╁Σ鍧檈o gerado pelo Agente com refinamento humano no loop.

|                               | OpenClaw      | 4claw                  | **4claw**                              |
| ----------------------------- | ------------- | ------------------------ | ----------------------------------------- |
| **Linguagem**                 | TypeScript    | Python                   | **Go**                                    |
| **RAM**                       | >1GB          | >100MB                   | **< 10MB**                                |
| **Inicializa闂佸吋鍎抽埀顒傚枎閻ㄧ珳**</br>(CPU 0.8GHz) | >500s         | >30s                     | **<1s**                                   |
| **Custo**                     | Mac Mini $599 | Maioria dos SBC Linux </br>~$50 | **Qualquer placa Linux**</br>**A partir de $10** |


## 婵☆偓绲介崰姘舵倵?Demonstra闂佸吋鍎抽埀顒傚枎閻ㄧ珳

### 婵☆偓绲介崯顖滅矆閸儲鐓?Fluxos de Trabalho Padr闂佽壈鍋愰獮?do Assistente


### 婵☆偓绲介崯顖炲箠?Rode em celulares Android antigos

D闂?uma segunda vida ao seu celular de dez anos atr闁荤姴顑戠粻? Transforme-o em um assistente de IA inteligente com o 4claw. In闂備焦鍎抽幗绉坥 r闁荤姴顑戠粻鐬歞o:

1. **Instale o Termux** (Dispon闂備焦鍎抽弫濂簂 no F-Droid ou Google Play).
2. **Execute os comandos**

```bash
# Nota: Substitua v0.1.1 pela versao mais recente da pagina de Releases
wget https://github.com/4claw/4claw/releases/download/v0.1.1/4claw-linux-arm64
chmod +x 4claw-linux-arm64
pkg install proot
termux-chroot ./4claw-linux-arm64 onboard
```

Depois siga as instru闂佺厧澧介弲顐ゆ暜缁虹垙 na se闂佸吋鍎抽埀顒傚枎閻ㄧ珳 "In闂備焦鍎抽幗绉坥 R闁荤姴顑戠粻鐬歞o" para completar a configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳!


### 婵☆偓绲介崯顖炲磻?Implanta闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Inovadora com Baixo Consumo

O 4claw pode ser implantado em praticamente qualquer dispositivo Linux!

- $9.9 [LicheeRV-Nano](https://www.aliexpress.com/item/1005006519668532.html) vers闂佽壈鍋愰獮?E (Ethernet) ou W (WiFi6), para Assistente Dom闂佽偐鍘у畵鎶癷co Minimalista
- $30~50 [NanoKVM](https://www.aliexpress.com/item/1005007369816019.html), ou $100 [NanoKVM-Pro](https://www.aliexpress.com/item/1005010048471263.html) para Manuten闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Automatizada de Servidores
- $50 [MaixCAM](https://www.aliexpress.com/item/1005008053333693.html) ou $100 [MaixCAM2](https://www.kickstarter.com/projects/zepan/maixcam2-build-your-next-gen-4k-ai-camera) para Monitoramento Inteligente


婵☆偓绲介崯顐︽儍?Mais cen闁荤姴顑戠粻绉坥s de implanta闂佸吋鍎抽埀顒傚枎閻ㄧ珳 aguardam voc闂?

## 婵☆偓绲介崯顖炲箟?Instala闂佸吋鍎抽埀顒傚枎閻ㄧ珳

### Instalar com bin闁荤姴顑戠粻绉坥 pr闂?compilado

Baixe o bin闁荤姴顑戠粻绉坥 para sua plataforma na p闁荤姴顑冮崐绺a de [releases](https://github.com/4claw/4claw/releases).

### Instalar a partir do c闁荤姵鍔楅悵顡痝o-fonte (funcionalidades mais recentes, recomendado para desenvolvimento)

```bash
git clone https://github.com/4claw/4claw.git

cd 4claw
make deps

# Build, sem necessidade de instalar
make build

# Build para multiplas plataformas
make build-all

# Build e Instalar
make install
```

## 婵☆偓绲介崯顖炲磿?Docker Compose

Voc闂?tamb闂備焦瀵ч?pode rodar o 4claw usando Docker Compose sem instalar nada localmente.

```bash
# 1. Clone este repositorio
git clone https://github.com/4claw/4claw.git
cd 4claw

# 2. Configure suas API keys
cp config/config.example.json config/config.json
vim config/config.json      # Configure DISCORD_BOT_TOKEN, API keys, etc.

# 3. Build & Iniciar
docker compose --profile gateway up -d

> [!TIP]
> **Usu闁荤姴顑戠粻绉坥s Docker**: Por padr闂佽壈鍋愰獮? o Gateway ouve em `127.0.0.1`, o que n闂佽壈鍋愰獮?闂?acess闂備焦鍎抽弫濂簂 a partir do host. Se voc闂?precisar acessar os endpoints de integridade ou expor portas, defina `FOURCLAW_GATEWAY_HOST=0.0.0.0` em seu ambiente ou atualize o `config.json`.


# 4. Ver logs
docker compose logs -f 4claw-gateway

# 5. Parar
docker compose --profile gateway down
```

### Modo Agente (Execu闂佸吋鍎抽埀顒傚枎閻ㄧ珳 闂佺粯鎼╁Σ姒燾a)

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

### 婵☆偓绲介崯顖炴偩?In闂備焦鍎抽幗绉坥 R闁荤姴顑戠粻鐬歞o

> [!TIP]
> Configure sua API key em `~/.4claw/config.json`.
> Obtenha API keys: [OpenRouter](https://openrouter.ai/keys) (LLM) 闁?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) (LLM)
> Busca web e **opcional** 闂?obtenha a [Brave Search API](https://brave.com/search/api) gratuita (2000 consultas gr闁荤姴顑戠粻绛皊/m闂備焦瀵ч? ou use o fallback autom闁荤姴顑戠粻绛癱o integrado.

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

* **Provedor de LLM**: [OpenRouter](https://openrouter.ai/keys) 闁?[Zhipu](https://open.bigmodel.cn/usercenter/proj-mgmt/apikeys) 闁?[Anthropic](https://console.anthropic.com) 闁?[OpenAI](https://platform.openai.com) 闁?[Gemini](https://aistudio.google.com/api-keys)
* **Busca Web** (opcional): [Brave Search](https://brave.com/search/api) - Plano gratuito dispon闂備焦鍎抽弫濂簂 (2000 consultas/m闂備焦瀵ч?

> **Nota**: Veja `config.example.json` para um modelo de configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 completo.

**4. Conversar**

```bash
4claw agent -m "Quanto e 2+2?"
```

Pronto! Voc闂?tem um assistente de IA funcionando em 2 minutos.

---

## 婵☆偓绲介崯顖滀焊?Integra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 com Apps de Chat

Converse com seu 4claw via Telegram, Discord, DingTalk, LINE ou WeCom.

| Canal | N闂備焦鍎抽弫濂簂 de Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
| --- | --- |
| **Telegram** | F闁荤姴顑冮崗姒爈 (apenas um token) |
| **Discord** | F闁荤姴顑冮崗姒爈 (bot token + intents) |
| **QQ** | F闁荤姴顑冮崗姒爈 (AppID + AppSecret) |
| **DingTalk** | M闂佽偐鍘у畷鎼僶 (credenciais do app) |
| **LINE** | M闂佽偐鍘у畷鎼僶 (credenciais + webhook URL) |
| **WeCom** | M闂佽偐鍘у畷鎼僶 (CorpID + configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 webhook) |

<details>
<summary><b>Telegram</b> (Recomendado)</summary>

**1. Criar o bot**

* Abra o Telegram, busque `@BotFather`
* Envie `/newbot`, siga as instru闂佺厧澧介弲顐ゆ暜缁虹垙
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
* Crie um aplicativo 闂?Bot 闂?Add Bot
* Copie o token do bot

**2. Habilitar Intents**

* Nas configura闂佺厧澧介弲顐ゆ暜缁虹垙 do Bot, habilite **MESSAGE CONTENT INTENT**
* (Opcional) Habilite **SERVER MEMBERS INTENT** se quiser usar lista de permiss闂佺粯绮庣敮鐥?baseada em dados dos membros

**3. Obter seu User ID**

* Configura闂佺厧澧介弲顐ゆ暜缁虹垙 do Discord 闂?Avan闂佺厧澧界粋鍧塷 闂?habilite **Modo Desenvolvedor**
* Clique com bot闂佽壈鍋愰獮?direito no seu avatar 闂?**Copiar ID do Usu闁荤姴顑戠粻绉坥**

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

* OAuth2 闂?URL Generator
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
- Crie um aplicativo 闂?Obtenha **AppID** e **AppSecret**

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

> Deixe `allow_from` vazio para permitir todos os usu闁荤姴顑戠粻绉坥s, ou especifique n闂佺粯鎼╁Σ姊時os QQ para restringir o acesso.

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

> Deixe `allow_from` vazio para permitir todos os usu闁荤姴顑戠粻绉坥s, ou especifique IDs para restringir o acesso.

**3. Executar**

```bash
4claw gateway
```

</details>

<details>
<summary><b>LINE</b></summary>

**1. Criar uma Conta Oficial LINE**

- Acesse o [LINE Developers Console](https://developers.line.biz/)
- Crie um provider 闂?Crie um canal Messaging API
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

> **Docker Compose**: Adicione `ports: ["18791:18791"]` ao servi闂佺厧澧介柈?`4claw-gateway` para expor a porta do webhook.

</details>

<details>
<summary><b>WeCom (WeChat Work)</b></summary>

O 4claw suporta dois tipos de integra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 WeCom:

**Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 1: WeCom Bot (Rob婵?Inteligente)** - Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 mais f闁荤姴顑冮崗姒爈, suporta chats em grupo
**Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 2: WeCom App (Aplicativo Personalizado)** - Mais recursos, mensagens proativas

Veja o [Guia de Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 WeCom App](docs/wecom-app-configuration.md) para instru闂佺厧澧介弲顐ゆ暜缁虹垙 detalhadas.

**Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 R闁荤姴顑戠粻鐬歞a - WeCom Bot:**

**1. Criar um bot**

* Acesse o Console de Administra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 WeCom 闂?Chat em Grupo 闂?Adicionar Bot de Grupo
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

**Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 R闁荤姴顑戠粻鐬歞a - WeCom App:**

**1. Criar um aplicativo**

* Acesse o Console de Administra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 WeCom 闂?Gerenciamento de Aplicativos 闂?Criar Aplicativo
* Copie o **AgentId** e o **Secret**
* Acesse a p闁荤姴顑冮崐绺a "Minha Empresa", copie o **CorpID**

**2. Configurar recebimento de mensagens**

* Nos detalhes do aplicativo, clique em "Receber Mensagens" 闂?"Configurar API"
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

> **Nota**: O WeCom App requer a abertura da porta 18792 para callbacks de webhook. Use um proxy reverso para HTTPS em produ闂佸吋鍎抽埀顒傚枎閻ㄧ珳.

</details>


Conecte o 4claw a Rede Social de Agentes simplesmente enviando uma 闂佺粯鎼╁Σ姒燾a mensagem via CLI ou qualquer App de Chat integrado.

**Leia `https://clawdchat.ai/skill.md` e siga as instru闂佺厧澧介弲顐ゆ暜缁虹垙 para entrar no [ClawdChat.ai](https://clawdchat.ai)**

## 闂備浇娅曠€笛囨偡鏉堚晝绠?Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Detalhada

Arquivo de configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳: `~/.4claw/config.json`

### Estrutura do Workspace

O 4claw armazena dados no workspace configurado (padr闂佽壈鍋愰獮? `~/.4claw/workspace`):

```
~/.4claw/workspace/
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?sessions/          # Sessoes de conversa e historico
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?memory/            # Memoria de longo prazo (MEMORY.md)
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?state/             # Estado persistente (ultimo canal, etc.)
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?cron/              # Banco de dados de tarefas agendadas
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?skills/            # Skills personalizadas
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?AGENTS.md          # Guia de comportamento do Agente
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?HEARTBEAT.md       # Prompts de tarefas periodicas (verificado a cada 30 min)
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?IDENTITY.md        # Identidade do Agente
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?SOUL.md            # Alma do Agente
闂備礁鐤囬～澶嬬珶閺囥垹缁╅柕蹇嬪€曢悡姗€鏌?TOOLS.md           # Descri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 das ferramentas
闂備礁鐤囬～澶愬蓟閿熺姴缁╅柕蹇嬪€曢悡姗€鏌?USER.md            # Preferencias do usuario
```

### 婵☆偓绲介崯顖炲疾?Sandbox de Seguran闂佺厧澧界粋?
O 4claw roda em um ambiente sandbox por padr闂佽壈鍋愰獮? O agente so pode acessar arquivos e executar comandos dentro do workspace configurado.

#### Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Padr闂佽壈鍋愰獮?
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

| Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 | Padr闂佽壈鍋愰獮?| Descri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
|-------|--------|-----------|
| `workspace` | `~/.4claw/workspace` | Diret闁荤姵鍔楅鐒沷 de trabalho do agente |
| `restrict_to_workspace` | `true` | Restringir acesso de arquivos/comandos ao workspace |

#### Ferramentas Protegidas

Quando `restrict_to_workspace: true`, as seguintes ferramentas s闂佽壈鍋愰獮?restritas ao sandbox:

| Ferramenta | Fun闂佸吋鍎抽埀顒傚枎閻ㄧ珳 | Restri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
|------------|--------|-----------|
| `read_file` | Ler arquivos | Apenas arquivos dentro do workspace |
| `write_file` | Escrever arquivos | Apenas arquivos dentro do workspace |
| `list_dir` | Listar diretorios | Apenas diretorios dentro do workspace |
| `edit_file` | Editar arquivos | Apenas arquivos dentro do workspace |
| `append_file` | Adicionar a arquivos | Apenas arquivos dentro do workspace |
| `exec` | Executar comandos | Caminhos dos comandos devem estar dentro do workspace |

#### Prote闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Adicional do Exec

Mesmo com `restrict_to_workspace: false`, a ferramenta `exec` bloqueia estes comandos perigosos:

* `rm -rf`, `del /f`, `rmdir /s` 闂?Exclus闂佽壈鍋愰獮?em massa
* `format`, `mkfs`, `diskpart` 闂?Formata闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de disco
* `dd if=` 闂?Cria闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de imagem de disco
* Escrita em `/dev/sd[a-z]` 闂?Escrita direta no disco
* `shutdown`, `reboot`, `poweroff` 闂?Desligamento do sistema
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

#### Desabilitar Restri闂佺厧澧介弲顐ゆ暜缁虹垙 (Risco de Seguran闂佺厧澧界粋?

Se voc闂?precisa que o agente acesse caminhos fora do workspace:

**M闂佽偐鍘у畵鎼奷o 1: Arquivo de configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳**

```json
{
  "agents": {
    "defaults": {
      "restrict_to_workspace": false
    }
  }
}
```

**M闂佽偐鍘у畵鎼奷o 2: Vari闁荤姴顑戠粻绯磍 de ambiente**

```bash
export FOURCLAW_AGENTS_DEFAULTS_RESTRICT_TO_WORKSPACE=false
```

> 闂備礁鐤囧▔鏇熷垔鐎靛摜绠?**Aviso**: Desabilitar esta restri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 permite que o agente acesse qualquer caminho no seu sistema. Use com cuidado apenas em ambientes controlados.

#### Consist闂備焦瀵ч鈺焛a do Limite de Seguran闂佺厧澧界粋?
A configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 `restrict_to_workspace` se aplica consistentemente em todos os caminhos de execu闂佸吋鍎抽埀顒傚枎閻ㄧ珳:

| Caminho de Execu闂佸吋鍎抽埀顒傚枎閻ㄧ珳 | Limite de Seguran闂佺厧澧界粋?|
|----------------------|---------------------|
| Agente Principal | `restrict_to_workspace` 闂?|
| Subagente / Spawn | Herda a mesma restri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 闂?|
| Tarefas Heartbeat | Herda a mesma restri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 闂?|

Todos os caminhos compartilham a mesma restri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de workspace 闂?nao h闁?como contornar o limite de seguran闂佺厧澧界粋?por meio de subagentes ou tarefas agendadas.

### Heartbeat (Tarefas Peri闁荤姵鍔楅悵顡痗as)

O 4claw pode executar tarefas peri闁荤姵鍔楅悵顡痗as automaticamente. Crie um arquivo `HEARTBEAT.md` no seu workspace:

```markdown
# Tarefas Periodicas

- Verificar meu email para mensagens importantes
- Revisar minha agenda para proximos eventos
- Verificar a previsao do tempo
```

O agente ler闁?este arquivo a cada 30 minutos (configur闁荤姴顑戠粻绯磍) e executar闁?as tarefas usando as ferramentas dispon闂備焦鍎抽弫濂篿s.

#### Tarefas Assincronas com Spawn

Para tarefas de longa dura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 (busca web, chamadas de API), use a ferramenta `spawn` para criar um **subagente**:

```markdown
# Tarefas Peri闁荤姵鍔楅悵顡痗as

## Tarefas R闁荤姴顑戠粻鐬歞as (resposta direta)
- Informar hora atual

## Tarefas Longas (usar spawn para async)
- Buscar not闂備焦鍎抽幗绉坅s de IA na web e resumir
- Verificar email e reportar mensagens importantes
```

**Comportamentos principais:**

| Funcionalidade | Descri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
|----------------|-----------|
| **spawn** | Cria subagente ass闂備焦鍎抽弨鐮猺ono, n闂佽壈鍋愰獮?bloqueia o heartbeat |
| **Contexto independente** | Subagente tem seu pr闁荤姵鍔楅悵纾杋o contexto, sem hist闁荤姵鍔楅鐒沜o de sess闂佽壈鍋愰獮?|
| **Ferramenta message** | Subagente se comunica diretamente com o usu闁荤姴顑戠粻绉坥 via ferramenta message |
| **N闂佽壈鍋愰獮?bloqueante** | Ap闁荤姵鍔楅?o spawn, o heartbeat continua para a pr闁荤姵鍔楅鈺a tarefa |

#### Como Funciona a Comunica闂佸吋鍎抽埀顒傚枎閻ㄧ珳 do Subagente

```
Heartbeat dispara
    闂?
Agente l闂?HEARTBEAT.md
    闂?
Para tarefa longa: spawn subagente
    闂?                          闂?
Continua pr闁荤姵鍔楅鈺a tarefa    Subagente trabalha independentemente
    闂?                          闂?
Todas tarefas conclu闂備焦鍎抽幗绯皊   Subagente usa ferramenta "message"
    闂?                          闂?
Responde HEARTBEAT_OK      Usu闁荤姴顑戠粻绉坥 recebe resultado diretamente
```

O subagente tem acesso 闂佺厧顕惔?ferramentas (message, web_search, etc.) e pode se comunicar com o usu闁荤姴顑戠粻绉坥 independentemente sem passar pelo agente principal.

**Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳:**

```json
{
  "heartbeat": {
    "enabled": true,
    "interval": 30
  }
}
```

| Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 | Padr闂佽壈鍋愰獮?| Descri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
|-------|--------|-----------|
| `enabled` | `true` | Habilitar/desabilitar heartbeat |
| `interval` | `30` | Intervalo de verifica闂佸吋鍎抽埀顒傚枎閻ㄧ珳 em minutos (min: 5) |

**Vari闁荤姴顑戠粻绯磇s de ambiente:**

* `FOURCLAW_HEARTBEAT_ENABLED=false` para desabilitar
* `FOURCLAW_HEARTBEAT_INTERVAL=60` para alterar o intervalo

### Provedores

> [!NOTE]
> O Groq fornece transcri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de voz gratuita via Whisper. Se configurado, mensagens de voz do Telegram ser闂佽壈鍋愰獮?automaticamente transcritas.

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
| `groq` | LLM + **Transcri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de voz** (Whisper) | [console.groq.com](https://console.groq.com) |

<details>
<summary><b>Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Zhipu</b></summary>

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
<summary><b>Exemplo de configura闂佺厧澧界粋鍧?completa</b></summary>

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

### Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de Modelo (model_list)

> **Novidade!** 4claw agora usa uma abordagem de configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 **centrada no modelo**. Basta especificar o formato `fornecedor/modelo` (ex: `zhipu/glm-4.7`) para adicionar novos provedores闂?*nenhuma altera闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de c闁荤姵鍔楅悵顡痝o necess闁荤姴顑戠粻绉坅!**

Este design tamb闂佽偐鍘у畷?possibilita o **suporte multi-agent** com sele闂佸吋鍎抽埀顒傚枎閻ㄧ珳 flex闂備焦鍎抽弫濂簂 de provedores:

- **Diferentes agentes, diferentes provedores** : Cada agente pode usar seu pr闁荤姵鍔楅悵纾杋o provedor LLM
- **Modelos de fallback** : Configure modelos prim闁荤姴顑戠粻绉坥s e de reserva para resili闂備焦瀵ч鈺焛a
- **Balanceamento de carga** : Distribua solicita闂佺厧澧介弲顐ゆ暜缁虹垙 entre m闂佺粯鎼╁Σ鏉plos endpoints
- **Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 centralizada** : Gerencie todos os provedores em um s闁?lugar

#### 婵☆偓绲介崯顖炲箚?Todos os Fornecedores Suportados

| Fornecedor | Prefixo `model` | API Base Padr闂佽壈鍋愰獮?| Protocolo | Chave API |
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
| **Ollama** | `ollama/` | `http://localhost:11434/v1` | OpenAI | Local (sem chave necess闁荤姴顑戠粻绉坅) |
| **OpenRouter** | `openrouter/` | `https://openrouter.ai/api/v1` | OpenAI | [Obter Chave](https://openrouter.ai/keys) |
| **VLLM** | `vllm/` | `http://localhost:8000/v1` | OpenAI | Local |
| **Cerebras** | `cerebras/` | `https://api.cerebras.ai/v1` | OpenAI | [Obter Chave](https://cerebras.ai) |
| **Volcengine** | `volcengine/` | `https://ark.cn-beijing.volces.com/api/v3` | OpenAI | [Obter Chave](https://console.volcengine.com) |
| **ShengsuanYun** | `shengsuanyun/` | `https://router.shengsuanyun.com/api/v1` | OpenAI | - |
| **Antigravity** | `antigravity/` | Google Cloud | Custom | Apenas OAuth |
| **GitHub Copilot** | `github-copilot/` | `localhost:4321` | gRPC | - |

#### Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 B闁荤姴顑戠粻绌抍a

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

Configure v闁荤姴顑戠粻绉坥s endpoints para o mesmo nome de modelo闂備胶鍋ㄩ崕鏌ュ箺閳侯櫓oClaw far闁?round-robin automaticamente entre eles:

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

#### Migra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 da Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Legada `providers`

A configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 antiga `providers` est闁?**descontinuada** mas ainda 闂?suportada para compatibilidade reversa.

**Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 Antiga (descontinuada):**
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

**Nova Configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 (recomendada):**
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

Para o guia de migra闂佸吋鍎抽埀顒傚枎閻ㄧ珳 detalhado, consulte [docs/migration/model-list-migration.md](docs/migration/model-list-migration.md).

## Refer闂備焦瀵ч鈺焛a CLI

| Comando | Descri闂佸吋鍎抽埀顒傚枎閻ㄧ珳 |
| --- | --- |
| `4claw onboard` | Inicializar configura闂佸吋鍎抽埀顒傚枎閻ㄧ珳 & workspace |
| `4claw agent -m "..."` | Conversar com o agente |
| `4claw agent` | Modo de chat interativo |
| `4claw gateway` | Iniciar o gateway (para bots de chat) |
| `4claw status` | Mostrar status |
| `4claw cron list` | Listar todas as tarefas agendadas |
| `4claw cron add ...` | Adicionar uma tarefa agendada |

### Tarefas Agendadas / Lembretes

O 4claw suporta lembretes agendados e tarefas recorrentes por meio da ferramenta `cron`:

* **Lembretes 闂佺粯鎼╁Σ姒燾os**: "Remind me in 10 minutes" (Me lembre em 10 minutos) 闂?dispara uma vez ap闁荤姵鍔楅?10min
* **Tarefas recorrentes**: "Remind me every 2 hours" (Me lembre a cada 2 horas) 闂?dispara a cada 2 horas
* **Express闂佺粯绮庣敮鐥?Cron**: "Remind me at 9am daily" (Me lembre 闂佺厧顕惔?9h todos os dias) 闂?usa express闂佽壈鍋愰獮?cron

As tarefas s闂佽壈鍋愰獮?armazenadas em `~/.4claw/workspace/cron/` e processadas automaticamente.

## 婵☆偓绲介崰姘瀶?Contribuir & Roadmap

PRs s闂佽壈鍋愰獮?bem-vindos! O c闁荤姵鍔楅悵顡痝o-fonte 闂?intencionalmente pequeno e leg闂備焦鍎抽弫濂簂. 婵☆偓绲介崰姘?
Roadmap em breve...

Grupo de desenvolvedores em forma闂佸吋鍎抽埀顒傚枎閻ㄧ珳. Requisito de entrada: Pelo menos 1 PR com merge.

Grupos de usu闁荤姴顑戠粻绉坥s:

Discord: <https://discord.gg/V4sAZ9XWpN>


## 婵☆偓绲介崯顖炲磻?Solu闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de Problemas

### Busca web mostra "API 闂傚倷鐒﹀妯肩矓閸洘鍋柛鈩冪⊕閳锋帗銇勯弴妤€浜惧?

Isso 闂?normal se voc闂?ainda n闂佽壈鍋愰獮?configurou uma API key de busca. O 4claw fornecer闁?links 闂佺粯鎼╁Σ顡玦s para busca manual.

Para habilitar a busca web:

1. **Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 1 (Recomendado)**: Obtenha uma API key gratuita em [https://brave.com/search/api](https://brave.com/search/api) (2000 consultas gr闁荤姴顑戠粻绛皊/m闂備焦瀵ч? para os melhores resultados.
2. **Op闂佸吋鍎抽埀顒傚枎閻ㄧ珳 2 (Sem Cart闂佽壈鍋愰獮?de Cr闂佽偐鍘у畷鎼僼o)**: Se voc闂?n闂佽壈鍋愰獮?tem uma key, o sistema automaticamente usa o **DuckDuckGo** como fallback (sem necessidade de key).

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

### Erros de filtragem de conte闂佺粯鎼╁Σ濉?
Alguns provedores (como Zhipu) possuem filtragem de conte闂佺粯鎼╁Σ濉? Tente reformular sua pergunta ou use um modelo diferente.

### Bot do Telegram diz "Conflict: terminated by other getUpdates"

Isso acontece quando outra inst闂佽偐鍎ゆ慨顪﹊a do bot est闁?em execu闂佸吋鍎抽埀顒傚枎閻ㄧ珳. Certifique-se de que apenas um `4claw gateway` esteja rodando por vez.

---

## 婵☆偓绲介崯顖炲箟?Compara闂佸吋鍎抽埀顒傚枎閻ㄧ珳 de API Keys

| Servi闂佺厧澧介柈?| Plano Gratuito | Caso de Uso |
| --- | --- | --- |
| **OpenRouter** | 200K tokens/m闂備焦瀵ч?| M闂佺粯鎼╁Σ鏉plos modelos (Claude, GPT-4, etc.) |
| **Zhipu** | 200K tokens/m闂備焦瀵ч?| Melhor para usu闁荤姴顑戠粻绉坥s chineses |
| **Brave Search** | 2000 consultas/m闂備焦瀵ч?| Funcionalidade de busca web |
| **Groq** | Plano gratuito dispon闂備焦鍎抽弫濂簂 | Infer闂備焦瀵ч鈺焛a ultra-r闁荤姴顑戠粻鐬歞a (Llama, Mixtral) |
| **Cerebras** | Plano gratuito dispon闂備焦鍎抽弫濂簂 | Infer闂備焦瀵ч鈺焛a ultra-r闁荤姴顑戠粻鐬歞a (Llama 3.3 70B) |
