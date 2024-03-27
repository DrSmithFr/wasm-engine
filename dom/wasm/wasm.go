package wasm

import (
	"syscall/js"
)

type Wasm interface {
	Js() js.Value
}

type Entity struct {
	entity js.Value
}

func (w *Entity) Js() js.Value {
	return w.entity
}

func New(value js.Value) *Entity {
	if value.IsNull() || value.IsUndefined() {
		return nil
	}

	return &Entity{entity: value}
}
