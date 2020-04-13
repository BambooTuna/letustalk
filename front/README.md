# letustalk

## 環境変数
- VUE_APP_PAYMENT_PUB_KEY

    決済サービスの公開鍵

- VUE_APP_SERVER_ENDPOINT

    バックエンドのエンドポイント

## ローカル動作確認
```bash
$ VUE_APP_PAYMENT_PUB_KEY=pk_test_6acc32ac21aa1b0f55d5e3b8 \
VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 \
npm run serve -- --port 8080 --host 0.0.0.0


$ VUE_APP_PAYMENT_PUB_KEY=pk_test_6acc32ac21aa1b0f55d5e3b8 \
VUE_APP_SERVER_ENDPOINT=http://localhost:8080/v1 \
npm run build

```
