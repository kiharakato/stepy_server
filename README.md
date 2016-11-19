
## 依存ライブラリ
```
godep get
```

## 実行されるの必要なもの
redis
mysql

## 実行
```
GO_ENV=DEV gin -a 8080 run main.go
```

アクセスをする
http://localhost:3000/ping
