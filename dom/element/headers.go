package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type HeadersReadOnly struct {
	*wasm.Entity
}

// NewHeadersReadOnly returns a new Headers object.
func NewHeadersReadOnly(headers js.Value) *HeadersReadOnly {
	return &HeadersReadOnly{
		Entity: wasm.New(headers),
	}
}

// Entries Returns an iterator allowing to go through all key/value pairs contained in this object.
func (h *HeadersReadOnly) Entries() *Iterator {
	return NewIterator(h.Js().Call("entries"))
}

// forEach Allows to perform a given action for each key/value pair in this object.
func (h *HeadersReadOnly) ForEach(callback js.Func) {
	h.Js().Call("forEach", callback)
}

// Get Returns the first value of the specified header from a Headers object.
func (h *HeadersReadOnly) Get(name string) string {
	return h.Js().Call("get", name).String()
}

// GetSetCookie Returns the first value of the specified header from a Headers object.
func (h *HeadersReadOnly) GetSetCookie() string {
	return h.Js().Call("get", "Set-Cookie").String()
}

// Has Returns a boolean stating if a Headers object contains a certain header.
func (h *HeadersReadOnly) Has(name string) bool {
	return h.Js().Call("has", name).Bool()
}

// Keys Returns an iterator allowing to go through all keys of the key/value pairs contained in this object.
func (h *HeadersReadOnly) Keys() *Iterator {
	return NewIterator(h.Js().Call("keys"))
}

// Values Returns an iterator allowing to go through all values of the key/value pairs contained in this object.
func (h *HeadersReadOnly) Values() *Iterator {
	return NewIterator(h.Js().Call("values"))
}

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
