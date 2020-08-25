// +build wasm js

package main

import (
	"syscall/js"

	"github.com/depadiernos/golangoflife/cmd"
)

func main() {
	println("GolangOfLife has been initialized")
	cmd.Cmd()
}

func other() {
	js.Global().Get("document")
}
