package main

import (
	"fmt"
	"syscall/js"

	"github.com/juicedata/juicefs/wasm"
)

func logToJS(message string) {
	js.Global().Get("console").Call("log", "[Go]:", message)
}

// 注册一个函数，用于从JS传递字符串到Go并返回，测试中文编码
func registerStringEcho() {
	js.Global().Set("goEchoString", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "No parameters provided"
		}
		input := args[0].String()
		fmt.Println("Received string from JS:", input)
		return input
	}))
}

func registerFuncs() {
	js.Global().Set("format", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) < 1 {
			return "Database URL parameter required"
		}
		dbUrl := args[0].String()
		logToJS("Starting format: " + dbUrl)

		// 尝试格式化操作
		err := wasm.Format(dbUrl)
		if err != nil {
			errMsg := fmt.Sprintf("Format failed: %v", err)
			logToJS(errMsg)
			return errMsg
		}

		logToJS("Format completed successfully")
		return "Format completed"
	}))
}

func main() {
	registerStringEcho()
	registerFuncs() 

	logToJS("JuiceFS WebAssembly module initialized")

	select {}
}
