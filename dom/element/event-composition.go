package element

type CompositionEvent struct {
	*Event
}

// Data Returns a DOMString representing the data entered by the user.
func (e *CompositionEvent) Data() string {
	return e.Js().Get("data").String()
}

// Locale Returns a DOMString representing the locale of the keyboard.
func (e *CompositionEvent) Locale() string {
	return e.Js().Get("locale").String()
}
