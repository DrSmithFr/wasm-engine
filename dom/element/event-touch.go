package element

import "syscall/js"

// TouchEvent https://developer.mozilla.org/en-US/docs/Web/API/TouchEvent
type TouchEvent struct {
	*Event
}

func NewTouchEvent(entity js.Value) *TouchEvent {
	return &TouchEvent{
		Event: NewEvent(entity),
	}
}

// AltKey Returns a Boolean that is true if the alt key was down when the touch event was fired.
func (e *TouchEvent) AltKey() bool {
	return e.Js().Get("altKey").Bool()
}

// ChangedTouches Returns a list of all the Touch objects for touch points that have changed since the event was last fired.
func (e *TouchEvent) ChangedTouches() []*Touch {
	var touches []*Touch
	for i := 0; i < e.Js().Get("changedTouches").Length(); i++ {
		touches = append(touches, NewTouch(e.Js().Get("changedTouches").Index(i)))
	}
	return touches
}

// CtrlKey Returns a Boolean that is true if the control key was down when the touch event was fired.
func (e *TouchEvent) CtrlKey() bool {
	return e.Js().Get("ctrlKey").Bool()
}

// MetaKey Returns a Boolean that is true if the meta key was down when the touch event was fired.
func (e *TouchEvent) MetaKey() bool {
	return e.Js().Get("metaKey").Bool()
}

// ShiftKey Returns a Boolean that is true if the shift key was down when the touch event was fired.
func (e *TouchEvent) ShiftKey() bool {
	return e.Js().Get("shiftKey").Bool()
}

// TargetTouches Returns a list of all the Touch objects for touch points that are still in contact with the surface and where the touchstart event occurred on the same target element as the current target element.
func (e *TouchEvent) TargetTouches() []*Touch {
	var touches []*Touch
	for i := 0; i < e.Js().Get("targetTouches").Length(); i++ {
		touches = append(touches, NewTouch(e.Js().Get("targetTouches").Index(i)))
	}
	return touches
}

// Touches Returns a list of all the Touch objects representing all current points of contact with the surface, regardless of target or changed status.
func (e *TouchEvent) Touches() []*Touch {
	var touches []*Touch
	for i := 0; i < e.Js().Get("touches").Length(); i++ {
		touches = append(touches, NewTouch(e.Js().Get("touches").Index(i)))
	}
	return touches
}
