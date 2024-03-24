package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type StylePropertyMap struct {
	raw js.Value
}

func newStylePropertyMap(raw js.Value) *StylePropertyMap {
	return &StylePropertyMap{raw: raw}
}

// enforce interface compliance
var _ wasm.WASM = (*StylePropertyMap)(nil)

func (m StylePropertyMap) Js() js.Value {
	return m.raw
}

//
// StylePropertyMap methods
//

// Append Adds a property to the map, or updates it if it already exists.
func (m StylePropertyMap) Append(name string, value string) {
	m.raw.Call("append", name, value)
}

// Clear Removes all properties from the map.
func (m StylePropertyMap) Clear() {
	m.raw.Call("clear")
}

// Delete Removes a property from the map.
func (m StylePropertyMap) Delete(name string) {
	m.raw.Call("delete", name)
}

// Set Sets the value of a property.
func (m StylePropertyMap) Set(name string, value string) {
	m.raw.Call("set", name, value)
}

// Get Returns the value of a property.
func (m StylePropertyMap) Get(name string) string {
	return m.raw.Call("get", name).String()
}
