package element

import "syscall/js"

type HTMLSlotElement struct {
	*HTMLElement
}

func NewHTMLSlotElement(raw js.Value) *HTMLSlotElement {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &HTMLSlotElement{
		HTMLElement: NewHTMLElement(raw),
	}
}

// Name Returns the name of the slot.
func (h *HTMLSlotElement) Name() string {
	return h.Js().Get("name").String()
}

// SetName Sets the name of the slot.
func (h *HTMLSlotElement) SetName(name string) {
	h.Js().Set("name", name)
}

// Assign Assigns a node to the slot.
func (h *HTMLSlotElement) Assign(node *Node) {
	h.Js().Call("assign", node.Js())
}

// AssignedNodes Returns the assigned nodes.
func (h *HTMLSlotElement) AssignedNodes(flatten bool) []*Node {
	rawNodes := h.Js().Call("assignedNodes", flatten)

	if rawNodes.IsNull() || rawNodes.IsUndefined() {
		panic("assignedNodes() returned null or undefined")
	}

	nodes := make([]*Node, rawNodes.Length())

	for i := 0; i < rawNodes.Length(); i++ {
		nodes[i] = NewNode(rawNodes.Index(i))
	}

	return nodes
}

// AssignedElements Returns the assigned elements.
func (h *HTMLSlotElement) AssignedElements(flatten bool) []*Element {
	rawNodes := h.Js().Call("assignedElements", flatten)

	if rawNodes.IsNull() || rawNodes.IsUndefined() {
		panic("assignedElements() returned null or undefined")
	}

	nodes := make([]*Element, rawNodes.Length())

	for i := 0; i < rawNodes.Length(); i++ {
		nodes[i] = NewElement(rawNodes.Index(i))
	}

	return nodes
}
