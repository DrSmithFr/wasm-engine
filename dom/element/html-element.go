package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

type HTMLElement struct {
	element *Element
}

func NewHTMLElement(raw js.Value) *HTMLElement {
	return &HTMLElement{element: newElement(raw)}
}

func NewHTMLElementList(items js.Value) []*HTMLElement {
	elements := make([]*HTMLElement, items.Length())

	for i := 0; i < items.Length(); i++ {
		elements[i] = NewHTMLElement(items.Index(i))
	}

	return elements
}

// enforce interface compliance
var _ wasm.WASM = (*HTMLElement)(nil)

func (d *HTMLElement) Js() js.Value {
	return d.element.Js()
}

//
// HTMLElement methods
//

func (d *HTMLElement) Element() *Element {
	return d.element
}

// AccessKey A DOMString representing the access key assigned to the element.
func (d *HTMLElement) AccessKey() string {
	return d.Js().Get("accessKey").String()
}

func (d *HTMLElement) SetAccessKey(value string) {
	d.Js().Set("accessKey", value)
}

// AccessKeyLabel A DOMString representing the element's assigned access key.
func (d *HTMLElement) AccessKeyLabel() string {
	return d.Js().Get("accessKeyLabel").String()
}

// AttributeStyleMap Returns a StylePropertyMap object for the element's inline style attribute.
func (d *HTMLElement) AttributeStyleMap() *StylePropertyMap {
	return newStylePropertyMap(d.Js().Get("attributeStyleMap"))
}

// Autofocus A Boolean indicating whether or not the element is focused.
func (d *HTMLElement) Autofocus() bool {
	return d.Js().Get("autofocus").Bool()
}

func (d *HTMLElement) SetAutofocus(value bool) {
	d.Js().Set("autofocus", value)
}

// ContentEditable A DOMString indicating whether or not the content of the element can be edited.
func (d *HTMLElement) ContentEditable() bool {
	return d.Js().Get("contentEditable").String() == "true"
}

func (d *HTMLElement) SetContentEditable(value bool) {
	v := "false"

	if value {
		v = "true"
	}

	d.Js().Set("contentEditable", v)
}

// Style A CSSStyleDeclaration object that represents the element's style attribute.
func (d *HTMLElement) Style() *CSSStyleDeclaration {
	return newCSSStyleDeclaration(d.Js().Get("style"))
}
