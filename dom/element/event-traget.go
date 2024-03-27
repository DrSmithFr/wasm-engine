package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type EventTargetI interface {
	// AddEventListener Registers an event handler of a specific event type on the EventTarget.
	// The event target may be an HTMLElement in a document, the Document itself, a Window, or any other object that supports events
	// (such as XMLHttpRequest).
	// If the event type is one of the DOMActivate, DOMFocusIn, DOMFocusOut, DOMNodeRemoved, DOMSubtreeModified, DOMAttrModified,
	// DOMCharacterDataModified, DOMAttributeNameChanged, DOMElementNameChanged, or DOMElementNameChanged events, this method has no effect.
	AddEventListener(eventType string, listener js.Func)

	// DispatchEvent Dispatches an Event at the specified EventTarget, invoking the affected EventListeners in the appropriate order.
	// The normal event processing rules (including the capturing and optional bubbling phase) also apply to events dispatched manually with dispatchEvent().
	// Note: This method is a proprietary Mozilla extension and is not on a standards track.
	DispatchEvent(event js.Value)

	// RemoveEventListener Removes an event listener from the EventTarget.
	// If an event listener is removed from an EventTarget while it is processing an event, it will not be triggered by the current actions.
	// Event listeners can never be invoked after being removed.
	RemoveEventListener(eventType string, listener js.Func)
}

type EventTarget struct {
	*wasm.Entity
}

func NewEventTarget(raw js.Value) *EventTarget {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &EventTarget{
		Entity: wasm.New(raw),
	}
}

func (e *EventTarget) AddEventListener(eventType string, listener js.Func) {
	e.Js().Call("addEventListener", eventType, listener)
}

func (e *EventTarget) DispatchEvent(event js.Value) {
	e.Js().Call("dispatchEvent", event)
}

func (e *EventTarget) RemoveEventListener(eventType string, listener js.Func) {
	e.Js().Call("removeEventListener", eventType, listener)
}
