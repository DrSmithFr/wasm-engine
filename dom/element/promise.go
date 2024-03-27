package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type Promise struct {
	*wasm.Entity
}

func NewPromise(raw js.Value) *Promise {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Promise{
		Entity: wasm.New(raw),
	}
}

// Then Appends fulfillment and rejection handlers to the promise, and returns a new promise resolving to the return value of the called handler.
func (p *Promise) Then(onFulfilled js.Func) *Promise {
	p.Js().Call("then", onFulfilled)
	return p
}

// Catch Appends a rejection handler callback to the promise, and returns a new promise resolving to the return value of the callback if it is called, or to its original fulfillment value if the promise is instead fulfilled.
func (p *Promise) Catch(onRejected js.Func) *Promise {
	p.Js().Call("catch", onRejected)
	return p
}
