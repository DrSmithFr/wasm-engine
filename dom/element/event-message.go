package element

import "syscall/js"

// MessageEvent https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent
type MessageEvent struct {
	*Event
}

func NewMessageEvent(raw js.Value) *MessageEvent {
	return &MessageEvent{
		Event: NewEvent(raw),
	}
}

// Data The data sent by the message emitter
func (e *MessageEvent) Data() string {
	return e.Js().Get("data").String()
}

// Origin A string representing the origin of the message emitter.
func (e *MessageEvent) Origin() string {
	return e.Js().Get("origin").String()
}

// LastEventId A string representing a unique ID for the event.
func (e *MessageEvent) LastEventId() string {
	return e.Js().Get("lastEventId").String()
}

// Source A MessageEventSource (which can be a WindowProxy, MessagePort, or ServiceWorker object) representing the message emitter.
func (e *MessageEvent) Source() js.Value {
	return e.Js().Get("source")
}

// Ports An array of MessagePort objects representing the ports associated with the channel the message is being sent through (where appropriate, e.g. in channel messaging or when sending a message to a shared worker).
func (e *MessageEvent) Ports() js.Value {
	return e.Js().Get("ports")
}
