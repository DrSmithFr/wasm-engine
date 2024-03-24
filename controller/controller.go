package controller

import (
	"go-webgl/dom/element"
)

type ActionState struct {
	Up        bool
	Down      bool
	Left      bool
	Right     bool
	TurnLeft  bool
	TurnRight bool

	ShowMap bool
	LockMap bool
	Action  bool
}

type Interface interface {
	Init(window *element.Window)
	GetState() ActionState
}
