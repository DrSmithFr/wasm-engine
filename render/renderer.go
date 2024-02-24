package render

import (
    "image/color"
)

type RenderFn func(r Renderer) bool

type Renderer interface {
    Init(width, height, x, y int)

    SetSize(width, height int)
    Size() (int, int)

    SetFPS(maxFPS int)
    Start(maxFPS int, render RenderFn)
    Stop()

    SetColor(c color.RGBA)
    DrawCircle(x, y, width float64)
    DrawLine(x1, y1, x2, y2, width float64)
    DrawRect(x1, y1, x2, y2 float64)

    Clear() // Clear the canvas
    Flush() // set image as rendered
}
