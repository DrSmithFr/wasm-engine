package render

import (
	"go-webgl/dom/element"
	"image/color"
)

type RenderFn func(r Renderer) bool

type Renderer interface {
	Init(window *element.Window)

	SetSize(width, height int)
	Size() (int, int)

	SetFPS(maxFPS int)
	Start(maxFPS int, render RenderFn)
	Stop()

	SetColor(c color.RGBA)
	DrawCircle(x, y, width float64)
	DrawLine(x1, y1, x2, y2, width float64)
	DrawRect(x1, y1, x2, y2 float64)

	GetCanvas() *element.CanvasElement

	Clear() // Clear the canvas
	Flush() // set image as rendered
}
