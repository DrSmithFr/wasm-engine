package canvas

import (
    "fmt"
    "syscall/js"
)

type Canvas2d struct {
    done chan struct{} // Used as part of 'run forever' in the render handler

    // DOM properties
    window js.Value
    doc    js.Value
    body   js.Value

    // canvas properties
    canvas *js.Value

    // size properties
    width  int
    height int

    // position properties
    top    *int
    left   *int
    bottom *int
    right  *int
    zIndex *int

    // extra css properties
    css map[string]string
}

func New2d(create bool) (*Canvas2d, error) {

    var c Canvas2d

    c.window = js.Global()
    c.doc = c.window.Get("document")
    c.body = c.doc.Call("getElementById", "main")

    // If create, make a canvas that fills the windows
    if create {
        c.Create(
            c.window.Get("innerWidth").Int(),
            c.window.Get("innerHeight").Int(),
        )
    }

    c.css = make(map[string]string)

    return &c, nil
}

// Implements the Canvas interface
var _ Canvas = (*Canvas2d)(nil)

// Create a new canvas in the DOM, and append it to the Body.
// This also calls Bind to create relevant shadow Buffer etc

// TODO suspect this needs to be fleshed out with more options
func (c *Canvas2d) Create(width int, height int) {

    // Make the canvas
    canvas := c.doc.Call("createElement", "canvas")

    canvas.Set("height", height)
    canvas.Set("width", width)

    c.body.Call("appendChild", canvas)

    c.Bind(&canvas, width, height)
}

// Used to setup with an existing Canvas element which was obtained from JS
func (c *Canvas2d) Bind(canvas *js.Value, width int, height int) {
    c.canvas = canvas
    c.height = height
    c.width = width
}

func (c *Canvas2d) SetSize(width int, height int) {
    c.height = height
    c.width = width

    c.canvas.Set("height", height)
    c.canvas.Set("width", width)
}

func (c *Canvas2d) Size() (int, int) {
    return c.width, c.height
}

func (c *Canvas2d) Js() *js.Value {
    return c.canvas
}

func (c *Canvas2d) SetZIndex(z int) {
    c.zIndex = &z
}

func (c *Canvas2d) SetTop(x int) {
    c.top = &x
    c.UpdateCanvasStyle()
}

func (c *Canvas2d) SetLeft(x int) {
    c.left = &x
    c.UpdateCanvasStyle()
}

func (c *Canvas2d) SetBottom(x int) {
    c.bottom = &x
    c.UpdateCanvasStyle()
}

func (c *Canvas2d) SetRight(x int) {
    c.right = &x
    c.UpdateCanvasStyle()
}

func (c *Canvas2d) SetCssProperty(property, value string) {
    old := c.css[property]

    if old != value {
        c.css[property] = value
        c.UpdateCanvasStyle()
    }
}

func (c *Canvas2d) UpdateCanvasStyle() {
    style := "position: absolute;"

    for property, value := range c.css {
        style += fmt.Sprintf("%s: %s;", property, value)
    }

    c.canvas.Set("style", style)
}
