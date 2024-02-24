package controller

import "go-webgl/browser"

type ActionState struct {
    Up     bool
    Down   bool
    Left   bool
    Right  bool
    Action bool
}

type Interface interface {
    Init(dom browser.DOM)
    GetState() ActionState
}