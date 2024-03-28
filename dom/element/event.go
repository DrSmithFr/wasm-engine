package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Event https://developer.mozilla.org/en-US/docs/Web/API/Event
type Event struct {
	*wasm.Entity
}

func NewEvent(raw js.Value) *Event {
	return &Event{
		Entity: wasm.New(raw),
	}
}

// Bubbles Returns a Boolean that indicates whether or not an event is a bubbling event.
func (e *Event) Bubbles() bool {
	return e.Js().Get("bubbles").Bool()
}

// Cancelable Returns a Boolean that indicates whether the event can be canceled.
func (e *Event) Cancelable() bool {
	return e.Js().Get("cancelable").Bool()
}

// Composed Returns a Boolean that indicates whether the event will trigger listeners outside of a shadow root.
func (e *Event) Composed() bool {
	return e.Js().Get("composed").Bool()
}

// CurrentTarget Returns the object whose event listener's callback is currently being invoked.
func (e *Event) CurrentTarget() *EventTarget {
	return NewEventTarget(e.Js().Get("currentTarget"))
}

// DefaultPrevented Returns a Boolean that indicates whether the default action has been prevented.
func (e *Event) DefaultPrevented() bool {
	return e.Js().Get("defaultPrevented").Bool()
}

type EventPhase int

const (
	EventPhaseNone      EventPhase = 0
	EventPhaseCapturing EventPhase = 1
	EventPhaseAtTarget  EventPhase = 2
	EventPhaseBubbling  EventPhase = 3
)

// EventPhase Returns an unsigned short representing the current phase of event flow.
func (e *Event) EventPhase() EventPhase {
	return EventPhase(e.Js().Get("eventPhase").Int())
}

// IsTrusted Returns a Boolean that indicates whether or not the event is trusted.
func (e *Event) IsTrusted() bool {
	return e.Js().Get("isTrusted").Bool()
}

// Target Returns the object to which the event was originally dispatched.
func (e *Event) Target() *EventTarget {
	return NewEventTarget(e.Js().Get("target"))
}

// TimeStamp Returns a DOMHighResTimeStamp representing the time at which the event was created.
func (e *Event) TimeStamp() float64 {
	return e.Js().Get("timeStamp").Float()
}

// Type Returns a DOMString representing the type of event.
func (e *Event) Type() string {
	return e.Js().Get("type").String()
}

// ComposedPath Returns the event’s path (objects on which listeners will be invoked).
func (e *Event) ComposedPath() []*EventTarget {
	path := e.Js().Call("composedPath")

	if path.IsNull() || path.IsUndefined() {
		panic("composedPath() returned null or undefined")
	}

	targets := make([]*EventTarget, path.Length())
	for i := 0; i < path.Length(); i++ {
		targets[i] = NewEventTarget(path.Index(i))
	}

	return targets
}

// PreventDefault Cancels the event (if it is cancelable).
func (e *Event) PreventDefault() {
	e.Js().Call("preventDefault")
}

// StopImmediatePropagation Prevents other listeners of the same event from being called.
func (e *Event) StopImmediatePropagation() {
	e.Js().Call("stopImmediatePropagation")
}

// StopPropagation Prevents further propagation of the current event in the capturing and bubbling phases.
func (e *Event) StopPropagation() {
	e.Js().Call("stopPropagation")
}
