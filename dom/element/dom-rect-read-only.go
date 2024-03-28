package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// DOMRectReadOnly https://developer.mozilla.org/en-US/docs/Web/API/DOMRectReadOnly
type DOMRectReadOnly struct {
	*wasm.Entity
}

func NewDOMRectReadOnly(raw js.Value) *DOMRectReadOnly {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &DOMRectReadOnly{
		Entity: wasm.New(raw),
	}
}

// X Returns the x-coordinate of the DOMRectReadOnly.
func (d *DOMRectReadOnly) X() float64 {
	return d.Js().Get("x").Float()
}

// Y Returns the y-coordinate of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Y() float64 {
	return d.Js().Get("y").Float()
}

// Width Returns the width of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Width() float64 {
	return d.Js().Get("width").Float()
}

// Height Returns the height of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Height() float64 {
	return d.Js().Get("height").Float()
}

// Top Returns the top of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Top() float64 {
	return d.Js().Get("top").Float()
}

// Right Returns the right of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Right() float64 {
	return d.Js().Get("right").Float()
}

// Bottom Returns the bottom of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Bottom() float64 {
	return d.Js().Get("bottom").Float()
}

// Left Returns the left of the DOMRectReadOnly.
func (d *DOMRectReadOnly) Left() float64 {
	return d.Js().Get("left").Float()
}
