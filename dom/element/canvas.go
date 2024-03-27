package element

import (
	"go-webgl/dom/element/ctx"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/HTMLElement

type CanvasElement struct {
	*HTMLElement
}

func NewCanvasElement(element *HTMLElement) *CanvasElement {
	if element == nil {
		return nil
	}

	return &CanvasElement{
		HTMLElement: element,
	}
}

//
// CanvasElement methods
//

// Width A long representing the width of the canvas in coordinate space units.
func (d *CanvasElement) Width() int {
	return d.Js().Get("width").Int()
}

func (d *CanvasElement) SetWidth(value int) {
	d.Js().Set("width", value)
}

// Height A long representing the height of the canvas in coordinate space units.
func (d *CanvasElement) Height() int {
	return d.Js().Get("height").Int()
}

func (d *CanvasElement) SetHeight(value int) {
	d.Js().Set("height", value)
}

func (d *CanvasElement) Size() (int, int) {
	return d.Width(), d.Height()
}

func (d *CanvasElement) SetSize(width, height int) {
	if width != 0 {
		d.SetWidth(width)
	}

	if height != 0 {
		d.SetHeight(height)
	}
}

// getContext Returns a drawing context on the canvas, or null if the context identifier is not supported.
func (d *CanvasElement) GetContext(ctx ctx.Type) js.Value {
	return d.Js().Call("getContext", string(ctx))
}

// GetContext2d Returns a drawing context on the canvas, or null if the context identifier is not supported.
func (d *CanvasElement) GetContext2d() *ctx.Context2D {
	return ctx.NewContext2D(d.GetContext(ctx.Ctx2D))
}
