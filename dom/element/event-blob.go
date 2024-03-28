package element

import "syscall/js"

// BlobEvent https://developer.mozilla.org/en-US/docs/Web/API/BlobEvent
type BlobEvent struct {
	*Event
}

// Data Returns a Blob object representing the data associated with the event.
func (e *BlobEvent) Data() js.Value {
	return e.Js().Get("data")
}

// Timecode Returns a DOMString representing a timestamp associated with the event.
func (e *BlobEvent) Timecode() int {
	return e.Js().Get("timecode").Int()
}
