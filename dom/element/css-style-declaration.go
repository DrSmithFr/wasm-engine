package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type CSSStyleDeclaration struct {
	*wasm.Entity
}

func NewCSSStyleDeclaration(raw js.Value) *CSSStyleDeclaration {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &CSSStyleDeclaration{
		Entity: wasm.New(raw),
	}
}

//
// CSSStyleDeclaration methods
//

// Length Returns the number of properties.
func (d *CSSStyleDeclaration) Length() int {
	return d.Js().Get("length").Int()
}

// GetPropertyPriority Returns the optional priority, "important".
func (d *CSSStyleDeclaration) GetPropertyPriority(property string) string {
	return d.Js().Call("getPropertyPriority", property).String()
}

// GetPropertyValue Returns the property value.
func (d *CSSStyleDeclaration) GetPropertyValue(property string) string {
	return d.Js().Call("getPropertyValue", property).String()
}

// Item Returns the property name.
func (d *CSSStyleDeclaration) Item(index int) string {
	return d.Js().Call("item", index).String()
}

// RemoveProperty Removes the property.
func (d *CSSStyleDeclaration) RemoveProperty(property string) string {
	return d.Js().Call("removeProperty", property).String()
}

// SetProperty Sets the property.
func (d *CSSStyleDeclaration) SetProperty(property, value string, important bool) *CSSStyleDeclaration {
	priority := ""

	if important {
		priority = "important"
	}

	d.Js().Call("setProperty", property, value, priority)

	return d
}
