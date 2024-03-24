package element

import (
	"go-webgl/dom/node"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Element

type Element struct {
	raw js.Value
}

func newElement(raw js.Value) *Element {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Element{raw: raw}
}

// enforce interface compliance
var _ wasm.WASM = (*Element)(nil)

func (d *Element) Js() js.Value {
	return d.raw
}

//
// HTMLElement methods
//

// AssignedSlot Returns a HTMLSlotElement representing the <slot> the node is inserted in.
func (d *Element) AssignedSlot() js.Value {
	return AssignedSlot(d.raw)
}

// Attributes Returns a NamedNodeMap object containing the assigned attributes of the corresponding HTML element.
func (d *Element) Attributes() js.Value {
	return Attributes(d.raw)
}

// ChildElementCount Returns the number of child elements of this element.
func (d *Element) ChildElementCount() int {
	return ChildElementCount(d.raw)
}

// Children Returns the child elements of this element.
func (d *Element) Children() *Element {
	return newElement(Children(d.raw))
}

// ClassName A string representing the class of the element.
func (d *Element) ClassName() string {
	return ClassName(d.raw).String()
}

func (d *Element) SetClassName(className string) {
	SetClassName(d.raw, className)
}

// ClientHeight Returns a number representing the inner height of the element.
func (d *Element) ClientHeight() int {
	return ClientHeight(d.raw)
}

// ClientLeft Returns a number representing the width of the left border of the element.
func (d *Element) ClientLeft() int {
	return ClientLeft(d.raw)
}

// ClientTop Returns a number representing the width of the top border of the element.
func (d *Element) ClientTop() int {
	return ClientTop(d.raw)
}

// ClientWidth Returns a number representing the inner width of the element.
func (d *Element) ClientWidth() int {
	return ClientWidth(d.raw)
}

// ElementTiming A string reflecting the elementtiming attribute which
// marks an element for observation in the PerformanceElementTiming API.
func (d *Element) ElementTiming() js.Value {
	return ElementTiming(d.raw)
}

// FirstElementChild Returns the first child element of this element.
func (d *Element) FirstElementChild() *Element {
	return newElement(FirstElementChild(d.raw))
}

// Id A string representing the id of the element.
func (d *Element) Id() string {
	return Id(d.raw)
}

func (d *Element) SetId(id string) {
	SetId(d.raw, id)
}

// InnerHTML A DOMString representing the markup of the element's content.
func (d *Element) InnerHTML() string {
	return InnerHTML(d.raw)
}

func (d *Element) SetInnerHTML(html string) {
	SetInnerHTML(d.raw, html)
}

// LastElementChild Returns the last child element of this element.
func (d *Element) LastElementChild() *Element {
	return newElement(LastElementChild(d.raw))
}

// NextElementSibling Returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
func (d *Element) NextElementSibling() *Element {
	return newElement(NextElementSibling(d.raw))
}

// LocalName A string representing the local part of the qualified name of the element.
func (d *Element) LocalName() string {
	return LocalName(d.raw)
}

// NamespaceURI A string representing the namespace URI of the element.
func (d *Element) NamespaceURI() string {
	return NamespaceURI(d.raw)
}

// OuterHTML A DOMString representing the element and its contents.
func (d *Element) OuterHTML() string {
	return OuterHTML(d.raw)
}

func (d *Element) SetOuterHTML(html string) {
	SetOuterHTML(d.raw, html)
}

//
// Node methods
//

// FirstChild Returns the first child of a node.
func (d *Element) FirstChild() *Element {
	return newElement(node.FirstChild(d.raw))
}

// IsConnected Returns a Boolean value indicating whether or not a node is connected to the DOM.
func (d *Element) IsConnected(elem wasm.WASM) bool {
	return node.IsConnected(d.raw, elem.Js())
}

// LastChild Returns the last child of a node.
func (d *Element) LastChild() *Element {
	return newElement(node.LastChild(d.raw))
}

// NextSibling Returns the node immediately following the specified one in its parent's childNodes list, or null if the specified node is the last node in that list.
func (d *Element) NextSibling() *Element {
	return newElement(node.NextSibling(d.raw))
}

// NodeName Returns a string representing the name of the node in the form of a tag name.
func (d *Element) NodeName() string {
	return node.NodeName(d.raw)
}

// NodeType Returns an unsigned short representing the type of the node.
func (d *Element) NodeType() node.Types {
	return node.NodeType(d.raw)
}

// NodeValue Sets or returns the value of the current node.
func (d *Element) NodeValue() string {
	return node.NodeValue(d.raw)
}

func (d *Element) SetNodeValue(value string) {
	node.SetNodeValue(d.raw, value)
}

// OwnerDocument Returns the Document object associated with a node.
func (d *Element) OwnerDocument() *Element {
	return newElement(node.OwnerDocument(d.raw))
}

// ParentNode Returns the parent of the specified node in the DOM tree.
func (d *Element) ParentNode() *Element {
	return newElement(node.ParentNode(d.raw))
}

// ParentElement Returns the parent element of the specified element.
func (d *Element) ParentElement() *Element {
	return newElement(node.ParentElement(d.raw))
}

// PreviousSibling Returns the node immediately preceding the specified one in its parent's childNodes list, or null if the specified node is the first in that list.
func (d *Element) PreviousSibling() *Element {
	return newElement(node.PreviousSibling(d.raw))
}

// TextContent Sets or returns the textual content of a node and its descendants.
func (d *Element) TextContent() string {
	return node.TextContent(d.raw)
}

func (d *Element) SetTextContent(value string) {
	node.SetTextContent(d.raw, value)
}

// AppendChild Adds a node to the end of the list of children of a specified parent node.
func (d *Element) AppendChild(child wasm.WASM) {
	node.AppendChild(d.raw, child.Js())
}

// CloneNode Returns a duplicate of the node on which this method was called.
func (d *Element) CloneNode(deep bool) *Element {
	return newElement(node.CloneNode(d.raw, deep))
}

// CompareDocumentPosition Returns an unsigned short representing the relationship between the node and the reference node.
func (d *Element) CompareDocumentPosition(other wasm.WASM) int {
	return node.CompareDocumentPosition(d.raw, other.Js())
}

// Contains Returns a Boolean value indicating whether a node is a descendant of a given node.
func (d *Element) Contains(other wasm.WASM) bool {
	return node.Contains(d.raw, other.Js())
}

// GetRootNode Returns the topmost node in the DOM tree.
func (d *Element) GetRootNode() *Element {
	return newElement(node.GetRootNode(d.raw))
}

// HasChildNodes Returns a Boolean value indicating whether the element has any child nodes or not.
func (d *Element) HasChildNodes() bool {
	return node.HasChildNodes(d.raw)
}

// InsertBefore Inserts a node before a reference node as a child of a specified parent node.
func (d *Element) InsertBefore(child wasm.WASM, reference wasm.WASM) {
	node.InsertBefore(d.raw, child.Js(), reference.Js())
}

// IsDefaultNamespace Returns a Boolean value indicating whether or not the document has a namespace.
func (d *Element) IsDefaultNamespace(namespace string) bool {
	return node.IsDefaultNamespace(d.raw, namespace)
}

// IsEqualNode Returns a Boolean value indicating whether or not two nodes are of the same type and all their defining data are equal.
func (d *Element) IsEqualNode(other wasm.WASM) bool {
	return node.IsEqualNode(d.raw, other.Js())
}

// IsSameNode Returns a Boolean value indicating whether or not the two nodes are the same (that is, they reference the same object).
func (d *Element) IsSameNode(other wasm.WASM) bool {
	return node.IsSameNode(d.raw, other.Js())
}

// LookupPrefix Returns a DOMString containing the prefix for a given namespace URI, if present, and null if not.
func (d *Element) LookupPrefix(namespace string) string {
	return node.LookupPrefix(d.raw, namespace)
}

// LookupNamespaceURI Returns a DOMString representing the namespace prefix associated with the given namespace URI, if present, and null if not.
func (d *Element) LookupNamespaceURI(prefix string) string {
	return node.LookupNamespaceURI(d.raw, prefix)
}

// Normalize Puts the specified node and all of its subtree into a "normalized" form.
func (d *Element) Normalize() {
	node.Normalize(d.raw)
}

// RemoveChild Removes a child node from the DOM.
func (d *Element) RemoveChild(child wasm.WASM) {
	node.RemoveChild(d.raw, child.Js())
}

// ReplaceChild Replaces one child node of the specified element with another.
func (d *Element) ReplaceChild(newChild wasm.WASM, oldChild wasm.WASM) {
	node.ReplaceChild(d.raw, newChild.Js(), oldChild.Js())
}
