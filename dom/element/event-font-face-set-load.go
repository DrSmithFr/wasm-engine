package element

import "syscall/js"

// FontFaceSetLoadEvent https://developer.mozilla.org/en-US/docs/Web/API/FontFaceSetLoadEvent
type FontFaceSetLoadEvent struct {
	*Event
}

// Fontfaces Returns a FontFaceSet object representing the font faces that were loaded.
func (e *FontFaceSetLoadEvent) Fontfaces() js.Value {
	return e.Js().Get("fontfaces")
}
