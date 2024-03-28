package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// DOMStringList https://developer.mozilla.org/en-US/docs/Web/API/DOMStringList
type DOMStringList struct {
	*wasm.Entity
}

func NewDOMStringList(raw js.Value) *DOMStringList {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &DOMStringList{
		Entity: wasm.New(raw),
	}
}

// Length Returns the number of strings in the list.
func (d *DOMStringList) Length() int {
	return d.Js().Get("length").Int()
}

// Item Returns the string at the specified index in the list.
func (d *DOMStringList) Item(index int) string {
	return d.Js().Call("item", index).String()
}

// Contains Returns true if the string is in the list.
func (d *DOMStringList) Contains(string string) bool {
	return d.Js().Call("contains", string).Bool()
}
