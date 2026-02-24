# 浼佷笟寰俊鑷缓搴旂敤 (WeCom App) 閰嶇疆鎸囧崡

鏈枃妗ｄ粙缁嶅浣曞湪 4claw 涓厤缃紒涓氬井淇¤嚜寤哄簲鐢?(wecom-app) 閫氶亾銆?

## 鍔熻兘鐗规€?

| 鍔熻兘 | 鏀寔鐘舵€?|
|------|---------|
| 琚姩鎺ユ敹娑堟伅 | 鉁?|
| 涓诲姩鍙戦€佹秷鎭?| 鉁?|
| 绉佽亰 | 鉁?|
| 缇よ亰 | 鉂?|

## 閰嶇疆姝ラ

### 1. 浼佷笟寰俊鍚庡彴閰嶇疆

1. 鐧诲綍 [浼佷笟寰俊绠＄悊鍚庡彴](https://work.weixin.qq.com/wework_admin)
2. 杩涘叆"搴旂敤绠＄悊" 鈫?閫夋嫨鑷缓搴旂敤
3. 璁板綍浠ヤ笅淇℃伅锛?
   - **AgentId**: 搴旂敤璇︽儏椤垫樉绀?
   - **Secret**: 鐐瑰嚮"鏌ョ湅"鑾峰彇
4. 杩涘叆"鎴戠殑浼佷笟"椤甸潰锛岃褰?**浼佷笟ID** (CorpID)

### 2. 鎺ユ敹娑堟伅閰嶇疆

1. 鍦ㄥ簲鐢ㄨ鎯呴〉锛岀偣鍑?鎺ユ敹娑堟伅"鐨?璁剧疆API鎺ユ敹"
2. 濉啓浠ヤ笅淇℃伅锛?
   - **URL**: `http://your-server:18792/webhook/wecom-app`
   - **Token**: 闅忔満鐢熸垚鎴栬嚜瀹氫箟锛堢敤浜庣鍚嶉獙璇侊級
   - **EncodingAESKey**: 鐐瑰嚮"闅忔満鐢熸垚"鐢熸垚43瀛楃鐨勫瘑閽?
3. 鐐瑰嚮"淇濆瓨"鏃讹紝浼佷笟寰俊浼氬彂閫侀獙璇佽姹?

### 3. 4claw 閰嶇疆

鍦?`config.json` 涓坊鍔犱互涓嬮厤缃細

```json
{
  "channels": {
    "wecom_app": {
      "enabled": true,
      "corp_id": "wwxxxxxxxxxxxxxxxx",           // 浼佷笟ID
      "corp_secret": "xxxxxxxxxxxxxxxxxxxxxxxx", // 搴旂敤Secret
      "agent_id": 1000002,                        // 搴旂敤AgentId
      "token": "your_token",                      // 鎺ユ敹娑堟伅閰嶇疆鐨凾oken
      "encoding_aes_key": "your_encoding_aes_key", // 鎺ユ敹娑堟伅閰嶇疆鐨凟ncodingAESKey
      "webhook_host": "0.0.0.0",
      "webhook_port": 18792,
      "webhook_path": "/webhook/wecom-app",
      "allow_from": [],
      "reply_timeout": 5
    }
  }
}
```

## 甯歌闂

### 1. 鍥炶皟URL楠岃瘉澶辫触

**鐥囩姸**: 浼佷笟寰俊淇濆瓨API鎺ユ敹娑堟伅鏃舵彁绀洪獙璇佸け璐?

**妫€鏌ラ」**:
- 纭鏈嶅姟鍣ㄩ槻鐏宸插紑鏀?18792 绔彛
- 纭 `corp_id`銆乣token`銆乣encoding_aes_key` 閰嶇疆姝ｇ‘
- 鏌ョ湅 4claw 鏃ュ織鏄惁鏈夎姹傚埌杈?

### 2. 涓枃娑堟伅瑙ｅ瘑澶辫触

**鐥囩姸**: 鍙戦€佷腑鏂囨秷鎭椂鍑虹幇 `invalid padding size` 閿欒

**鍘熷洜**: 浼佷笟寰俊浣跨敤闈炴爣鍑嗙殑 PKCS7 濉厖锛?2瀛楄妭鍧楀ぇ灏忥級

**瑙ｅ喅**: 纭繚浣跨敤鏈€鏂扮増鏈殑 4claw锛屽凡淇姝ら棶棰樸€?

### 3. 绔彛鍐茬獊

**鐥囩姸**: 鍚姩鏃舵彁绀虹鍙ｅ凡琚崰鐢?

**瑙ｅ喅**: 淇敼 `webhook_port` 涓哄叾浠栫鍙ｏ紝濡?18794

## 鎶€鏈粏鑺?

### 鍔犲瘑绠楁硶

- **绠楁硶**: AES-256-CBC
- **瀵嗛挜**: EncodingAESKey Base64瑙ｇ爜鍚庣殑32瀛楄妭
- **IV**: AESKey鐨勫墠16瀛楄妭
- **濉厖**: PKCS7锛堝潡澶у皬涓?2瀛楄妭锛岄潪鏍囧噯16瀛楄妭锛?
- **娑堟伅鏍煎紡**: XML

### 娑堟伅缁撴瀯

瑙ｅ瘑鍚庣殑娑堟伅鏍煎紡锛?
```
random(16B) + msg_len(4B) + msg + receiveid
```

鍏朵腑 `receiveid` 瀵逛簬鑷缓搴旂敤鏄?`corp_id`銆?

## 璋冭瘯

鍚敤璋冭瘯妯″紡鏌ョ湅璇︾粏鏃ュ織锛?

```bash
4claw gateway --debug
```

鍏抽敭鏃ュ織鏍囪瘑锛?
- `wecom_app`: WeCom App 閫氶亾鐩稿叧鏃ュ織
- `wecom_common`: 鍔犲瘑瑙ｅ瘑鐩稿叧鏃ュ織

## 鍙傝€冩枃妗?

- [浼佷笟寰俊瀹樻柟鏂囨。 - 鎺ユ敹娑堟伅](https://developer.work.weixin.qq.com/document/path/96211)
- [浼佷笟寰俊瀹樻柟鍔犺В瀵嗗簱](https://github.com/sbzhu/weworkapi_golang)
