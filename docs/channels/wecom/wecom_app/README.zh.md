# 娴间椒绗熷顔讳繆閼奉亜缂撴惔鏃傛暏

娴间椒绗熷顔讳繆閼奉亜缂撴惔鏃傛暏閺勵垱瀵氭导浣风瑹閸︺劋绱掓稉姘簳娣団€茶厬閸掓稑缂撻惃鍕安閻㈩煉绱濇稉鏄忣洣閻劋绨导浣风瑹閸愬懘鍎存担璺ㄦ暏閵嗗倿鈧俺绻冩导浣风瑹瀵邦喕淇婇懛顏勭紦鎼存梻鏁ら敍灞肩磼娑撴艾褰叉禒銉ョ杽閻滈绗岄崨妯轰紣閻ㄥ嫰鐝弫鍫熺煛闁艾鎷伴崡蹇庣稊閿涘本褰佹妯轰紣娴ｆ粍鏅ラ悳鍥モ偓?

## 闁板秶鐤?
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
      "allow_from": [],
      "reply_timeout": 5
    }
  }
}
```

| 鐎涙顔?            | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                                    |
| ---------------- | ------ | ---- | ---------------------------------------- |
| corp_id          | string | 閺?  | 娴间椒绗?ID                                  |
| corp_secret      | string | 閺?  | 鎼存梻鏁ょ粙瀣碍鐎靛棝鎸?                            |
| agent_id         | int    | 閺?  | 鎼存梻鏁ょ粙瀣碍娴狅絿鎮?ID                          |
| token            | string | 閺?  | 閸ョ偠鐨熸宀冪槈娴犮倗澧?                            |
| encoding_aes_key | string | 閺?  | 43 鐎涙顑?AES 鐎靛棝鎸?                        |
| webhook_host     | string | 閸?  | HTTP 閺堝秴濮熼崳銊х拨鐎规艾婀撮崸鈧?                     |
| webhook_port     | int    | 閸?  | HTTP 閺堝秴濮熼崳銊ь伂閸欙綇绱欐妯款吇閿?8792閿?          |
| webhook_path     | string | 閸?  | Webhook 鐠侯垰绶為敍鍫ョ帛鐠併倧绱?webhook/wecom-app閿?|
| allow_from       | array  | 閸?  | 閻劍鍩?ID 閻ц棄鎮曢崡?                          |
| reply_timeout    | int    | 閸?  | 閸ョ偛顦茬搾鍛閺冨爼妫块敍鍫㈩潡閿?                      |

## 鐠佸墽鐤嗗ù浣衡柤

1. 閻ц缍?[娴间椒绗熷顔讳繆缁狅紕鎮婇崥搴″酱](https://work.weixin.qq.com/)
2. 鏉╂稑鍙嗛垾婊冪安閻劎顓搁悶鍡忊偓?-> 閳ユ粌鍨卞鍝勭安閻劉鈧?
3. 閼惧嘲褰囨导浣风瑹 ID (CorpID) 閸滃苯绨查悽?Secret
4. 閸︺劌绨查悽銊啎缂冾喕鑵戦柊宥囩枂閳ユ粍甯撮弨鑸电Х閹垪鈧繐绱濋懢宄板絿 Token 閸?EncodingAESKey
5. 鐠佸墽鐤嗛崶鐐剁殶 URL 娑?`http://<your-server-ip>:<port>/webhook/wecom-app`
6. 鐏?CorpID, Secret, AgentID 缁涘淇婇幁顖氾綖閸忋儵鍘ょ純顔芥瀮娴?
