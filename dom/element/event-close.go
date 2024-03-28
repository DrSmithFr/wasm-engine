package element

// CloseEvent https://developer.mozilla.org/en-US/docs/Web/API/CloseEvent
type CloseEvent struct {
	*Event
}

// Code Returns an unsigned short representing the close code of the event.
func (e *CloseEvent) Code() int {
	return e.Js().Get("code").Int()
}

// Reason Returns a DOMString representing the reason the server closed the connection.
func (e *CloseEvent) Reason() string {
	return e.Js().Get("reason").String()
}

// WasClean Returns a Boolean that indicates whether or not the connection was cleanly closed.
func (e *CloseEvent) WasClean() bool {
	return e.Js().Get("wasClean").Bool()
}
