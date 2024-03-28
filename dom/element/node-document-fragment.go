package element

import "syscall/js"

// DocumentFragment https://developer.mozilla.org/en-US/docs/Web/API/DocumentFragment
type DocumentFragment struct {
	*Node
}

func NewDocumentFragment(raw js.Value) *DocumentFragment {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &DocumentFragment{
		Node: NewNode(raw),
	}
}

// ChildElementCount Returns the number of child elements.
func (d *DocumentFragment) ChildElementCount() int {
	return d.Js().Get("childElementCount").Int()
}

// Children Returns the child elements.
func (d *DocumentFragment) Children() *HTMLCollection {
	return NewHTMLCollection(d.Js().Get("children"))
}

// FirstElementChild Returns the first child element.
func (d *DocumentFragment) FirstElementChild() *HTMLElement {
	return NewHTMLElement(d.Js().Get("firstElementChild"))
}

// LastElementChild Returns the last child element.
func (d *DocumentFragment) LastElementChild() *HTMLElement {
	return NewHTMLElement(d.Js().Get("lastElementChild"))
}

// Append Adds nodes to the end of the list of children.
func (d *DocumentFragment) Append(nodes ...*Node) {
	for _, node := range nodes {
		d.Js().Call("append", node.Js())
	}
}

// Prepend Adds nodes to the beginning of the list of children.
func (d *DocumentFragment) Prepend(nodes ...*Node) {
	for _, node := range nodes {
		d.Js().Call("prepend", node.Js())
	}
}

// QuerySelector Returns the first element that matches a specified CSS selector.
func (d *DocumentFragment) QuerySelector(selector string) *HTMLElement {
	return NewHTMLElement(d.Js().Call("querySelector", selector))
}

// QuerySelectorAll Returns a list of elements that matches a specified CSS selector.
func (d *DocumentFragment) QuerySelectorAll(selector string) []*HTMLElement {
	return NewHTMLElementList(d.Js().Call("querySelectorAll", selector))
}

// ReplaceChildren Replaces the children of the node with the specified nodes.
func (d *DocumentFragment) ReplaceChildren(nodes ...*Node) {
	for _, node := range nodes {
		d.Js().Call("replaceChildren", node.Js())
	}
}

// GetElementById Returns the element with the specified ID.
func (d *DocumentFragment) GetElementById(id string) *HTMLElement {
	return NewHTMLElement(d.Js().Call("getElementById", id))
}
