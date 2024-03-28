package element

import "syscall/js"

// ErrorEvent https://developer.mozilla.org/en-US/docs/Web/API/ErrorEvent
type ErrorEvent struct {
	*Event
}

// Message Returns a DOMString representing a message from the server.
func (e *ErrorEvent) Message() string {
	return e.Js().Get("message").String()
}

// Filename Returns a DOMString representing the name of the script file in which the error occurred.
func (e *ErrorEvent) Filename() string {
	return e.Js().Get("filename").String()
}

// Lineno Returns an integer representing the line number in the script file on which the error occurred.
func (e *ErrorEvent) Lineno() int {
	return e.Js().Get("lineno").Int()
}

// Colno Returns an integer representing the column number in the script file on which the error occurred.
func (e *ErrorEvent) Colno() int {
	return e.Js().Get("colno").Int()
}

// Error Returns an Error object representing the error that occurred.
// TODO
func (e *ErrorEvent) Error() js.Value {
	return e.Js().Get("error")
}
