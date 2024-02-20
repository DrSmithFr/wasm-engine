package browser

import (
	"syscall/js"
)

func LoadDOM() *DOM {
	// loading default DOM data
	window := js.Global()
	document := window.Get("document")
	body := document.Get("body")

	// loading DOM size data
	size := Size{
		Height: window.Get("innerHeight").Float(),
		Width:  window.Get("innerWidth").Float(),
	}

	// returning DOM
	return &DOM{
		Window:   window,
		Document: document,
		Body:     body,
		Size:     size,
	}
}

type DOM struct {
	Window   js.Value
	Document js.Value
	Body     js.Value
	Canvas   js.Value
	Ball     js.Value
	Size     Size
}

type Size struct {
	Width  float64
	Height float64
}

func (dom *DOM) Log(args ...interface{}) {
	dom.Window.Get("console").Call("log", args...)
}
