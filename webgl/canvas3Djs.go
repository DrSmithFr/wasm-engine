package webgl

import (
    "syscall/js"
)

type Canvas3D struct {
    done chan struct{} // Used as part of 'run forever' in the render handler

    Canvas Canvas
    gl     WebGL

    // DOM properties
    window js.Value

    reqID    js.Value // Storage of the current annimationFrame requestID - For Cancel
    timeStep float64  // Min Time delay between frames. - Calculated as   maxFPS/1000
}

func NewCanvas3D(c Canvas, gl WebGL) *Canvas3D {
    return &Canvas3D{
        Canvas: c,
        gl:     gl,
    }
}

type RenderFunc func(gl *WebGL) bool

// Starts the annimationFrame callbacks running.   (Recently seperated from Create / Set to give better control for when things start / stop)
func (c Canvas3D) Start(maxFPS float64, rf RenderFunc) {
    c.SetFPS(maxFPS)
    c.initFrameUpdate(rf)
}

// This needs to be called on an 'beforeUnload' trigger, to properly close out the render callback, and prevent browser errors on page Refresh
func (c Canvas3D) Stop() {
    c.window.Call("cancelAnimationFrame", c.reqID)
    c.done <- struct{}{}
    close(c.done)
}

// Sets the maximum FPS (Frames per Second).  This can be changed on the fly and will take affect next frame.
func (c Canvas3D) SetFPS(maxFPS float64) {
    c.timeStep = 1000 / maxFPS
}

// handles calls from Render, and copies the image over.
func (c *Canvas3D) initFrameUpdate(rf RenderFunc) {
    // Hold the callbacks without blocking
    go func() {
        var renderFrame js.Func
        var lastTimestamp float64

        renderFrame = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

            timestamp := args[0].Float()
            if timestamp-lastTimestamp >= c.timeStep { // Constrain FPS
                if rf != nil { // If required, call the requested render function, before copying the frame
                    rf(&c.gl)
                } else { // Just do the copy, rendering must be being done elsewhere
                    panic("runner is nil")
                }
                lastTimestamp = timestamp
            }

            c.reqID = js.Global().Call("requestAnimationFrame", renderFrame) // Captures the requestID to be used in Close / Cancel
            return nil
        })
        defer renderFrame.Release()
        js.Global().Call("requestAnimationFrame", renderFrame)
        <-c.done
    }()
}
