package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// HTMLCollection https://developer.mozilla.org/en-US/docs/Web/API/HTMLCollection
type HTMLCollection struct {
	*wasm.Entity
}

func NewHTMLCollection(raw js.Value) *HTMLCollection {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &HTMLCollection{
		Entity: wasm.New(raw),
	}
}

// Length Returns the number of elements in the collection.
func (c *HTMLCollection) Length() int {
	return c.Js().Get("length").Int()
}

// Item Returns the element at the specified index in the collection.
func (c *HTMLCollection) Item(index int) *HTMLElement {
	return NewHTMLElement(c.Js().Call("item", index))
}

// NamedItem Returns the element with the specified ID or name.
func (c *HTMLCollection) NamedItem(name string) *HTMLElement {
	return NewHTMLElement(c.Js().Call("namedItem", name))
}
