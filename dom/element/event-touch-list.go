package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// TouchList https://developer.mozilla.org/en-US/docs/Web/API/TouchList
type TouchList struct {
	*wasm.Entity
}

func NewTouchList(entity js.Value) *TouchList {
	if entity.IsNull() || entity.IsUndefined() {
		return nil
	}

	return &TouchList{
		Entity: wasm.New(entity),
	}
}

// Length Returns the number of Touch objects in the TouchList.
func (t *TouchList) Length() int {
	return t.Js().Get("length").Int()
}

// Item Returns the Touch object at the specified index in the TouchList.
func (t *TouchList) Item(index int) *Touch {
	return NewTouch(t.Js().Call("item", index))
}

// Touch https://developer.mozilla.org/en-US/docs/Web/API/Touch
type Touch struct {
	*wasm.Entity
}

func NewTouch(entity js.Value) *Touch {
	if entity.IsNull() || entity.IsUndefined() {
		return nil
	}

	return &Touch{
		Entity: wasm.New(entity),
	}
}

// Identifier Returns a unique identifier for this Touch object.
func (t *Touch) Identifier() int {
	return t.Js().Get("identifier").Int()
}

// ScreenX Returns the X coordinate of the touch point relative to the screen.
func (t *Touch) ScreenX() int {
	return t.Js().Get("screenX").Int()
}

// ScreenY Returns the Y coordinate of the touch point relative to the screen.
func (t *Touch) ScreenY() int {
	return t.Js().Get("screenY").Int()
}

// ClientX Returns the X coordinate of the touch point relative to the viewport, excluding any scroll offset.
func (t *Touch) ClientX() int {
	return t.Js().Get("clientX").Int()
}

// ClientY Returns the Y coordinate of the touch point relative to the viewport, excluding any scroll offset.
func (t *Touch) ClientY() int {
	return t.Js().Get("clientY").Int()
}

// PageX Returns the X coordinate of the touch point relative to the viewport, including any scroll offset.
func (t *Touch) PageX() int {
	return t.Js().Get("pageX").Int()
}

// PageY Returns the Y coordinate of the touch point relative to the viewport, including any scroll offset.
func (t *Touch) PageY() int {
	return t.Js().Get("pageY").Int()
}

// Target Returns the EventTarget on which the touch point started when it was created.
func (t *Touch) Target() *Element {
	return NewElement(t.Js().Get("target"))
}

// RadiusX Returns the radius of the ellipse that most closely circumscribes the area of contact with the screen.
func (t *Touch) RadiusX() float64 {
	return t.Js().Get("radiusX").Float()
}

// RadiusY Returns the radius of the ellipse that most closely circumscribes the area of contact with the screen.
func (t *Touch) RadiusY() float64 {
	return t.Js().Get("radiusY").Float()
}

// RotationAngle Returns the angle (in degrees) that the ellipse described by radiusX and radiusY must be rotated, clockwise, to most accurately cover the area of contact between the user and the screen.
func (t *Touch) RotationAngle() float64 {
	return t.Js().Get("rotationAngle").Float()
}

// Force Returns the amount of pressure being applied to the surface by the user, as a float between 0.0 (no pressure) and 1.0 (maximum pressure).
func (t *Touch) Force() float64 {
	return t.Js().Get("force").Float()
}
