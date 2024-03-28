package element

import "syscall/js"

// MouseEvent https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent
type MouseEvent struct {
	*Event
}

func NewMouseEvent(raw js.Value) *MouseEvent {
	return &MouseEvent{
		Event: NewEvent(raw),
	}
}

// AltKey Returns a Boolean that is true if the Alt (Option or ⌥ on OS X) key was active when the event was generated.
func (e *MouseEvent) AltKey() bool {
	return e.Js().Get("altKey").Bool()
}

type MouseButton int

const (
	// MouseButtonMain Main button pressed, usually the left button or the un-initialized state
	MouseButtonMain MouseButton = 0

	// MouseButtonAuxiliary Auxiliary button pressed, usually the wheel button or the middle button (if present)
	MouseButtonAuxiliary MouseButton = 1

	// MouseButtonSecondary Secondary button pressed, usually the right button
	MouseButtonSecondary MouseButton = 2

	// MouseButtonFourth Fourth button pressed, typically the Browser Back button
	MouseButtonFourth MouseButton = 3

	// MouseButtonFifth Fifth button pressed, typically the Browser Forward button
	MouseButtonFifth MouseButton = 4
)

// Button Returns an unsigned short representing the button that was pressed when the mouse event was fired.
func (e *MouseEvent) Button() MouseButton {
	return MouseButton(e.Js().Get("button").Int())
}

// Buttons Returns an unsigned long representing a bitmask of the buttons that are currently pressed.
func (e *MouseEvent) Buttons() int {
	return e.Js().Get("buttons").Int()
}

// ClientX Returns an integer representing the horizontal coordinate of the event within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (e *MouseEvent) ClientX() int {
	return e.Js().Get("clientX").Int()
}

// ClientY Returns an integer representing the vertical coordinate of the event within the application's client area at which the event occurred (as opposed to the coordinates within the page).
func (e *MouseEvent) ClientY() int {
	return e.Js().Get("clientY").Int()
}

// CtrlKey Returns a Boolean that is true if the Ctrl key was active when the event was generated.
func (e *MouseEvent) CtrlKey() bool {
	return e.Js().Get("ctrlKey").Bool()
}

// MetaKey Returns a Boolean that is true if the Meta key was active when the event was generated.
func (e *MouseEvent) MetaKey() bool {
	return e.Js().Get("metaKey").Bool()
}

// MovementX Returns an integer representing the horizontal component of the movement of the pointer between this event and the previous mousemove event.
func (e *MouseEvent) MovementX() int {
	return e.Js().Get("movementX").Int()
}

// MovementY Returns an integer representing the vertical component of the movement of the pointer between this event and the previous mousemove event.
func (e *MouseEvent) MovementY() int {
	return e.Js().Get("movementY").Int()
}

// OffsetX Returns an integer representing the horizontal coordinate of the event relative to the element the event fired on.
func (e *MouseEvent) OffsetX() int {
	return e.Js().Get("offsetX").Int()
}

// OffsetY Returns an integer representing the vertical coordinate of the event relative to the element the event fired on.
func (e *MouseEvent) OffsetY() int {
	return e.Js().Get("offsetY").Int()
}

// PageX Returns an integer representing the horizontal coordinate of the event relative to the whole document.
func (e *MouseEvent) PageX() int {
	return e.Js().Get("pageX").Int()
}

// PageY Returns an integer representing the vertical coordinate of the event relative to the whole document.
func (e *MouseEvent) PageY() int {
	return e.Js().Get("pageY").Int()
}

// RelatedTarget Returns the secondary target for the event, if there is one.
func (e *MouseEvent) RelatedTarget() *EventTarget {
	return NewEventTarget(e.Js().Get("relatedTarget"))
}

// ScreenX Returns an integer representing the horizontal coordinate of the event relative to the whole screen.
func (e *MouseEvent) ScreenX() int {
	return e.Js().Get("screenX").Int()
}

// ScreenY Returns an integer representing the vertical coordinate of the event relative to the whole screen.
func (e *MouseEvent) ScreenY() int {
	return e.Js().Get("screenY").Int()
}

// ShiftKey Returns a Boolean that is true if the Shift key was active when the event was generated.
func (e *MouseEvent) ShiftKey() bool {
	return e.Js().Get("shiftKey").Bool()
}

// X Returns an integer representing the horizontal coordinate of the event relative to the whole document.
func (e *MouseEvent) X() int {
	return e.Js().Get("x").Int()
}

// Y Returns an integer representing the vertical coordinate of the event relative to the whole document.
func (e *MouseEvent) Y() int {
	return e.Js().Get("y").Int()
}

// GetModifierState Returns a Boolean that is true if a modifier key is active when the event is generated.
func (e *MouseEvent) GetModifierState(keyArg string) bool {
	return e.Js().Call("getModifierState", keyArg).Bool()
}
