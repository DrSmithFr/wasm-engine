package element

import (
	"syscall/js"
)

// MessagePort https://developer.mozilla.org/en-US/docs/Web/API/MessagePort
type MessagePort struct {
	*EventTarget
}

func NewMessagePort(entity js.Value) *MessagePort {
	return &MessagePort{
		EventTarget: NewEventTarget(entity),
	}
}

// PostMessage Sends a message from the port, and optionally, transfers ownership of objects to other browsing contexts.
func (m *MessagePort) PostMessage(message js.Value, transfer []js.Value) {
	m.Js().Call("postMessage", message, transfer)
}

// Start Starts the sending of messages queued on the port.
func (m *MessagePort) Start() {
	m.Js().Call("start")
}

// Close Disconnects the port, so that it is no longer active.
func (m *MessagePort) Close() {
	m.Js().Call("close")
}

// AddMessageEventListener Registers an event handler of a specific event type on the MessagePort.
func AddMessageEventListener(m *MessagePort, eventType string, callback func(e *MessageEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback(NewMessageEvent(args[0]))
		return nil
	})

	m.Js().Call("addEventListener", eventType, listener)
}

// AddMessageErrorEventListener Registers an event handler of a specific event type on the MessagePort.
func AddMessageErrorEventListener(m *MessagePort, eventType string, callback func(e *MessageEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		callback(NewMessageEvent(args[0]))
		return nil
	})

	m.Js().Call("addEventListener", eventType, listener)
}
