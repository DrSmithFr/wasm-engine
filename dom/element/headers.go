package element

import (
	"syscall/js"
)

// Headers https://developer.mozilla.org/en-US/docs/Web/API/Headers
type Headers struct {
	*HeadersReadOnly
}

// NewHeaders returns a new Headers object.
func NewHeaders(headers js.Value) *Headers {
	return &Headers{
		HeadersReadOnly: NewHeadersReadOnly(headers),
	}
}

// Append Appends a new value onto an existing header inside a Headers object, or adds the header if it does not already exist.
func (h *Headers) Append(name string, value string) {
	h.Js().Call("append", name, value)
}

// Delete Deletes a header from a Headers object.
func (h *Headers) Delete(name string) {
	h.Js().Call("delete", name)
}

// Set Sets a new value for an existing header inside a Headers object, or adds the header if it does not already exist.
func (h *Headers) Set(name string, value string) {
	h.Js().Call("set", name, value)
}
