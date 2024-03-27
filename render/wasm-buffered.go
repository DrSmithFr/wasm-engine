package render

import (
	"fmt"
	"github.com/golang/freetype/truetype"
	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
	"go-webgl/dom/element"
	"go-webgl/dom/element/ctx"
	"image"
	"image/color"
	"math"
	"syscall/js"
)

type WasmBuffered struct {
	// canvas properties
	canvas *element.CanvasElement
	ctx    *draw2dimg.GraphicContext
	ctx2D  *ctx.Context2D

	// Drawing Context
	image   *image.RGBA // The Shadow frame we actually draw on
	imgData js.Value
	buffer  js.Value

	// Font Context
	font     *truetype.Font
	fontData draw2d.FontData

	// position properties
	width  int
	height int

	// RequestAnimationFrame
	done     chan bool // Used as part of 'run forever' in the render handler
	window   js.Value
	reqID    js.Value // Storage of the current annimationFrame requestID - For Cancel
	timeStep float64  // Min Time delay between frames. - Calculated as   maxFPS/1000

}

func (r *WasmBuffered) GetCanvas() *element.CanvasElement {
	return r.canvas
}

func NewWasmBuffered(c *element.CanvasElement) *WasmBuffered {
	return &WasmBuffered{
		canvas: c,
		ctx2D:  c.GetContext2d(),
		window: js.Global(),
		done:   make(chan bool),
	}
}

// implement Renderer interface
var _ Renderer = (*WasmBuffered)(nil)

func (r *WasmBuffered) Init(window *element.Window) {
	if r.width == 0 {
		r.width = window.InnerWidth()
	}

	if r.height == 0 {
		r.height = window.InnerHeight()
	}

	r.SetSize(r.width, r.height)
}

func (r *WasmBuffered) SetSize(width int, height int) {
	fmt.Print("reszing")
	r.canvas.SetSize(width, height)

	r.image = image.NewRGBA(image.Rect(0, 0, width, height))
	r.imgData = r.ctx2D.CreateImageData(width, height)             // Note Width, then Height
	r.buffer = js.Global().Get("Uint8Array").New(len(r.image.Pix)) // Static JS drawCtx for copying data out to JS. Defined once and re-used to save on un-needed allocations

	r.ctx = draw2dimg.NewGraphicContext(r.image)
}

func (r *WasmBuffered) Size() (int, int) {
	return r.canvas.Size()
}

func (r *WasmBuffered) SetFPS(maxFPS int) {
	r.timeStep = 1000. / float64(maxFPS)
}

func (r *WasmBuffered) Start(maxFPS int, renderFn RenderFn) {
	r.SetFPS(maxFPS)
	r.initFrameUpdate(renderFn)
}

func (r *WasmBuffered) Stop() {
	r.window.Call("cancelAnimationFrame", r.reqID)
	r.done <- true
	close(r.done)
}

func (r *WasmBuffered) Clear() {
	r.ctx.Clear()
	r.SetColor(color.RGBA{0, 0, 0, 0})
	r.DrawRect(0, 0, float64(r.width), float64(r.height))
}

func (r *WasmBuffered) Flush() {

}

func (r *WasmBuffered) SetColor(c color.RGBA) {
	r.ctx.SetStrokeColor(c)
	r.ctx.SetFillColor(c)
}

func (r *WasmBuffered) DrawCircle(x, y, width float64) {
	radius := width / 2

	r.ctx.BeginPath()
	r.ctx.SetLineWidth(1)
	r.ctx.ArcTo(x, y, radius, radius, 0, -math.Pi*2)
	r.ctx.Fill()
	r.ctx.FillStroke()
}

func (r *WasmBuffered) DrawLine(x1, y1, x2, y2, width float64) {
	r.ctx.BeginPath()
	r.ctx.SetLineWidth(width)
	r.ctx.MoveTo(x1, y1)
	r.ctx.LineTo(x2, y2)
	r.ctx.FillStroke()
}

func (r *WasmBuffered) DrawRect(x1, y1, x2, y2 float64) {
	r.ctx.BeginPath()
	r.ctx.MoveTo(x1, y1)
	r.ctx.LineTo(x2, y1)
	r.ctx.LineTo(x2, y2)
	r.ctx.LineTo(x1, y2)
	r.ctx.Close()
	r.ctx.FillStroke()
}

// internals

// handles calls from Render, and copies the image over.
func (r *WasmBuffered) initFrameUpdate(renderingFn RenderFn) {
	// Hold the callbacks without blocking
	go func() {
		var renderFrame js.Func
		var lastTimestamp float64

		renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			timestamp := args[0].Float()
			if timestamp-lastTimestamp >= r.timeStep { // Constrain FPS
				if renderingFn(r) { // Only copy the image back if RenderFunction returns TRUE. (i.e. stuff has changed.)  This allows Render to return false, saving time this cycle if nothing changed.  (Keep frame as before)
					r.imgCopy()
				}
				lastTimestamp = timestamp
			}

			r.reqID = js.Global().Call("requestAnimationFrame", renderFrame) // Captures the requestID to be used in Close / Cancel
			return nil
		})
		defer renderFrame.Release()
		js.Global().Call("requestAnimationFrame", renderFrame)
		<-r.done
	}()
}

// Does the actually copy over of the image data for the 'render' call.
func (r *WasmBuffered) imgCopy() {
	js.CopyBytesToJS(r.buffer, r.image.Pix)
	r.imgData.Get("data").Call("set", r.buffer)
	r.ctx2D.PutImageData(r.imgData, 0, 0)
}
