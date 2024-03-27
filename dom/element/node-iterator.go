package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type NodeIterator struct {
	*wasm.Entity
}

func NewNodeIterator(raw js.Value) *NodeIterator {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &NodeIterator{
		Entity: wasm.New(raw),
	}
}

// Root Returns the root node.
func (n *NodeIterator) Root() *Node {
	return NewNode(n.Js().Get("root"))
}

// WhatToShow Returns the whatToShow value.
func (n *NodeIterator) WhatToShow() NodeFilter {
	return NodeFilter(n.Js().Get("whatToShow").Int())
}

// Filter Returns the filter.
func (n *NodeIterator) Filter() *js.Func {
	filter := n.Js().Get("filter")

	if filter.IsNull() || filter.IsUndefined() || filter.Type() != js.TypeFunction {
		return nil
	}

	return &js.Func{Value: filter}
}

// ReferenceNode Returns the reference node.
func (n *NodeIterator) ReferenceNode() *Node {
	return NewNode(n.Js().Get("referenceNode"))
}

// PointerBeforeReferenceNode Returns true if the pointer is before the reference node.
func (n *NodeIterator) PointerBeforeReferenceNode() bool {
	return n.Js().Get("pointerBeforeReferenceNode").Bool()
}

// Detach Detaches the node iterator.
func (n *NodeIterator) Detach() {
	n.Js().Call("detach")
}

// PreviousNode Returns the previous node.
func (n *NodeIterator) PreviousNode() *Node {
	return NewNode(n.Js().Call("previousNode"))
}

// NextNode Returns the next node.
func (n *NodeIterator) NextNode() *Node {
	return NewNode(n.Js().Call("nextNode"))
}
