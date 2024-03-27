package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

func NewHTMLElementList(items js.Value) []*HTMLElement {
	elements := make([]*HTMLElement, items.Length())

	for i := 0; i < items.Length(); i++ {
		elements[i] = NewHTMLElement(items.Index(i))
	}

	return elements
}

///////////////////////

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
