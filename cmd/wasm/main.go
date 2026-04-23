//go:build js && wasm

package main

import (
	"syscall/js"

	"github.com/tryuuu/ai-formatter/internal/formatter"
)

func main() {
	js.Global().Set("aiFormat", js.FuncOf(func(_ js.Value, args []js.Value) any {
		if len(args) == 0 {
			return ""
		}
		return formatter.Format(args[0].String())
	}))

	// プロセスを生かし続けて関数を有効に保つ
	select {}
}
