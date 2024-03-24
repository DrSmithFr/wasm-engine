package helper

import "syscall/js"

func NewUint8Array(length int) js.Value {
	return js.Global().Get("Uint8Array").New(length)
}
