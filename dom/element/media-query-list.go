package element

import (
	"syscall/js"
)

type MediaQueryList struct {
	*EventTarget
}

func NewMediaQueryList(raw js.Value) *MediaQueryList {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &MediaQueryList{
		EventTarget: NewEventTarget(raw),
	}
}

// Matches Returns a Boolean that is true if the media query string matches the current environment.
func (m *MediaQueryList) Matches() bool {
	return m.Js().Get("matches").Bool()
}

// Media Returns the serialized media query list.
func (m *MediaQueryList) Media() string {
	return m.Js().Get("media").String()
}
