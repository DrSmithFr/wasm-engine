package element

// ClipboardEvent https://developer.mozilla.org/en-US/docs/Web/API/ClipboardEvent
type ClipboardEvent struct {
	*Event
}

// ClipboardData Returns a DataTransfer object containing the data affected by the user-initiated cut, copy, or paste operation.
func (e *ClipboardEvent) ClipboardData() *DataTransfer {
	return NewDataTransfer(e.Js().Get("clipboardData"))
}
