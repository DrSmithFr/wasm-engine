package ctx

import (
	"go-webgl/dom/helper"
	"go-webgl/dom/wasm"
	"image"
	"syscall/js"
)

type ImageData struct {
	data     js.Value
	copybuff js.Value
}

func NewImageData(data js.Value, width, height int) *ImageData {
	if data.IsNull() || data.IsUndefined() {
		return nil
	}

	return &ImageData{
		data:     data,
		copybuff: helper.NewUint8Array(width * height * 4),
	}
}

// enforce interface compliance
var _ wasm.Wasm = (*ImageData)(nil)

func (i *ImageData) Js() js.Value {
	return i.data
}

//
// ImageData methods
//

// SetData Sets the data for the ImageData object
func (i *ImageData) SetData(img *image.RGBA) {
	// TODO:  This currently does multiple data copies.   go image drawCtx -> JS Uint8Array,   Then JS Uint8Array -> ImageData,  then ImageData into the canvas.
	// Would like to eliminate at least one of them, however currently CopyBytesToJS only supports Uint8Array  rather than the Uint8ClampedArray of ImageData

	i.data.Get("data").Call("set", i.copybuff)
}
