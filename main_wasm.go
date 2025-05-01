package main

import (
	"fmt"
	"syscall/js"

	"github.com/juicedata/juicefs/wasm"
)

func logToJS(message string) {
	js.Global().Get("console").Call("log", "[Go]:", message)
}

func registerFuncs() {
	js.Global().Set("format", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "Database URL parameter required"
		}
		dbUrl := args[0].String()

		err := wasm.Format(dbUrl)
		if err != nil {
			errMsg := fmt.Sprintf("Format failed: %v", err)
			logToJS(errMsg)
			return errMsg
		}

		logToJS("Format completed successfully")
		return nil
	}))
}

func main() {
	registerFuncs()

	logToJS("JuiceFS WebAssembly module initialized")

	select {}
}
