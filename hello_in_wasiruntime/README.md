## 実行方法
```foo.sh
# 参考
# https://github.com/bytecodealliance/wasmtime-go/blob/main/README.md

# コンパイルする
tinygo build -o guest.wasm -target wasi ./guest.go

# 吐き出されたWasmをみたいなら下記を実行し任意のエディタで確認
wasm2wat guest.wasm > guest.wat

# 実行する①
# 単純にwasmtimeの上で実行
wasmtime guest.wasm

# 実行する②
# wasmtimeのGoパッケージ上で実行
go run host.go
```