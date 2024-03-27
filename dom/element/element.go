package element

import (
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Element

type Element struct {
	*Node
}

func newElement(raw js.Value) *Element {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Element{
		Node: NewNode(raw),
	}
}

//
// HTMLElement methods
//

// AssignedSlot Returns a HTMLSlotElement representing the <slot> the node is inserted in.
func (d *Element) AssignedSlot() js.Value {
	return AssignedSlot(d.Js())
}

// Attributes Returns a NamedNodeMap object containing the assigned attributes of the corresponding HTML element.
func (d *Element) Attributes() js.Value {
	return Attributes(d.Js())
}

// ChildElementCount Returns the number of child elements of this element.
func (d *Element) ChildElementCount() int {
	return ChildElementCount(d.Js())
}

// Children Returns the child elements of this element.
func (d *Element) Children() *Element {
	return newElement(Children(d.Js()))
}

// ClassName A string representing the class of the element.
func (d *Element) ClassName() string {
	return ClassName(d.Js()).String()
}

func (d *Element) SetClassName(className string) {
	SetClassName(d.Js(), className)
}

// ClientHeight Returns a number representing the inner height of the element.
func (d *Element) ClientHeight() int {
	return ClientHeight(d.Js())
}

// ClientLeft Returns a number representing the width of the left border of the element.
func (d *Element) ClientLeft() int {
	return ClientLeft(d.Js())
}

// ClientTop Returns a number representing the width of the top border of the element.
func (d *Element) ClientTop() int {
	return ClientTop(d.Js())
}

// ClientWidth Returns a number representing the inner width of the element.
func (d *Element) ClientWidth() int {
	return ClientWidth(d.Js())
}

// ElementTiming A string reflecting the elementtiming attribute which
// marks an element for observation in the PerformanceElementTiming API.
func (d *Element) ElementTiming() js.Value {
	return ElementTiming(d.Js())
}

// FirstElementChild Returns the first child element of this element.
func (d *Element) FirstElementChild() *Element {
	return newElement(FirstElementChild(d.Js()))
}

// Id A string representing the id of the element.
func (d *Element) Id() string {
	return Id(d.Js())
}

func (d *Element) SetId(id string) {
	SetId(d.Js(), id)
}

// InnerHTML A DOMString representing the markup of the element's content.
func (d *Element) InnerHTML() string {
	return InnerHTML(d.Js())
}

func (d *Element) SetInnerHTML(html string) {
	SetInnerHTML(d.Js(), html)
}

// LastElementChild Returns the last child element of this element.
func (d *Element) LastElementChild() *Element {
	return newElement(LastElementChild(d.Js()))
}

// NextElementSibling Returns the element immediately following the specified one in its parent's children list, or null if the specified element is the last one in the list.
func (d *Element) NextElementSibling() *Element {
	return newElement(NextElementSibling(d.Js()))
}

// LocalName A string representing the local part of the qualified name of the element.
func (d *Element) LocalName() string {
	return LocalName(d.Js())
}

// NamespaceURI A string representing the namespace URI of the element.
func (d *Element) NamespaceURI() string {
	return NamespaceURI(d.Js())
}

// OuterHTML A DOMString representing the element and its contents.
func (d *Element) OuterHTML() string {
	return OuterHTML(d.Js())
}

func (d *Element) SetOuterHTML(html string) {
	SetOuterHTML(d.Js(), html)
}
