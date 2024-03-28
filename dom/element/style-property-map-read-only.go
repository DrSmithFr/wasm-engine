package element

import (
	"syscall/js"
)

// StylePropertyMap https://developer.mozilla.org/en-US/docs/Web/API/StylePropertyMap
type StylePropertyMapReadOnly struct {
	raw js.Value
}

func newStylePropertyMapReadOnly(raw js.Value) *StylePropertyMapReadOnly {
	return &StylePropertyMapReadOnly{raw: raw}
}

// Get Returns the value of a property.
func (m StylePropertyMapReadOnly) Get(name string) string {
	return m.raw.Call("get", name).String()
}
