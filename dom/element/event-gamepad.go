package element

import "syscall/js"

// GamepadEvent https://developer.mozilla.org/en-US/docs/Web/API/GamepadEvent
type GamepadEvent struct {
	*Event
}

// Gamepad Returns a Gamepad object representing the gamepad that was connected or disconnected.
// todo
func (e *GamepadEvent) Gamepad() js.Value {
	return e.Js().Get("gamepad")
}
