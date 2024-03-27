package element

import (
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

type HTMLElement struct {
	*Element
	css *CSSStyleDeclaration
}

func NewHTMLElement(raw js.Value) *HTMLElement {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &HTMLElement{
		Element: NewElement(raw),
		css:     newCSSStyleDeclaration(raw.Get("style")),
	}
}

//
// HTMLElement methods
//

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
