package element

import "syscall/js"

// TransitionEvent https://developer.mozilla.org/en-US/docs/Web/API/TransitionEvent
type TransitionEvent struct {
	*Event
}

func NewTransitionEvent(entity js.Value) *TransitionEvent {
	return &TransitionEvent{
		Event: NewEvent(entity),
	}
}

// PropertyName Returns the name of the CSS property associated with the transition.
func (e *TransitionEvent) PropertyName() string {
	return e.Js().Get("propertyName").String()
}

// ElapsedTime Returns the amount of time the transition has been running, in seconds, when this event fired.
func (e *TransitionEvent) ElapsedTime() float64 {
	return e.Js().Get("elapsedTime").Float()
}

// PseudoElement Returns the name of the pseudo-element of the element the animation runs on.
func (e *TransitionEvent) PseudoElement() string {
	return e.Js().Get("pseudoElement").String()
}
