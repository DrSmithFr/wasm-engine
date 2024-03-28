package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Iterator https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Iteration_protocols
type Iterator struct {
	*wasm.Entity
}

// NewIterator returns a new Iterator object.
func NewIterator(iterator js.Value) *Iterator {
	return &Iterator{
		Entity: wasm.New(iterator),
	}
}

// Next Returns a key/value pair.
func (i *Iterator) Next() js.Value {
	return i.Js().Call("next")
}

// Done Returns a boolean indicating whether the iterator has been consumed.
func (i *Iterator) Done() bool {
	return i.Js().Get("done").Bool()
}

// Value Returns the current value of the iterator.
func (i *Iterator) Value() js.Value {
	return i.Js().Get("value")
}

func (i *Iterator) Values() []js.Value {
	var values []js.Value

	for !i.Done() {
		values = append(values, i.Value())
		i.Next()
	}

	return values
}
