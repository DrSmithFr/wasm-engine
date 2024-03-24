package node

import (
	"syscall/js"
)

//
// Node helpers
//

// ChildNodes Returns a live NodeList containing all the children of this node (including elements,
// text and comments). NodeList being live means that if the children of the Node change,
// the NodeList object is automatically updated.
func ChildNodes(node js.Value) []js.Value {
	children := node.Get("childNodes")

	var elements []js.Value
	for i := 0; i < children.Length(); i++ {
		elem := children.Index(i)
		elements = append(elements, elem)
	}

	return elements
}

// FirstChild Returns the first child of a node.
func FirstChild(node js.Value) js.Value {
	return node.Get("firstChild")
}

// IsConnected GetAttribute Returns the value of a specified attribute on the element.
// If the given attribute does not exist, the value returned will either be null or "" (the empty string);
func IsConnected(node js.Value, elem js.Value) bool {
	return node.Call("isConnected", elem).Bool()
}

// LastChild Returns the last child of a node.
func LastChild(node js.Value) js.Value {
	return node.Get("lastChild")
}

// NextSibling Returns the node immediately following the specified one in its parent's childNodes list,
// or null if the specified node is the last node in that list.
func NextSibling(node js.Value) js.Value {
	return node.Get("nextSibling")
}

// NodeName ParentElement Returns the parent element of the specified element.
func NodeName(node js.Value) string {
	return node.Get("nodeName").String()
}

// NodeType Returns an unsigned short representing the type of the node.
func NodeType(node js.Value) Types {
	return Types(node.Get("nodeType").Int())
}

// NodeValue Returns or sets the value of the current node.
func NodeValue(node js.Value) string {
	return node.Get("nodeValue").String()
}

func SetNodeValue(node js.Value, value string) {
	node.Set("nodeValue", value)
}

// OwnerDocument Returns the Document that this node belongs to.
func OwnerDocument(node js.Value) js.Value {
	return node.Get("ownerDocument")
}

// ParentNode Returns the parent of the specified node in the DOM tree.
func ParentNode(node js.Value) js.Value {
	return node.Get("parentNode")
}

// ParentElement Returns the parent element of the specified element.
func ParentElement(node js.Value) js.Value {
	return node.Get("parentElement")
}

// PreviousSibling Returns the node immediately before the specified one in its parent's childNodes list,
// or null if the specified node is the first in that list.
func PreviousSibling(node js.Value) js.Value {
	return node.Get("previousSibling")
}

// TextContent Represents the text content of a node and its descendants.
func TextContent(node js.Value) string {
	return node.Get("textContent").String()
}

func SetTextContent(node js.Value, text string) {
	node.Set("textContent", text)
}

// AppendChild Adds the specified childNode argument as the last child to the current node.
// If the argument referenced an existing node on the DOM tree,
// the node will be detached from its current position and attached at the new position.
func AppendChild(node js.Value, child js.Value) {
	node.Call("appendChild", child)
}

// CloneNode Returns a duplicate of the node on which this method was called.
func CloneNode(node js.Value, deep bool) js.Value {
	return node.Call("cloneNode", deep)
}

// CompareDocumentPosition Returns an unsigned short, which is a bitmask,
// representing the position of the reference node relative to the node on which this method is called.
func CompareDocumentPosition(node js.Value, other js.Value) int {
	return node.Call("compareDocumentPosition", other).Int()
}

// Contains Returns a Boolean value indicating whether a node is a descendant of a given node or not.
func Contains(node js.Value, other js.Value) bool {
	return node.Call("contains", other).Bool()
}

// GetRootNode Returns the context object's shadow-including root.
func GetRootNode(node js.Value) js.Value {
	return node.Call("getRootNode")
}

// HasChildNodes Returns a Boolean value indicating whether the current Node has child nodes or not.
func HasChildNodes(node js.Value) bool {
	return node.Call("hasChildNodes").Bool()
}

// InsertBefore Inserts a node before a reference node as a child of a specified parent node.
func InsertBefore(node js.Value, newNode js.Value, referenceNode js.Value) {
	node.Call("insertBefore", newNode, referenceNode)
}

// IsDefaultNamespace Returns a Boolean value indicating whether or not the namespace of the node is the default namespace for the document.
func IsDefaultNamespace(node js.Value, namespace string) bool {
	return node.Call("isDefaultNamespace", namespace).Bool()
}

// IsEqualNode Returns a Boolean value indicating whether the two nodes are equal.
func IsEqualNode(node js.Value, other js.Value) bool {
	return node.Call("isEqualNode", other).Bool()
}

// IsSameNode Returns a Boolean value indicating whether two nodes are the same node.
func IsSameNode(node js.Value, other js.Value) bool {
	return node.Call("isSameNode", other).Bool()
}

// LookupPrefix Returns a string containing the prefix for a given namespace URI, if present, and null if not.
// When multiple prefixes are possible, the result is implementation-dependent.
func LookupPrefix(node js.Value, namespace string) string {
	return node.Call("lookupPrefix", namespace).String()
}

// LookupNamespaceURI Returns a string containing the namespace URI for a given prefix, if present, and null if not.
func LookupNamespaceURI(node js.Value, prefix string) string {
	return node.Call("lookupNamespaceURI", prefix).String()
}

// Normalize Puts the specified node and all of its subtree into a "normalized" form.
// In a normalized subtree, no text nodes in the subtree are empty and there are no adjacent text nodes.
func Normalize(node js.Value) {
	node.Call("normalize")
}

// RemoveChild Removes a child node from the DOM and returns the removed node.
func RemoveChild(node js.Value, child js.Value) js.Value {
	return node.Call("removeChild", child)
}

// ReplaceChild Replaces one child node of the specified node with another.
func ReplaceChild(node js.Value, newChild js.Value, oldChild js.Value) {
	node.Call("replaceChild", newChild, oldChild)
}
