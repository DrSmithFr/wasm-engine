package render

import (
	"fmt"
	"go-webgl/dom/element"
	"go-webgl/dom/element/ctx"
	"image/color"
	"math"
	"syscall/js"
)

type DirectCtx struct {
	// canvas properties
	canvas *element.CanvasElement
	ctx    *ctx.Context2D

	// position properties
	width  int
	height int

	// RequestAnimationFrame
	done     chan struct{} // Used as part of 'run forever' in the render handler
	window   js.Value
	reqID    js.Value // Storage of the current annimationFrame requestID - For Cancel
	timeStep float64  // Min Time delay between frames. - Calculated as   maxFPS/1000
}

func NewDirectCtx(c *element.CanvasElement) *DirectCtx {
	return &DirectCtx{
		canvas: c,
		ctx:    c.GetContext2d(),
		window: js.Global(),
		done:   make(chan struct{}),
	}
}

// implement Renderer interface
var _ Renderer = (*DirectCtx)(nil)

func (r *DirectCtx) GetCanvas() *element.CanvasElement {
	return r.canvas
}

func (r *DirectCtx) Init(window *element.Window) {
	if r.width == 0 {
		r.width = r.canvas.Js().Get("width").Int()
	}

	if r.height == 0 {
		r.height = r.canvas.Js().Get("height").Int()
	}

	r.SetSize(r.width, r.height)
}

func (r *DirectCtx) SetSize(width int, height int) {
	r.canvas.SetSize(width, height)
}

func (r *DirectCtx) Size() (int, int) {
	return r.canvas.Size()
}

func (r *DirectCtx) SetFPS(maxFPS int) {
	r.timeStep = 1000. / float64(maxFPS)
}

func (r *DirectCtx) Start(maxFPS int, renderFn RenderFn) {
	r.SetFPS(maxFPS)
	r.initFrameUpdate(renderFn)
}

func (r *DirectCtx) Stop() {
	r.window.Call("cancelAnimationFrame", r.reqID)
	r.done <- struct{}{}
	close(r.done)
}

func (r *DirectCtx) Clear() {
	r.ctx.ClearRect(0, 0, r.width, r.height)
}

func (r *DirectCtx) Flush() {
}

func (r *DirectCtx) SetColor(c color.RGBA) {
	cHex := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	r.ctx.SetFillStyle(cHex)
	r.ctx.SetStrokeStyle(cHex)
}

func (r *DirectCtx) DrawCircle(x, y, width float64) {
	r.ctx.BeginPath()

	r.ctx.Arc(
		math.Round(x),
		math.Round(y),
		math.Round(width),
		0,
		math.Pi*2,
		true,
	)

	r.ctx.Fill()
	r.ctx.ClosePath()
}

func (r *DirectCtx) DrawLine(x1, y1, x2, y2, width float64) {
	r.ctx.BeginPath()
	r.ctx.LineWidth(int(math.Round(width)))

	r.ctx.MoveTo(int(math.Round(x1)), int(math.Round(y1)))
	r.ctx.LineTo(int(math.Round(x2)), int(math.Round(y2)))

	r.ctx.Stroke()
	r.ctx.ClosePath()

	r.ctx.LineWidth(1)
}

func (r *DirectCtx) DrawRect(x1, y1, x2, y2 float64) {
	r.ctx.BeginPath()

	r.ctx.MoveTo(int(math.Round(x1)), int(math.Round(y1)))
	r.ctx.LineTo(int(math.Round(x2)), int(math.Round(y1)))
	r.ctx.LineTo(int(math.Round(x2)), int(math.Round(y2)))
	r.ctx.LineTo(int(math.Round(x1)), int(math.Round(y2)))

	r.ctx.Fill()
	r.ctx.ClosePath()
}

// handles calls from Render, and copies the image over.
func (r *DirectCtx) initFrameUpdate(renderingFn RenderFn) {
	// Hold the callbacks without blocking
	go func() {
		var renderFrame js.Func
		var lastTimestamp float64

		renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			timestamp := args[0].Float()
			if timestamp-lastTimestamp >= r.timeStep { // Constrain FPS
				renderingFn(r)
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
