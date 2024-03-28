package element

import (
	"syscall/js"
)

// DOMRect https://developer.mozilla.org/en-US/docs/Web/API/DOMRect
type DOMRect struct {
	*DOMRectReadOnly
}

func NewDOMRect(raw js.Value) *DOMRect {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &DOMRect{
		DOMRectReadOnly: NewDOMRectReadOnly(raw),
	}
}

// FromRect Creates a new DOMRect object with a given position and size.
func FromRect(other *DOMRectReadOnly) *DOMRect {
	elem := js.Global().Get("DOMRect").New(other.Js())
	return NewDOMRect(elem)
}

// SetX Sets the x-coordinate of the DOMRect.
func (d *DOMRect) SetX(x float64) *DOMRect {
	d.Js().Set("x", x)
	return d
}

// SetY Sets the y-coordinate of the DOMRect.
func (d *DOMRect) SetY(y float64) *DOMRect {
	d.Js().Set("y", y)
	return d
}

// SetWidth Sets the width of the DOMRect.
func (d *DOMRect) SetWidth(width float64) *DOMRect {
	d.Js().Set("width", width)
	return d
}

// SetHeight Sets the height of the DOMRect.
func (d *DOMRect) SetHeight(height float64) *DOMRect {
	d.Js().Set("height", height)
	return d
}

// SetTop Sets the top of the DOMRect.
func (d *DOMRect) SetTop(top float64) *DOMRect {
	d.Js().Set("top", top)
	return d
}

// SetRight Sets the right of the DOMRect.
func (d *DOMRect) SetRight(right float64) *DOMRect {
	d.Js().Set("right", right)
	return d
}

// SetBottom Sets the bottom of the DOMRect.
func (d *DOMRect) SetBottom(bottom float64) *DOMRect {
	d.Js().Set("bottom", bottom)
	return d
}

// SetLeft Sets the left of the DOMRect.
func (d *DOMRect) SetLeft(left float64) *DOMRect {
	d.Js().Set("left", left)
	return d
}
