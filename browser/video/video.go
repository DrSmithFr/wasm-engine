package video

import (
    "go-webgl/browser"
    "syscall/js"
)

type Video struct {
    Dom browser.Element
}

func New(src string) *Video {
    v := &Video{
        Dom: browser.Element{
            Js:   js.Global().Get("document").Call("createElement", "video"),
            Size: browser.Size{},
            Css:  browser.Css{},
        },
    }

    v.SetSrc(src)

    return v
}

func Bind(elementJs js.Value) *Video {
    return &Video{
        Dom: browser.Element{
            Js: elementJs,
            Size: browser.Size{
                Width:  elementJs.Get("width").Float(),
                Height: elementJs.Get("height").Float(),
            },
            Css: browser.Css{},
        },
    }
}

func BindById(id string) *Video {
    return Bind(js.Global().Get("document").Call("getElementById", id))
}

func (v *Video) SetSrc(src string) {
    v.Dom.Js.Set("src", src)
}

type Player interface {
    Play()
    Pause()
    Stop()
}

func (v *Video) Play() {
    v.Dom.Js.Call("play")
}

func (v *Video) Pause() {
    v.Dom.Js.Call("pause")
}

func (v *Video) Stop() {
    v.Dom.Js.Call("stop")
}
