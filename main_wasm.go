package main

import (
	"strconv"
	"github.com/juicedata/juicefs/wasm"
)

func main() {
  wasm.RegisterFuncs()
  sum := wasm.TestAdd(1, 2)
  wasm.LogToJS("TestAdd(1, 2) = " + strconv.Itoa(sum))

	// wasm.LogToJS("JuiceFS WebAssembly module initialized")
	select {}
}
