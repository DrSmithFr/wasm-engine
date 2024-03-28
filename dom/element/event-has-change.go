package element

// HashChangeEvent https://developer.mozilla.org/en-US/docs/Web/API/HashChangeEvent
type HashChangeEvent struct {
	*Event
}

// NewURL Returns a DOMString representing the new URL to which the window is navigating.
func (e *HashChangeEvent) NewURL() string {
	return e.Js().Get("newURL").String()
}

// OldURL Returns a DOMString representing the old URL from which the window was navigated.
func (e *HashChangeEvent) OldURL() string {
	return e.Js().Get("oldURL").String()
}
