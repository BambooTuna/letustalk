# letustalk

## ローカル動作確認
1. インフラの構築
```bash
$ cd infra
$ docker-compose up --build
```

2. フロントのVueプロジェクトをBuildして静的ファイルを作る
```bash
$ cd front
$ VUE_APP_PAYMENT_PUB_KEY=pk_test_6acc32ac21aa1b0f55d5e3b8 \
VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 \
npm run build
```

3. APIサーバーを起動する
※フロントの静的ファイルのホスティングも兼ねている
```bash
$ go run main.go
```

## テスト
```bash
$ go test ./... -v

$ cd front
$ npm run lint
```

## バックエンドAPIについて
https://github.com/BambooTuna/letustalk/wiki
