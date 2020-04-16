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

## スケジュール

```bash
$ curl -X GET "localhost:8080/v1/account/1/schedule?from=20200420000000&to=20200420030000"

$ curl -X POST localhost:8080/v1/schedule/2/reserve

```

## Wiki
https://github.com/BambooTuna/letustalk/wiki

## Go Test
```bash
$ go test ./... -v
```
