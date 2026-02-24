# Telegram

Telegram Channel 闁俺绻?Telegram 閺堝搫娅掓禍?API 娴ｈ法鏁ら梹鑳枂鐠囥垹鐤勯悳鏉跨唨娴滃孩婧€閸ｃ劋姹夐惃鍕偓姘繆閵嗗倸鐣犻弨顖涘瘮閺傚洦婀板☉鍫熶紖閵嗕礁鐛熸担鎾绘娴犺绱欓悡褏澧栭妴浣筋嚔闂婄偨鈧線鐓舵０鎴欌偓浣规瀮濡楋綇绱氶妴渚€鈧俺绻?Groq Whisper 鏉╂稖顢戠拠顓㈢叾鏉烆剙缍嶆禒銉ュ挤閸愬懐鐤嗛崨鎴掓姢婢跺嫮鎮婇崳銊ｂ偓?

## 闁板秶鐤?
```json
{
  "channels": {
    "telegram": {
      "enabled": true,
      "token": "123456789:ABCdefGHIjklMNOpqrsTUVwxyz",
      "allow_from": ["123456789"],
      "proxy": ""
    }
  }
}
```

| 鐎涙顔?      | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                                                     |
| ---------- | ------ | ---- | --------------------------------------------------------- |
| enabled    | bool   | 閺?  | 閺勵垰鎯侀崥顖滄暏 Telegram 妫版垿浜?                                   |
| token      | string | 閺?  | Telegram 閺堝搫娅掓禍?API Token                                 |
| allow_from | array  | 閸?  | 閻劍鍩汭D閻ц棄鎮曢崡鏇礉缁岄缚銆冪粈鍝勫帒鐠佸憡澧嶉張澶屾暏閹?                         |
| proxy      | string | 閸?  | 鏉╃偞甯?Telegram API 閻ㄥ嫪鍞悶?URL (娓氬顩?http://127.0.0.1:7890) |

## 鐠佸墽鐤嗗ù浣衡柤

1. 閸?Telegram 娑擃厽鎮崇槐?`@BotFather`
2. 閸欐垿鈧?`/newbot` 閸涙垝鎶ら獮鑸靛瘻閻撗勫絹缁€鍝勫灡瀵ょ儤鏌婇張鍝勬珤娴?
3. 閼惧嘲褰?HTTP API Token
4. 鐏?Token 婵夘偄鍙嗛柊宥囩枂閺傚洣娆㈡稉?
5. (閸欘垶鈧? 闁板秶鐤?`allow_from` 娴犮儵妾洪崚璺哄帒鐠侀晲绨伴崝銊ф畱閻劍鍩?ID (閸欘垶鈧俺绻?`@userinfobot` 閼惧嘲褰?ID)
