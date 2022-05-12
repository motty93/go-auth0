## Auth0 access token
```
$ go run main.go
```

### 設定

```
$ cp env_sample .env
```
※.envはauth0のものを + 既存ユーザーのusername/passwordを設定

### コンテナで実行
```
$ docker build -t go-auth0:latest .

$ docker container run -it -d --name go-auth0-container go-auth0:latest sh

$ docker container cp go-auth0-container:/go/src/app/tmp ./

$ docker container rm -f go-auth0-container

$ ./tmp/main
```

もしくは
```
$ ./scripts/run.sh

$ ./scripts/build_run.sh
```

## 環境
```
$ go version
go version go1.17 linux/amd64
```

※docker imageはgo version1.18

## 詳細・注意事項

- 設定項目の環境変数を必ず設定してください
- default directory/audience等の設定は必要です
- auth0 databaseにユーザーが保存されている状態じゃないと動きません
- auth0 API referenceにある通り、sigle page applicationとnative applicationでは動きません
- regular web applicationでもadvanced settingでgrant typeにチェックがついてないと動きません
