package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Screen https://developer.mozilla.org/en-US/docs/Web/API/Screen
type Screen struct {
	*wasm.Entity
}

func NewScreen(raw js.Value) *Screen {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Screen{
		Entity: wasm.New(raw),
	}
}

// AvailHeight Returns the height of the available screen space.
func (s *Screen) AvailHeight() int {
	return s.Js().Get("availHeight").Int()
}

// AvailWidth Returns the width of the available screen space.
func (s *Screen) AvailWidth() int {
	return s.Js().Get("availWidth").Int()
}

// ColorDepth Returns the bit depth of the color palette for displaying images.
func (s *Screen) ColorDepth() int {
	return s.Js().Get("colorDepth").Int()
}

// Height Returns the height of the screen.
func (s *Screen) Height() int {
	return s.Js().Get("height").Int()
}

// IsExtended Returns true if the screen is in extended mode.
func (s *Screen) IsExtended() bool {
	return s.Js().Get("isExtended").Bool()
}

// Orientation Returns the current orientation of the screen.
func (s *Screen) Orientation() string {
	return s.Js().Get("orientation").String()
}

// PixelDepth Returns the bit depth of the screen.
func (s *Screen) PixelDepth() int {
	return s.Js().Get("pixelDepth").Int()
}
