# OneBot

OneBot 閺勵垯绔存稉顏堟桨閸?QQ 閺堝搫娅掓禍铏规畱瀵偓閺€鎯у礂鐠侇喗鐖ｉ崙鍡礉娑撳搫顦跨粔?QQ 閺堝搫娅掓禍鍝勭杽閻滃府绱欐笟瀣洤 go-cqhttp閵嗕府irai閿涘褰佹笟娑楃啊缂佺喍绔撮惃鍕复閸欙絻鈧倸鐣犳担璺ㄦ暏 WebSocket 鏉╂稖顢戦柅姘繆閵?

## 闁板秶鐤?
```json
{
  "channels": {
    "onebot": {
      "enabled": true,
      "ws_url": "ws://localhost:8080",
      "access_token": "",
      "allow_from": []
    }
  }
}
```

| 鐎涙顔?        | 缁鐎?  | 韫囧懎锝?| 閹诲繗鍫?                            |
| ------------ | ------ | ---- | -------------------------------- |
| enabled      | bool   | 閺?  | 閺勵垰鎯侀崥顖滄暏 OneBot 妫版垿浜?            |
| ws_url       | string | 閺?  | OneBot 閺堝秴濮熼崳銊ф畱 WebSocket URL    |
| access_token | string | 閸?  | 鏉╃偞甯?OneBot 閺堝秴濮熼崳銊ф畱鐠佸潡妫舵禒銈囧     |
| allow_from   | array  | 閸?  | 閻劍鍩汭D閻ц棄鎮曢崡鏇礉缁岄缚銆冪粈鍝勫帒鐠佸憡澧嶉張澶屾暏閹?|

## 鐠佸墽鐤嗗ù浣衡柤

1. 闁劎璁叉稉鈧稉?OneBot 閸忕厧顔愰惃鍕杽閻?娓氬顩apcat)
2. 闁板秶鐤?OneBot 鐎圭偟骞囨禒銉ユ儙閻?WebSocket 閺堝秴濮熼獮鎯邦啎缂冾喛顔栭梻顔绘姢閻?婵″倹鐏夐棁鈧憰?
3. 鐏?WebSocket URL 閸滃矁顔栭梻顔绘姢閻楀苯锝為崗銉╁帳缂冾喗鏋冩禒鏈佃厬
