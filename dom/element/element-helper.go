package element

import "syscall/js"

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Element

// AssignedSlot Returns a HTMLSlotElement representing the <slot> the node is inserted in.
func AssignedSlot(elem js.Value) js.Value {
	return elem.Get("assignedSlot")
}

// Attributes Returns a NamedNodeMap object containing the assigned attributes of the corresponding HTML element.
func Attributes(elem js.Value) js.Value {
	return elem.Get("attributes")
}

// ChildElementCount Returns the number of child elements of this element.
func ChildElementCount(elem js.Value) int {
	return elem.Get("childElementCount").Int()
}

// Children Returns the child elements of this element.
func Children(elem js.Value) js.Value {
	return elem.Get("children")
}

// ClassName A string representing the class of the element.
func ClassName(elem js.Value) js.Value {
	return elem.Get("className")
}

func SetClassName(elem js.Value, className string) {
	elem.Set("className", className)
}

// ClientHeight Returns a number representing the inner height of the element.
func ClientHeight(elem js.Value) int {
	return elem.Get("clientHeight").Int()
}

// ClientLeft Returns a number representing the width of the left border of the element.
func ClientLeft(elem js.Value) int {
	return elem.Get("clientLeft").Int()
}

// ClientTop Returns a number representing the width of the top border of the element.
func ClientTop(elem js.Value) int {
	return elem.Get("clientTop").Int()
}

// ClientWidth Returns a number representing the inner width of the element.
func ClientWidth(elem js.Value) int {
	return elem.Get("clientWidth").Int()
}

// ElementTiming A string reflecting the elementtiming attribute which
// marks an element for observation in the PerformanceElementTiming API.
func ElementTiming(elem js.Value) js.Value {
	return elem.Get("elementTiming")
}

// FirstElementChild Returns the first child element of this element.
func FirstElementChild(elem js.Value) js.Value {
	return elem.Get("firstElementChild")
}

// Id A string representing the id of the element.
func Id(elem js.Value) string {
	return elem.Get("id").String()
}

func SetId(elem js.Value, id string) {
	elem.Set("id", id)
}

// InnerHTML A DOMString representing the markup of the element's content.
func InnerHTML(elem js.Value) string {
	return elem.Get("innerHTML").String()
}

func SetInnerHTML(elem js.Value, innerHTML string) {
	elem.Set("innerHTML", innerHTML)
}

// LastElementChild Returns the last child element of this element.
func LastElementChild(elem js.Value) js.Value {
	return elem.Get("lastElementChild")
}

// LocalName A string representing the local part of the qualified name of the element.
func LocalName(elem js.Value) string {
	return elem.Get("localName").String()
}

// NamespaceURI A string representing the namespace URI of the element.
func NamespaceURI(elem js.Value) string {
	return elem.Get("namespaceURI").String()
}

// NextElementSibling Returns the element immediately following the specified one in its parent's children list,
// or null if the specified element is the last one in the list.
func NextElementSibling(elem js.Value) js.Value {
	return elem.Get("nextElementSibling")
}

// OuterHTML A DOMString representing the element and its contents.
func OuterHTML(elem js.Value) string {
	return elem.Get("outerHTML").String()
}

func SetOuterHTML(elem js.Value, outerHTML string) {
	elem.Set("outerHTML", outerHTML)
}
