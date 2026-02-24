# Slack

Slack 閺勵垰鍙忛悶鍐暙閸忓牏娈戞导浣风瑹缁狙冨祮閺冨爼鈧俺顔嗛獮鍐插酱閵嗕揪icoClaw 闁插洨鏁?Slack 閻?Socket Mode 鐎圭偟骞囩€圭偞妞傞崣灞芥倻闁矮淇婇敍灞炬￥闂団偓闁板秶鐤嗛崗顒€绱戦惃?Webhook 缁旑垳鍋ｉ妴?

## 闁板秶鐤?
```json
{
  "channels": {
    "slack": {
      "enabled": true,
      "bot_token": "xoxb-...",
      "app_token": "xapp-...",
      "allow_from": []
    }
  }
}
```

| 鐎涙顔?      | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                                                    |
| ---------- | ------ | ---- | -------------------------------------------------------- |
| enabled    | bool   | 閺?  | 閺勵垰鎯侀崥顖滄暏 Slack 妫版垿浜?                                     |
| bot_token  | string | 閺?  | Slack 閺堝搫娅掓禍铏规畱 Bot User OAuth Token (娴?xoxb- 瀵偓婢?      |
| app_token  | string | 閺?  | Slack 鎼存梻鏁ら惃?Socket Mode App Level Token (娴?xapp- 瀵偓婢? |
| allow_from | array  | 閸?  | 閻劍鍩汭D閻ц棄鎮曢崡鏇礉缁岄缚銆冪粈鍝勫帒鐠佸憡澧嶉張澶屾暏閹?                        |

## 鐠佸墽鐤嗗ù浣衡柤

1. 閸撳秴绶?[Slack API](https://api.slack.com/) 閸掓稑缂撴稉鈧稉顏呮煀閻?Slack 鎼存梻鏁?2. 閸氼垳鏁?Socket Mode 楠炴儼骞忛崣?App Level Token
3. 濞ｈ濮?Bot Token Scopes(娓氬顩chat:write`閵嗕梗im:history`缁?
4. 鐎瑰顥婃惔鏃傛暏閸掓澘浼愭担婊冨隘楠炴儼骞忛崣?Bot User OAuth Token
5. 鐏?Bot Token 閸?App Token 婵夘偄鍙嗛柊宥囩枂閺傚洣娆㈡稉?
