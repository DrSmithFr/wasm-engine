package render

import (
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
	canvas  *element.CanvasElement
	ctx     *ctx.Context2D
	imgData *ctx.ImageData

	// position properties
	width  int
	height int

	// RequestAnimationFrame
	done     chan struct{} // Used as part of 'run forever' in the render handler
	window   js.Value
	reqID    js.Value // Storage of the current annimationFrame requestID - For Cancel
	timeStep float64  // Min Time delay between frames. - Calculated as   maxFPS/1000

	// Drawing Context
	gctx     *draw2dimg.GraphicContext // Graphic Context
	image    *image.RGBA               // The Shadow frame we actually draw on
	font     *truetype.Font
	fontData draw2d.FontData
}

func (r *WasmBuffered) GetCanvas() *element.CanvasElement {
	return r.canvas
}

func NewWasmBuffered(c *element.CanvasElement, width, height int) *WasmBuffered {
	return &WasmBuffered{
		canvas: c,
		ctx:    c.GetContext2d(),

		window: js.Global(),
		done:   make(chan struct{}),

		width:  width,
		height: height,
	}
}

// implement Renderer interface
var _ Renderer = (*WasmBuffered)(nil)

func (r *WasmBuffered) Init(window *element.Window) {
	r.SetSize(r.width, r.height)
}

func (r *WasmBuffered) SetSize(width int, height int) {
	// Setup the 2D Drawing context
	r.canvas.SetSize(width, height)
	r.imgData = r.ctx.CreateImageData(r.width, r.height) // Note Width, then Height

	r.image = image.NewRGBA(image.Rect(0, 0, width, height))
	r.gctx = draw2dimg.NewGraphicContext(r.image)
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
	r.done <- struct{}{}
	close(r.done)
}

func (r *WasmBuffered) Clear() {
	r.gctx.Clear()
}

func (r *WasmBuffered) Flush() {

}

func (r *WasmBuffered) SetColor(c color.RGBA) {
	r.gctx.SetStrokeColor(c)
	r.gctx.SetFillColor(c)
}

func (r *WasmBuffered) DrawCircle(x, y, width float64) {
	radius := width / 2

	r.gctx.BeginPath()
	r.gctx.SetLineWidth(1)
	r.gctx.ArcTo(x, y, radius, radius, 0, -math.Pi*2)
	r.gctx.Fill()
	r.gctx.Close()
}

func (r *WasmBuffered) DrawLine(x1, y1, x2, y2, width float64) {
	r.gctx.BeginPath()
	r.gctx.SetLineWidth(width)
	r.gctx.MoveTo(x1, y1)
	r.gctx.LineTo(x2, y2)
	r.gctx.Stroke()
}

func (r *WasmBuffered) DrawRect(x1, y1, x2, y2 float64) {
	r.gctx.BeginPath()
	r.gctx.MoveTo(x1, y1)
	r.gctx.LineTo(x2, y1)
	r.gctx.LineTo(x2, y2)
	r.gctx.LineTo(x1, y2)
	r.gctx.Close()
	r.gctx.FillStroke()
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
				if renderingFn != nil { // If required, call the requested render function, before copying the frame
					if renderingFn(r) { // Only copy the image back if RenderFunction returns TRUE. (i.e. stuff has changed.)  This allows Render to return false, saving time this cycle if nothing changed.  (Keep frame as before)
						r.imgCopy()
					}
				} else { // Just do the copy, rendering must be being done elsewhere
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
	r.imgData.SetData(r.image)
	r.ctx.PutImageData(r.imgData, 0, 0)
}
