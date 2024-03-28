package element

type PageTransitionEvent struct {
	*Event
}

// Persisted Returns a Boolean that is true if the page transition was done via a swipe gesture.
func (e *PageTransitionEvent) Persisted() bool {
	return e.Js().Get("persisted").Bool()
}
