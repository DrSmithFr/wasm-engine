package controller

import (
    "go-webgl/browser"
    "log"
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
            ShowMap:   false,
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

    log.Printf("Keydown event: %s", code)

    switch code {
    case "ArrowUp", "KeyW":
        k.move.Up = true
    case "ArrowDown", "KeyS":
        k.move.Down = true
    case "ArrowLeft":
        k.move.Left = true
    case "ArrowRight":
        k.move.Right = true
    case "KeyD":
        k.move.TurnRight = true
    case "KeyA":
        k.move.TurnLeft = true
    case "KeyT":
        k.move.Action = true
    case "KeyM":
        k.move.ShowMap = !k.move.ShowMap
    case "Comma":
        k.move.LockMap = !k.move.LockMap
    }
}

func (k *KeyboardOnly) keyupEvent(event js.Value) {
    code := event.Get("code").String()

    switch code {
    case "ArrowUp", "KeyW":
        k.move.Up = false
    case "ArrowDown", "KeyS":
        k.move.Down = false
    case "ArrowLeft":
        k.move.Left = false
    case "ArrowRight":
        k.move.Right = false
    case "KeyD":
        k.move.TurnRight = false
    case "KeyA":
        k.move.TurnLeft = false
    case "KeyT":
        k.move.Action = false
    }
}
