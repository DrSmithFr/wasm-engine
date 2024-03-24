package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type CSSStyleDeclaration struct {
	raw js.Value
}

func newCSSStyleDeclaration(raw js.Value) *CSSStyleDeclaration {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &CSSStyleDeclaration{raw: raw}
}

// enforce interface compliance
var _ wasm.WASM = (*CSSStyleDeclaration)(nil)

func (d *CSSStyleDeclaration) Js() js.Value {
	return d.raw
}

//
// CSSStyleDeclaration methods
//

// Length Returns the number of properties.
func (d *CSSStyleDeclaration) Length() int {
	return d.raw.Get("length").Int()
}

// GetPropertyPriority Returns the optional priority, "important".
func (d *CSSStyleDeclaration) GetPropertyPriority(property string) string {
	return d.raw.Call("getPropertyPriority", property).String()
}

// GetPropertyValue Returns the property value.
func (d *CSSStyleDeclaration) GetPropertyValue(property string) string {
	return d.raw.Call("getPropertyValue", property).String()
}

// Item Returns the property name.
func (d *CSSStyleDeclaration) Item(index int) string {
	return d.raw.Call("item", index).String()
}

// RemoveProperty Removes the property.
func (d *CSSStyleDeclaration) RemoveProperty(property string) string {
	return d.raw.Call("removeProperty", property).String()
}

// SetProperty Sets the property.
func (d *CSSStyleDeclaration) SetProperty(property, value string, important bool) *CSSStyleDeclaration {
	priority := ""

	if important {
		priority = "important"
	}

	d.raw.Call("setProperty", property, value, priority)

	return d
}
