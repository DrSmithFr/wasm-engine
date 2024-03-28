package element

import "syscall/js"

// InputEvent https://developer.mozilla.org/en-US/docs/Web/API/InputEvent
type InputEvent struct {
	*Event
}

// Data Returns a DOMString representing the input data.
func (e *InputEvent) Data() string {
	return e.Js().Get("data").String()
}

// DataTransfer Returns a DataTransfer object representing the data being dragged during the drag operation.
func (e *InputEvent) DataTransfer() *DataTransfer {
	return NewDataTransfer(e.Js().Get("dataTransfer"))
}

// InputType Returns a DOMString representing the type of input control that was changed.
func (e *InputEvent) InputType() string {
	return e.Js().Get("inputType").String()
}

// IsComposing Returns a Boolean that is true if the event is fired after compositionstart and before compositionend.
func (e *InputEvent) IsComposing() bool {
	return e.Js().Get("isComposing").Bool()
}

// GetTargetRanges Returns a static Range object representing the range of text that is affected by the event.
func (e *InputEvent) GetTargetRanges() js.Value {
	return e.Js().Call("getTargetRanges")
}
