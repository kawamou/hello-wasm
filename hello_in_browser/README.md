## 実行方法
```foo.sh
# 参考
# https://zenn.dev/tomi/articles/2020-11-10-go-web11

# コンパイルする
tinygo build -o guest.wasm -target wasm ./guest.go

# 吐き出されたWasmをみたいなら下記を実行し任意のエディタで確認
wasm2wat guest.wasm > guest.wat

# 実行
go run server.go

# 確認（ブラウザで下記リンクにアクセスし検証ツールでコンソールを確認）
http://localhost:8080
```