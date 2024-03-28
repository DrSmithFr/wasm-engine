package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Request https://developer.mozilla.org/en-US/docs/Web/API/Request
type Request struct {
	*wasm.Entity
}

// NewRequest returns a new Request object.
func NewRequest(raw js.Value) *Request {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Request{
		Entity: wasm.New(raw),
	}
}
