# 娴间椒绗熷顔讳繆閺堝搫娅掓禍?

娴间椒绗熷顔讳繆閺堝搫娅掓禍鐑樻Ц娴间椒绗熷顔讳繆閹绘劒绶甸惃鍕缁夊秴鎻╅柅鐔稿复閸忋儲鏌熷蹇ョ礉閸欘垯浜掗柅姘崇箖 Webhook URL 閹恒儲鏁瑰☉鍫熶紖閵?

## 闁板秶鐤?
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
      "allow_from": [],
      "reply_timeout": 5
    }
  }
}
```

| 鐎涙顔?            | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                                        |
| ---------------- | ------ | ---- | -------------------------------------------- |
| token            | string | 閺?  | 缁涙儳鎮曟宀冪槈娴狅絽绔?                                |
| encoding_aes_key | string | 閺?  | 閻劋绨憴锝呯槕閻?43 鐎涙顑?AES 鐎靛棝鎸?                 |
| webhook_url      | string | 閺?  | 閻劋绨崣鎴︹偓浣告礀婢跺秶娈戞导浣风瑹瀵邦喕淇婄紘銈堜喊閺堝搫娅掓禍?Webhook URL |
| webhook_host     | string | 閸?  | HTTP 閺堝秴濮熼崳銊х拨鐎规艾婀撮崸鈧敍鍫ョ帛鐠併倧绱?.0.0.0閿?        |
| webhook_port     | int    | 閸?  | HTTP 閺堝秴濮熼崳銊ь伂閸欙綇绱欐妯款吇閿?8793閿?              |
| webhook_path     | string | 閸?  | Webhook 缁旑垳鍋ｇ捄顖氱窞閿涘牓绮拋銈忕窗/webhook/wecom閿?    |
| allow_from       | array  | 閸?  | 閻劍鍩?ID 閻ц棄鎮曢崡鏇礄缁屽搫鈧?= 閸忎浇顔忛幍鈧張澶屾暏閹村嚖绱?       |
| reply_timeout    | int    | 閸?  | 閸ョ偛顦茬搾鍛閺冨爼妫块敍鍫濆礋娴ｅ稄绱扮粔鎺炵礉姒涙顓婚崐纭风窗5閿?         |

## 鐠佸墽鐤嗗ù浣衡柤

1. 閸︺劋绱掓稉姘簳娣囷紕鍏㈡稉顓熷潑閸旂姵婧€閸ｃ劋姹?2. 閼惧嘲褰?Webhook URL
3. (婵″倿娓堕幒銉︽暪濞戝牊浼? 閸︺劍婧€閸ｃ劋姹夐柊宥囩枂妞ょ敻娼扮拋鍓х枂閹恒儲鏁瑰☉鍫熶紖閻?API 閸︽澘娼冮敍鍫濇礀鐠嬪啫婀撮崸鈧敍澶変簰閸?Token 閸?EncodingAESKey
4. 鐏忓棛娴夐崗鍏呬繆閹垰锝為崗銉╁帳缂冾喗鏋冩禒?
