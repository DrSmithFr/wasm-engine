package element

import (
	"syscall/js"
)

type Node struct {
	*EventTarget
}

func NewNode(raw js.Value) *Node {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Node{
		EventTarget: NewEventTarget(raw),
	}
}

//
// Node helpers
//

// ChildNodes Returns a live NodeList containing all the children of this node (including elements,
// text and comments). NodeList being live means that if the children of the Node change,
// the NodeList object is automatically updated.
func (n *Node) ChildNodes() []*Node {
	children := n.Js().Get("childNodes")

	var elements []*Node
	for i := 0; i < children.Length(); i++ {
		elem := children.Index(i)
		node := NewNode(elem)
		elements = append(elements, node)
	}

	return elements
}

// FirstChild Returns the first child of a node.
func (n *Node) FirstChild() *Node {
	return NewNode(n.Js().Get("firstChild"))
}

// IsConnected GetAttribute Returns the value of a specified attribute on the element.
// If the given attribute does not exist, the value returned will either be null or "" (the empty string);
func (n *Node) IsConnected(elem *Node) bool {
	return n.Js().Call("isConnected", elem).Bool()
}

// LastChild Returns the last child of a node.
func (n *Node) LastChild() *Node {
	return NewNode(n.Js().Get("lastChild"))
}

// NextSibling Returns the node immediately following the specified one in its parent's childNodes list,
// or null if the specified node is the last node in that list.
func (n *Node) NextSibling() *Node {
	return NewNode(n.Js().Get("nextSibling"))
}

// NodeName ParentElement Returns the parent element of the specified element.
func (n *Node) NodeName() string {
	return n.Js().Get("nodeName").String()
}

// NodeType Returns an unsigned short representing the type of the node.
func (n *Node) NodeType() NodeType {
	return NodeType(n.Js().Get("nodeType").Int())
}

// NodeValue Returns or sets the value of the current node.
func (n *Node) NodeValue() string {
	return n.Js().Get("nodeValue").String()
}

func (n *Node) SetNodeValue(value string) {
	n.Js().Set("nodeValue", value)
}

// OwnerDocument Returns the Document that this node belongs to.
func (n *Node) OwnerDocument() *Document {
	return NewDocument(n.Js().Get("ownerDocument"))
}

// ParentNode Returns the parent of the specified node in the DOM tree.
func (n *Node) ParentNode() *Node {
	return NewNode(n.Js().Get("parentNode"))
}

// ParentElement Returns the parent element of the specified element.
func (n *Node) ParentElement() *Node {
	return NewNode(n.Js().Get("parentElement"))
}

// PreviousSibling Returns the node immediately before the specified one in its parent's childNodes list,
// or null if the specified node is the first in that list.
func (n *Node) PreviousSibling() *Node {
	return NewNode(n.Js().Get("previousSibling"))
}

// TextContent Represents the text content of a node and its descendants.
func (n *Node) TextContent() string {
	return n.Js().Get("textContent").String()
}

func (n *Node) SetTextContent(text string) {
	n.Js().Set("textContent", text)
}

// AppendChild Adds the specified childNode argument as the last child to the current node.
// If the argument referenced an existing node on the DOM tree,
// the node will be detached from its current position and attached at the new position.
func (n *Node) AppendChild(child *Node) {
	n.Js().Call("appendChild", child.Js())
}

// CloneNode Returns a duplicate of the node on which this method was called.
func (n *Node) CloneNode(deep bool) *Node {
	return NewNode(n.Js().Call("cloneNode", deep))
}

// CompareDocumentPosition Returns an unsigned short, which is a bitmask,
// representing the position of the reference node relative to the node on which this method is called.
func (n *Node) CompareDocumentPosition(other *Node) int {
	return n.Js().Call("compareDocumentPosition", other.Js()).Int()
}

// Contains Returns a Boolean value indicating whether a node is a descendant of a given node or not.
func (n *Node) Contains(other *Node) bool {
	return n.Js().Call("contains", other.Js()).Bool()
}

// GetRootNode Returns the context object's shadow-including root.
func (n *Node) GetRootNode() *Node {
	return NewNode(n.Js().Call("getRootNode"))
}

// HasChildNodes Returns a Boolean value indicating whether the current Node has child nodes or not.
func (n *Node) HasChildNodes() bool {
	return n.Js().Call("hasChildNodes").Bool()
}

// InsertBefore Inserts a node before a reference node as a child of a specified parent node.
func (n *Node) InsertBefore(newNode *Node, referenceNode *Node) {
	n.Js().Call("insertBefore", newNode.Js(), referenceNode.Js())
}

// IsDefaultNamespace Returns a Boolean value indicating whether or not the namespace of the node is the default namespace for the document.
func (n *Node) IsDefaultNamespace(namespace string) bool {
	return n.Js().Call("isDefaultNamespace", namespace).Bool()
}

// IsEqualNode Returns a Boolean value indicating whether the two nodes are equal.
func (n *Node) IsEqualNode(other *Node) bool {
	return n.Js().Call("isEqualNode", other.Js()).Bool()
}

// IsSameNode Returns a Boolean value indicating whether two nodes are the same node.
func (n *Node) IsSameNode(other *Node) bool {
	return n.Js().Call("isSameNode", other.Js()).Bool()
}

// LookupPrefix Returns a string containing the prefix for a given namespace URI, if present, and null if not.
// When multiple prefixes are possible, the result is implementation-dependent.
func (n *Node) LookupPrefix(namespace string) string {
	return n.Js().Call("lookupPrefix", namespace).String()
}

// LookupNamespaceURI Returns a string containing the namespace URI for a given prefix, if present, and null if not.
func (n *Node) LookupNamespaceURI(prefix string) string {
	return n.Js().Call("lookupNamespaceURI", prefix).String()
}

// Normalize Puts the specified node and all of its subtree into a "normalized" form.
// In a normalized subtree, no text nodes in the subtree are empty and there are no adjacent text nodes.
func (n *Node) Normalize() {
	n.Js().Call("normalize")
}

// RemoveChild Removes a child node from the DOM and returns the removed node.
func (n *Node) RemoveChild(child *Node) *Node {
	return NewNode(n.Js().Call("removeChild", child.Js()))
}

// ReplaceChild Replaces one child node of the specified node with another.
func (n *Node) ReplaceChild(newChild *Node, oldChild *Node) {
	n.Js().Call("replaceChild", newChild.Js(), oldChild.Js())
}
