package element

// BeforeUnloadEvent https://developer.mozilla.org/en-US/docs/Web/API/BeforeUnloadEvent
type BeforeUnloadEvent struct {
	*Event
}

// ReturnValue Returns a string that the event handler may set to display a message to the user in a dialog box launched by the user agent to confirm that the page should be unloaded.
func (e *BeforeUnloadEvent) ReturnValue() string {
	return e.Js().Get("returnValue").String()
}
