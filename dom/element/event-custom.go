package element

import "syscall/js"

// CustomEvent https://developer.mozilla.org/en-US/docs/Web/API/CustomEvent
type CustomEvent struct {
	*Event
}

// Detail Returns any custom data event was created with.
func (e *CustomEvent) Detail() js.Value {
	return e.Js().Get("detail")
}

// InitCustomEvent Initializes a CustomEvent object.
func (e *CustomEvent) InitCustomEvent(typeArg string, canBubbleArg bool, cancelableArg bool, detailArg js.Value) {
	e.Js().Call("initCustomEvent", typeArg, canBubbleArg, cancelableArg, detailArg)
}
