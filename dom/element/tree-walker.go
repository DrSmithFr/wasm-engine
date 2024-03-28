package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// TreeWalker https://developer.mozilla.org/en-US/docs/Web/API/TreeWalker
type TreeWalker struct {
	*wasm.Entity
}

func NewTreeWalker(raw js.Value) *TreeWalker {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &TreeWalker{
		Entity: wasm.New(raw),
	}
}

// Root Returns the root node.
func (t *TreeWalker) Root() *Node {
	return NewNode(t.Js().Get("root"))
}

// WhatToShow Returns the whatToShow value.
func (t *TreeWalker) WhatToShow() NodeFilter {
	return NodeFilter(t.Js().Get("whatToShow").Int())
}

// Filter Returns the filter.
func (t *TreeWalker) Filter() *js.Func {
	filter := t.Js().Get("filter")

	if filter.IsNull() || filter.IsUndefined() || filter.Type() != js.TypeFunction {
		return nil
	}

	return &js.Func{Value: filter}
}

// CurrentNode Returns the current node.
func (t *TreeWalker) CurrentNode() *Node {
	return NewNode(t.Js().Get("currentNode"))
}

// SetCurrentNode Sets the current node.
func (t *TreeWalker) SetCurrentNode(node *Node) {
	t.Js().Set("currentNode", node.Js())
}

// ParentNode Returns the parent node.
func (t *TreeWalker) ParentNode() *Node {
	return NewNode(t.Js().Call("parentNode"))
}

// FirstChild Returns the first child.
func (t *TreeWalker) FirstChild() *Node {
	return NewNode(t.Js().Call("firstChild"))
}

// LastChild Returns the last child.
func (t *TreeWalker) LastChild() *Node {
	return NewNode(t.Js().Call("lastChild"))
}

// PreviousSibling Returns the previous sibling.
func (t *TreeWalker) PreviousSibling() *Node {
	return NewNode(t.Js().Call("previousSibling"))
}

// NextSibling Returns the next sibling.
func (t *TreeWalker) NextSibling() *Node {
	return NewNode(t.Js().Call("nextSibling"))
}

// PreviousNode Returns the previous node.
func (t *TreeWalker) PreviousNode() *Node {
	return NewNode(t.Js().Call("previousNode"))
}

// NextNode Returns the next node.
func (t *TreeWalker) NextNode() *Node {
	return NewNode(t.Js().Call("nextNode"))
}
