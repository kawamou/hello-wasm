package main

// https://text.baldanders.info/golang/wasi-with-tinygo/

import (
	_ "embed"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/bytecodealliance/wasmtime-go"
)

//go:embed guest/guest.wasm
var wasm []byte

func main() {
	dir, err := ioutil.TempDir("", "out")
	check(err)
	defer os.RemoveAll(dir)
	stdoutPath := filepath.Join(dir, "stdout")

	// Almost all operations in wasmtime require a contextual `store`
	// argument to share, so create that first
	store := wasmtime.NewStore(wasmtime.NewEngine())

	linker := wasmtime.NewLinker(store)

	// Configure WASI imports to write stdout into a file.
	wasiConfig := wasmtime.NewWasiConfig()
	wasiConfig.SetStdoutFile(stdoutPath)

	// Set the version to the same as in the WAT.
	wasi, err := wasmtime.NewWasiInstance(store, wasiConfig, "wasi_snapshot_preview1")
	check(err)

	// Link WASI
	err = linker.DefineWasi(wasi)
	check(err)

	// Once we have our binary `wasm` we can compile that into a `*Module`
	// which represents compiled JIT code.
	module, err := wasmtime.NewModule(store.Engine, wasm)
	check(err)

	// Next up we instantiate a module which is where we link in all our
	// imports. We've got one import so we pass that in here.
	instance, err := linker.Instantiate(module)
	check(err)

	// After we've instantiated we can lookup our `run` function and call
	// it.
	run := instance.GetExport("_start").Func()
	_, err = run.Call()
	check(err)

	// Print WASM stdout
	out, err := ioutil.ReadFile(stdoutPath)
	check(err)
	fmt.Print(string(out))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
