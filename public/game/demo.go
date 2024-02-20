package main

import (
    "go-webgl/webgl"
    "syscall/js"
)

func run() {
    canvas := js.Global().Get("document").Call("getElementById", "glcanvas")

    gl, err := webgl.New(canvas)

    if err != nil {
        panic(err)
    }

    width := gl.Canvas.ClientWidth()
    height := gl.Canvas.ClientHeight()

    gl.Viewport(0, 0, width, height)

    GlGameLoop(gl)

}

func GlGameLoop(gl *webgl.WebGL) bool {
    gl.ClearColor(0.5, 0.5, 0.5, 0.9)
    gl.Clear(gl.COLOR_BUFFER_BIT)

    return true
}

func main() {
    go run()
    select {}
}
