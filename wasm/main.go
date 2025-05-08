//go:build wasm
// +build wasm

package wasm

import (
	"fmt"
	"syscall/js"
)

func LogToJS(message string) {
	js.Global().Get("console").Call("log", "[JuiceFS]:", message)
}

func RegisterFuncs() {
	js.Global().Set("format", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "Database URL parameter required"
		}
		dbUrl := args[0].String()

		err := Format(dbUrl)
		if err != nil {
			errMsg := fmt.Sprintf("Format failed: %v", err)
			LogToJS(errMsg)
			return errMsg
			}

		LogToJS("Format completed successfully")
		return nil
	}))
}

func TestAdd(a, b int) int {
  return js.Global().Get("add").Invoke(a, b).Int()
}
