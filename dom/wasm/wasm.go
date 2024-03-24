package wasm

import (
	"syscall/js"
)

type WASM interface {
	Js() js.Value
}
