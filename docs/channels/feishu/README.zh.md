# 妞嬬偘鍔?
妞嬬偘鍔熼敍鍫濇禇闂勫懐澧楅崥宥囆為敍姝€ark閿涘妲哥€涙濡捄鍐插З閺冩ぞ绗呴惃鍕磼娑撴艾宕楁担婊冮挬閸欒埇鈧倸鐣犻柅姘崇箖娴滃娆㈡す鍗炲З閻?Webhook 閸氬本妞傞弨顖涘瘮娑擃厼娴楅崪灞藉弿閻炲啫绔堕崷鎭掆偓?

## 闁板秶鐤?
```json
{
  "channels": {
    "feishu": {
      "enabled": true,
      "app_id": "cli_xxx",
      "app_secret": "xxx",
      "encrypt_key": "",
      "verification_token": "",
      "allow_from": []
    }
  }
}
```

| 鐎涙顔?              | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                            |
| ------------------ | ------ | ---- | -------------------------------- |
| enabled            | bool   | 閺?  | 閺勵垰鎯侀崥顖滄暏妞嬬偘鍔熸０鎴︿壕                 |
| app_id             | string | 閺?  | 妞嬬偘鍔熸惔鏃傛暏閻?App ID(娴狀櫓li\_瀵偓婢?   |
| app_secret         | string | 閺?  | 妞嬬偘鍔熸惔鏃傛暏閻?App Secret            |
| encrypt_key        | string | 閸?  | 娴滃娆㈤崶鐐剁殶閸旂姴鐦戠€靛棝鎸?                |
| verification_token | string | 閸?  | 閻劋绨琖ebhook娴滃娆㈡宀冪槈閻ㄥ嚲oken       |
| allow_from         | array  | 閸?  | 閻劍鍩汭D閻ц棄鎮曢崡鏇礉缁岄缚銆冪粈鍝勫帒鐠佸憡澧嶉張澶屾暏閹?|

## 鐠佸墽鐤嗗ù浣衡柤

1. 閸撳秴绶?[妞嬬偘鍔熷鈧弨鎯ч挬閸欑櫓(https://open.feishu.cn/)閸掓稑缂撴惔鏃傛暏缁嬪绨?2. 閼惧嘲褰?App ID 閸?App Secret
3. 闁板秶鐤嗘禍瀣╂鐠併垽妲勯崪瀛竐bhook URL
4. 鐠佸墽鐤嗛崝鐘茬槕(閸欘垶鈧?閻㈢喍楠囬悳顖氼暔瀵ら缚顔呴崥顖滄暏)
5. 鐏?App ID閵嗕竸pp Secret閵嗕笒ncrypt Key 閸?Verification Token(婵″倹鐏夐崥顖滄暏閸旂姴鐦? 婵夘偄鍙嗛柊宥囩枂閺傚洣娆㈡稉?
