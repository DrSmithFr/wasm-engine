package element

import "syscall/js"

// KeyboardEvent https://developer.mozilla.org/en-US/docs/Web/API/KeyboardEvent
type KeyboardEvent struct {
	*Event
}

func NewKeyboardEvent(raw js.Value) *KeyboardEvent {
	return &KeyboardEvent{
		Event: NewEvent(raw),
	}
}

// AltKey Returns a Boolean that is true if the Alt (Option or ⌥ on OS X) key was active when the key event was generated.
func (e *KeyboardEvent) AltKey() bool {
	return e.Js().Get("altKey").Bool()
}

// Code Returns a DOMString representing the key value of the key represented by the event.
func (e *KeyboardEvent) Code() string {
	return e.Js().Get("code").String()
}

// CtrlKey Returns a Boolean that is true if the Ctrl key was active when the key event was generated.
func (e *KeyboardEvent) CtrlKey() bool {
	return e.Js().Get("ctrlKey").Bool()
}

// IsComposing Returns a Boolean that is true if the event is fired after compositionstart and before compositionend.
func (e *KeyboardEvent) IsComposing() bool {
	return e.Js().Get("isComposing").Bool()
}

// Key Returns a DOMString representing the key value of the key represented by the event.
func (e *KeyboardEvent) Key() string {
	return e.Js().Get("key").String()
}

type KeyboardLocation int

const (
	KeyboardLocationStandard KeyboardLocation = 0
	KeyboardLocationLeft     KeyboardLocation = 1
	KeyboardLocationRight    KeyboardLocation = 2
	KeyboardLocationNumpad   KeyboardLocation = 3
	KetboardLocationMobile   KeyboardLocation = 4
	KeyboardLocationJoystick KeyboardLocation = 5
)

// Location Returns an unsigned long representing the location of the key on the keyboard or other input device.
func (e *KeyboardEvent) Location() KeyboardLocation {
	return KeyboardLocation(e.Js().Get("location").Int())
}

// MetaKey Returns a Boolean that is true if the Meta key (on Mac keyboards, the ⌘ Command key; on Windows keyboards, the Windows key (⊞)) was active when the key event was generated.
func (e *KeyboardEvent) MetaKey() bool {
	return e.Js().Get("metaKey").Bool()
}

// Repeat Returns a Boolean that is true if the key is being held down such that it is automatically repeating.
func (e *KeyboardEvent) Repeat() bool {
	return e.Js().Get("repeat").Bool()
}

// ShiftKey Returns a Boolean that is true if the Shift key was active when the key event was generated.
func (e *KeyboardEvent) ShiftKey() bool {
	return e.Js().Get("shiftKey").Bool()
}
