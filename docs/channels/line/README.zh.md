# Line

4claw 闂侇偅淇虹换?LINE Messaging API 闂佹澘绉撮幃?Webhook 闁搞儳鍋犻惃鐔煎礉閻旇鍘撮悗鍦仧楠炲洨鈧?LINE 闁汇劌瀚弫顕€骞愭担纰樺亾?

## 闂佹澘绉堕悿?

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

| 閻庢稒顨嗛?                | 缂侇偉顕ч悗?  | 闊洤鎳庨敐?| 闁硅绻楅崼?                                      |
| -------------------- | ------ | ---- | ------------------------------------------ |
| enabled              | bool   | 闁?  | 闁哄嫷鍨伴幆渚€宕ラ婊勬殢 LINE Channel                      |
| channel_secret       | string | 闁?  | LINE Messaging API 闁?Channel Secret       |
| channel_access_token | string | 闁?  | LINE Messaging API 闁?Channel Access Token |
| webhook_host         | string | 闁?  | Webhook 闁烩晜鍨甸幆澶愭儍閸曨亜鐦滈柡鍫濇惈濠€鎾锤閳?(闂侇偅鑹鹃悥鑸电▔?0.0.0.0)    |
| webhook_port         | int    | 闁?  | Webhook 闁烩晜鍨甸幆澶愭儍閸曨収浼傞柛?(濮掓稒顭堥缁樼▔?18791)          |
| webhook_path         | string | 闁?  | Webhook 闁汇劌瀚惌鎯ь嚗?(濮掓稒顭堥缁樼▔?/webhook/line)      |
| allow_from           | array  | 闁?  | 闁活潿鍔嶉崺姹璂闁谎嗘閹洟宕￠弴顏嗙缂佸矂缂氶妴鍐矆閸濆嫬甯掗悹浣告啞婢у秹寮垫径灞炬殢闁?          |

## 閻犱礁澧介悿鍡椕规担琛℃煠

1. 闁告挸绉寸欢?[LINE Developers Console](https://developers.line.biz/console/) 闁告帗绋戠紓鎾寸▔閳ь剚绋夐鍛疀闁告柡鍓濊ぐ浣圭瑹濞戞ɑ娅岄柛婊冨缁斿瓨绋?Messaging API Channel
2. 闁兼儳鍢茶ぐ?Channel Secret 闁?Channel Access Token
3. 闂佹澘绉堕悿鍝bhook:
   - Line閻熸洑鐒﹂惇鐧bhook闊洤鎳橀妴蹇旀媴鐠恒劍鏆廐TTPS闁告绻楅鍛存晬鐏炶姤绀堟慨婵勫€濆〒鍓佹啺娓氣偓閸庡绱旈煫顓狀伇濞戞搩浜濋弫顕€骞愭稉鐕璗PS闁汇劌瀚﹢鍥礉閳ヨ櫕鐝ら柨娑樻湰閸ㄣ劑鎳撻崨顒€鈻忛柣顫妼瀵粙宕ラ幋婊冩暕闁荤偛妫楁导鎰板礂瀹勯偊娲grok閻忓繐妫欏﹢浼村捶閻楀牊绠涢柛鏂衡偓铏彜闁哄棙鎸冲﹢鍫曞礆閺夊灝褰嗙紓?
   - 閻?Webhook URL 閻犱礁澧介悿鍡樼▔?`https://your-domain.com/webhook/line`
   - 闁告凹鍨抽弫?Webhook 妤犵偛鐖奸悰娆戞嫚?URL
4. 閻?Channel Secret 闁?Channel Access Token 濠靛鍋勯崣鍡涙煀瀹ュ洨鏋傞柡鍌氭矗濞嗐垺绋?
