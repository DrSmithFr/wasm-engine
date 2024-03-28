package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// DataTransfer https://developer.mozilla.org/en-US/docs/Web/API/DataTransfer
type DataTransfer struct {
	*wasm.Entity
}

// NewDataTransfer returns a new DataTransfer object.
func NewDataTransfer(raw js.Value) *DataTransfer {
	return &DataTransfer{
		Entity: wasm.New(raw),
	}
}

type DataTransferDropEffect string

const (
	DataTransferDropEffectNone DataTransferDropEffect = "none"
	DataTransferDropEffectCopy DataTransferDropEffect = "copy"
	DataTransferDropEffectLink DataTransferDropEffect = "link"
	DataTransferDropEffectMove DataTransferDropEffect = "move"
)

// DropEffect Returns the effect that will be used on a drop.
func (d *DataTransfer) DropEffect() DataTransferDropEffect {
	return DataTransferDropEffect(d.Js().Get("dropEffect").String())
}

// SetDropEffect Sets the effect that will be used on a drop.
func (d *DataTransfer) SetDropEffect(effect DataTransferDropEffect) {
	d.Js().Set("dropEffect", string(effect))
}

type DataTransferEffectAllowed string

const (
	DataTransferEffectAllowedNone          DataTransferEffectAllowed = "none"
	DataTransferEffectAllowedCopy          DataTransferEffectAllowed = "copy"
	DataTransferEffectAllowedCopyLink      DataTransferEffectAllowed = "copyLink"
	DataTransferEffectAllowedCopyMove      DataTransferEffectAllowed = "copyMove"
	DataTransferEffectAllowedLink          DataTransferEffectAllowed = "link"
	DataTransferEffectAllowedLinkMove      DataTransferEffectAllowed = "linkMove"
	DataTransferEffectAllowedMove          DataTransferEffectAllowed = "move"
	DataTransferEffectAllowedAll           DataTransferEffectAllowed = "all"
	DataTransferEffectAllowedUninitialized DataTransferEffectAllowed = "uninitialized"
)

// EffectAllowed Returns the types of operations that are possible.
func (d *DataTransfer) EffectAllowed() DataTransferEffectAllowed {
	return DataTransferEffectAllowed(d.Js().Get("effectAllowed").String())
}

// SetEffectAllowed Sets the types of operations that are possible.
func (d *DataTransfer) SetEffectAllowed(effect DataTransferEffectAllowed) {
	d.Js().Set("effectAllowed", string(effect))
}

// Files Returns a FileList object.
func (d *DataTransfer) Files() js.Value {
	return d.Js().Get("files")
}

// Items Returns a DataTransferItemList object.
func (d *DataTransfer) Items() *DataTransferList {
	return NewDataTransferList(d.Js().Get("items"))
}

// Types Returns list of strings giving the formats that were set in the dragstart event.
func (d *DataTransfer) Types() []string {
	types := d.Js().Get("types")

	if types.IsNull() || types.IsUndefined() {
		panic("types is null or undefined")
	}

	var result []string
	for i := 0; i < types.Length(); i++ {
		result = append(result, types.Index(i).String())
	}

	return result
}
