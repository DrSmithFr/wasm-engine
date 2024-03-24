package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Window

type Window struct {
	raw      js.Value
	document *Document
}

// CurrentWindow enforce document Singleton
var CurrentWindow *Window

func LoadWindow() *Window {
	if CurrentWindow != nil {
		return CurrentWindow
	}

	CurrentWindow = &Window{raw: js.Global()}
	CurrentWindow.document = LoadDocument()

	return CurrentWindow
}

// enforce interface compliance
var _ wasm.WASM = (*Window)(nil)

func (d *Window) Bind(e js.Value) {
	d.raw = e
}

func (d *Window) Js() js.Value {
	return d.raw
}

//
// Window methods
//

// DevicePixelRatio Returns the ratio of the resolution in physical pixels to the resolution in CSS pixels for the current display device.
func (d *Window) DevicePixelRatio() float64 {
	return d.raw.Get("devicePixelRatio").Float()
}

// Document Returns the Document object that represents the document in the specified window.
func (d *Window) Document() *Document {
	return d.document
}

// InnerWidth Returns the width of the content area of the browser window including, if rendered, the vertical scrollbar.
func (d *Window) InnerWidth() int {
	return d.raw.Get("innerWidth").Int()
}

// InnerHeight Returns the height of the content area of the browser window including, if rendered, the horizontal scrollbar.
func (d *Window) InnerHeight() int {
	return d.raw.Get("innerHeight").Int()
}

func (d *Window) InnerSize() (int, int) {
	return d.InnerWidth(), d.InnerHeight()
}

// OuterWidth Returns the width of the outside of the browser window.
func (d *Window) OuterWidth() int {
	return d.raw.Get("outerWidth").Int()
}

// OuterHeight Returns the height of the outside of the browser window.
func (d *Window) OuterHeight() int {
	return d.raw.Get("outerHeight").Int()
}

func (d *Window) OuterSize() (int, int) {
	return d.OuterWidth(), d.OuterHeight()
}

// PageXOffset Returns the number of pixels that the document has already been scrolled horizontally.
func (d *Window) PageXOffset() int {
	return d.raw.Get("pageXOffset").Int()
}

// PageYOffset Returns the number of pixels that the document has already been scrolled vertically.
func (d *Window) PageYOffset() int {
	return d.raw.Get("pageYOffset").Int()
}

func (d *Window) PageOffset() (int, int) {
	return d.PageXOffset(), d.PageYOffset()
}
