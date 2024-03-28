package element

import "syscall/js"

// PointerEvent https://developer.mozilla.org/en-US/docs/Web/API/PointerEvent
type PointerEvent struct {
	*MouseEvent
}

func NewPointerEvent(entity js.Value) *PointerEvent {
	return &PointerEvent{
		MouseEvent: NewMouseEvent(entity),
	}
}

// PointerId Returns an identifier for the pointer causing the event.
func (e *PointerEvent) PointerId() int {
	return e.Js().Get("pointerId").Int()
}

// Width Returns the width of the pointer's contact geometry.
func (e *PointerEvent) Width() float64 {
	return e.Js().Get("width").Float()
}

// Height Returns the height of the pointer's contact geometry.
func (e *PointerEvent) Height() float64 {
	return e.Js().Get("height").Float()
}

// Pressure Returns the normalized pressure of the pointer input in the range 0 to 1, where 0 and 1 represent the minimum and maximum pressure the hardware is capable of detecting, respectively.
func (e *PointerEvent) Pressure() float64 {
	return e.Js().Get("pressure").Float()
}

// TangentialPressure The normalized tangential pressure of the pointer input (also known as barrel pressure or cylinder stress) in the range -1 to 1, where 0 is the neutral position of the control.
func (e *PointerEvent) TangentialPressure() float64 {
	return e.Js().Get("tangentialPressure").Float()
}

// TiltX Returns the plane angle between the Y-Z plane and the plane containing both the transducer (e.g. pen stylus) axis and the Y axis, in degrees, with a range of -90 to 90 degrees.
func (e *PointerEvent) TiltX() int {
	return e.Js().Get("tiltX").Int()
}

// TiltY Returns the plane angle between the X-Z plane and the plane containing both the transducer (e.g. pen stylus) axis and the X axis, in degrees, with a range of -90 to 90 degrees.
func (e *PointerEvent) TiltY() int {
	return e.Js().Get("tiltY").Int()
}

// Twist Returns the clockwise rotation of the transducer (e.g. pen stylus) around its own major axis, in degrees in the range 0 to 359.
func (e *PointerEvent) Twist() int {
	return e.Js().Get("twist").Int()
}

type PointerType string

const (
	PointerTypeMouse PointerType = "mouse"
	PointerTypePen   PointerType = "pen"
	PointerTypeTouch PointerType = "touch"
)

// PointerType Returns the type of pointer that triggered the event.
func (e *PointerEvent) PointerType() PointerType {
	return PointerType(e.Js().Get("pointerType").String())
}

// IsPrimary Returns true if the pointer represents the primary pointer in a multi-pointer environment.
func (e *PointerEvent) IsPrimary() bool {
	return e.Js().Get("isPrimary").Bool()
}

// GetCoalescedEvents Returns a list of all the PointerEvent objects representing all pointers associated with the pointer device that triggered the event.
func (e *PointerEvent) GetCoalescedEvents() []*PointerEvent {
	var events []*PointerEvent
	for i := 0; i < e.Js().Get("getCoalescedEvents").Length(); i++ {
		events = append(events, NewPointerEvent(e.Js().Get("getCoalescedEvents").Index(i)))
	}
	return events
}

// GetPredictedEvents Returns a list of all the PointerEvent objects representing all predicted pointer events that are expected to occur in the future.
func (e *PointerEvent) GetPredictedEvents() []*PointerEvent {
	var events []*PointerEvent
	for i := 0; i < e.Js().Get("getPredictedEvents").Length(); i++ {
		events = append(events, NewPointerEvent(e.Js().Get("getPredictedEvents").Index(i)))
	}
	return events
}
