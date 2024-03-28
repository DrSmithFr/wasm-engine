package element

// FocusEvent https://developer.mozilla.org/en-US/docs/Web/API/FocusEvent
type FocusEvent struct {
	*Event
}

// RelatedTarget Returns the secondary target for the event, if there is one.
func (e *FocusEvent) RelatedTarget() *EventTarget {
	return NewEventTarget(e.Js().Get("relatedTarget"))
}
