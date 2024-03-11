package browser

import (
    "syscall/js"
)

func Load() Document {
    // loading default Document data
    window := js.Global()
    document := window.Get("document")
    body := document.Get("body")

    // loading Document size data
    size := Size{
        Height: window.Get("innerHeight").Float(),
        Width:  window.Get("innerWidth").Float(),
    }

    // returning Document
    dom := Document{
        Window:   window,
        Document: document,
        Body:     body,
        Size:     size,
    }

    return dom
}

type Document struct {
    Window   js.Value
    Document js.Value
    Body     js.Value
    Size     Size
}

type Size struct {
    Width  float64
    Height float64
}

func (dom *Document) Log(args ...interface{}) {
    dom.Window.Get("console").Call("log", args...)
}

func (dom *Document) GetScreenSize() (int, int) {
    w := js.Global().Get("innerWidth").Int()
    h := js.Global().Get("innerHeight").Int()
    return w, h
}
