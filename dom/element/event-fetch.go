package element

import (
	"go-webgl/dom/promise"
	"syscall/js"
)

// FetchEvent https://developer.mozilla.org/en-US/docs/Web/API/FetchEvent
type FetchEvent struct {
	*Event
}

// ClientId Returns a string representing the id of the client that the current service worker is controlling (one of the clients within the service worker's scope.)
func (e *FetchEvent) ClientId() string {
	return e.Js().Get("clientId").String()
}

// Handled Returns A promise that is pending while the event has not been handled, and fulfilled once it has.
func (e *FetchEvent) Handled() *promise.Promise {
	return promise.New(e.Js().Get("handled"))
}

// PreloadResponse Returns a promise that resolves with a Response object.
func (e *FetchEvent) PreloadResponse() *promise.Promise {
	return promise.New(e.Js().Get("preloadResponse"))
}

// ReplaceClientIds Replaces the client IDs of the current service worker with a new client ID.
func (e *FetchEvent) ReplaceClientIds(oldClientId string, newClientId string) {
	e.Js().Call("replaceClientIds", oldClientId, newClientId)
}

// ResultingClientId Returns a string representing the id of the client that the current service worker is controlling (one of the clients within the service worker's scope.)
func (e *FetchEvent) ResultingClientId() string {
	return e.Js().Get("resultingClientId").String()
}

// Request Returns a Request object representing the request.
func (e *FetchEvent) Request() *Request {
	return NewRequest(e.Js().Get("request"))
}

// RespondWith Prevents the browser's default fetch handling, and allows you to provide a response yourself.
func (e *FetchEvent) RespondWith(response *Response) {
	e.Js().Call("respondWith", js.ValueOf(&response))
}

// WaitUntil Extends the lifetime of the event. It is intended to be called in the install event handler.
func (e *FetchEvent) WaitUntil(promise *promise.Promise) {
	e.Js().Call("waitUntil", promise.Js())
}
