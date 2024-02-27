package render

import (
    "fmt"
    "go-webgl/browser"
    "go-webgl/canvas"
    "image/color"
    "math"
    "syscall/js"
)

type DirectCtx struct {
    // canvas properties
    canvas canvas.Canvas
    ctx    js.Value

    // position properties
    width  int
    height int
    x      int
    y      int

    // RequestAnimationFrame
    done     chan struct{} // Used as part of 'run forever' in the render handler
    window   js.Value
    reqID    js.Value // Storage of the current annimationFrame requestID - For Cancel
    timeStep float64  // Min Time delay between frames. - Calculated as   maxFPS/1000
}

func NewDirectCtx(width, height, x, y int) *DirectCtx {
    c, err := canvas.New2d(true)

    if err != nil {
        panic(err)
    }

    return &DirectCtx{
        canvas: c,
        window: js.Global(),
        done:   make(chan struct{}),

        width:  width,
        height: height,
        x:      x,
        y:      y,
    }
}

// implement Renderer interface
var _ Renderer = (*DirectCtx)(nil)

func (r *DirectCtx) Init(dom browser.DOM) {
    c, err := canvas.New2d(false)

    if err != nil {
        panic(err)
    }

    c.Create(r.width, r.height)
    r.ctx = r.canvas.Js().Call("getContext", "2d")
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
    r.ctx.Call("clearRect", 0, 0, r.width, r.height)
}

func (r *DirectCtx) Flush() {

}

func (r *DirectCtx) SetColor(c color.RGBA) {
    cHex := fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
    r.ctx.Set("fillStyle", cHex)
    r.ctx.Set("strokeStyle", cHex)
}

func (r *DirectCtx) DrawCircle(x, y, width float64) {
    r.ctx.Call("beginPath")
    r.ctx.Call("arc", math.Round(x), math.Round(y), math.Round(width), 0, math.Pi*2, true)
    r.ctx.Call("fill")
    r.ctx.Call("closePath")
}

func (r *DirectCtx) DrawLine(x1, y1, x2, y2, width float64) {
    r.ctx.Call("beginPath")
    r.ctx.Set("lineWidth", int64(math.Round(width)))
    r.ctx.Call("moveTo", math.Round(x1), math.Round(y1))
    r.ctx.Call("lineTo", math.Round(x2), math.Round(y2))
    r.ctx.Call("stroke")
    r.ctx.Call("closePath")
    r.ctx.Set("lineWidth", 1)
}

func (r *DirectCtx) DrawRect(x1, y1, x2, y2 float64) {
    r.ctx.Call("beginPath")
    r.ctx.Call("moveTo", math.Round(x1), math.Round(y1))
    r.ctx.Call("lineTo", math.Round(x2), math.Round(y1))
    r.ctx.Call("lineTo", math.Round(x2), math.Round(y2))
    r.ctx.Call("lineTo", math.Round(x1), math.Round(y2))
    r.ctx.Call("fill")
    r.ctx.Call("closePath")
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
