package element

import "syscall/js"

type Attr struct {
	*Node
}

func NewAttr(raw js.Value) *Attr {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Attr{
		Node: NewNode(raw),
	}
}

// LocalName Returns the local name of the attribute.
func (a *Attr) LocalName() string {
	return a.Js().Get("localName").String()
}

// Name Returns the name of the attribute.
func (a *Attr) Name() string {
	return a.Js().Get("name").String()
}

// NamespaceURI Returns the namespace URI of the attribute.
func (a *Attr) NamespaceURI() string {
	return a.Js().Get("namespaceURI").String()
}

// OwnerElement Returns the element node the attribute is attached to.
func (a *Attr) OwnerElement() *HTMLElement {
	return NewHTMLElement(a.Js().Get("ownerElement"))
}

// Prefix Returns the namespace prefix of the attribute.
func (a *Attr) Prefix() string {
	return a.Js().Get("prefix").String()
}

// Specified Returns true if the attribute is specified.
func (a *Attr) Specified() bool {
	return a.Js().Get("specified").Bool()
}

// Value Returns the value of the attribute.
func (a *Attr) Value() string {
	return a.Js().Get("value").String()
}

// SetValue Sets the value of the attribute.
func (a *Attr) SetValue(value string) {
	a.Js().Set("value", value)
}
