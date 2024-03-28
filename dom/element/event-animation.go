package element

// AnimationEvent https://developer.mozilla.org/en-US/docs/Web/API/AnimationEvent
type AnimationEvent struct {
	*Event
}

// AnimationName Returns the name of the animation.
func (e *AnimationEvent) AnimationName() string {
	return e.Js().Get("animationName").String()
}

// ElapsedTime Returns the amount of time the animation has been running, in seconds, when this event fired, excluding any time the animation was paused.
func (e *AnimationEvent) ElapsedTime() float64 {
	return e.Js().Get("elapsedTime").Float()
}

// PseudoElement Returns the name of the pseudo-element of the animation.
func (e *AnimationEvent) PseudoElement() string {
	return e.Js().Get("pseudoElement").String()
}
