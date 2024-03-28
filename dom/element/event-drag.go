package element

// DragEvent https://developer.mozilla.org/en-US/docs/Web/API/DragEvent
type DragEvent struct {
	*Event
}

// DataTransfer Returns a DataTransfer object representing the data being dragged during the drag operation.
func (e *DragEvent) DataTransfer() *DataTransfer {
	return NewDataTransfer(e.Js().Get("dataTransfer"))
}
