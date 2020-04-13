# letustalk


## 請求
```bash
$ curl -X GET localhost:8080/v1/invoice/:invoiceId -i
$ curl -X POST -H "Content-Type: application/json" -d '{"amount":1000}' localhost:8080/v1/invoice -i

$ curl -X POST -H "Content-Type: application/json" -d '{"token":""}' -i localhost:8080/v1/pay/:invoiceId
```

## ユーザー
```bash
$ curl -X GET localhost:8080/v1/mentor -i
```


## Wiki
https://github.com/BambooTuna/letustalk/wiki
