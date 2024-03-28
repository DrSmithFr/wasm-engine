package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// DataTransferList https://developer.mozilla.org/en-US/docs/Web/API/DataTransferList
type DataTransferList struct {
	*wasm.Entity
}

// NewDataTransferList returns a new DataTransferList object.
func NewDataTransferList(raw js.Value) *DataTransferList {
	return &DataTransferList{
		Entity: wasm.New(raw),
	}
}

// Length Returns the number of items in the drag data store.
func (d *DataTransferList) Length() int {
	return d.Js().Get("length").Int()
}

// Add Adds an item (either a string or a URL) to the drag data store.
// Todo

// Remove Removes all the elements from the drag data store.
func (d *DataTransferList) Remove(index int) {
	d.Js().Call("remove", index)
}

// Clear Removes all the elements from the drag data store.
func (d *DataTransferList) Clear() {
	d.Js().Call("clear")
}

// Get Returns the data for a given type.
func (d *DataTransferList) Get(index int) string {
	return d.Js().Call("getData", index).String()
}
