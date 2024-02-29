package controller

import (
    "go-webgl/browser"
    "syscall/js"
)

type KeyboardOnly struct {
    move ActionState
}

func NewKeyboardOnly() *KeyboardOnly {
    return &KeyboardOnly{
        move: ActionState{
            Up:        false,
            Down:      false,
            Left:      false,
            Right:     false,
            TurnLeft:  false,
            TurnRight: false,
        },
    }
}

// implement Interface interface
var _ Interface = (*KeyboardOnly)(nil)

func (k *KeyboardOnly) Init(dom browser.DOM) {
    k.bindEvents(dom)
}

func (k *KeyboardOnly) GetState() ActionState {
    return k.move
}

// internal methods

func (k *KeyboardOnly) bindEvents(dom browser.DOM) {
    // let's handle key Down
    var keydownEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        k.keydownEvent(args[0])
        return nil
    })

    var keyupEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        k.keyupEvent(args[0])
        return nil
    })

    dom.Document.Call("addEventListener", "keydown", keydownEventHandler)
    dom.Document.Call("addEventListener", "keyup", keyupEventHandler)
}

func (k *KeyboardOnly) keydownEvent(event js.Value) {
    code := event.Get("code").String()

    switch code {
    case "ArrowUp", "KeyW":
        k.move.Up = true
    case "ArrowDown", "KeyS":
        k.move.Down = true
    case "KeyQ":
        k.move.Left = true
    case "KeyE":
        k.move.Right = true
    case "ArrowRight", "KeyD":
        k.move.TurnRight = true
    case "ArrowLeft", "KeyA":
        k.move.TurnLeft = true
    case "KeyT":
        k.move.Action = true
    }
}

func (k *KeyboardOnly) keyupEvent(event js.Value) {
    code := event.Get("code").String()

    switch code {
    case "ArrowUp", "KeyW":
        k.move.Up = false
    case "ArrowDown", "KeyS":
        k.move.Down = false
    case "KeyQ":
        k.move.Left = false
    case "KeyE":
        k.move.Right = false
    case "ArrowRight", "KeyD":
        k.move.TurnRight = false
    case "ArrowLeft", "KeyA":
        k.move.TurnLeft = false
    case "KeyT":
        k.move.Action = false
    }
}
