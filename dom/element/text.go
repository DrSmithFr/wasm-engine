package element

import "syscall/js"

type Text struct {
	*CharacterData
}

func NewText(raw js.Value) *Text {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Text{
		CharacterData: NewCharacterData(raw),
	}
}

// AssignedSlot Returns the slot the text node is assigned to.
func (t *Text) AssignedSlot() *HTMLSlotElement {
	return NewHTMLSlotElement(t.Js().Get("assignedSlot"))
}

// WholeText Returns all text of text nodes in a given range.
func (t *Text) WholeText() string {
	return t.Js().Get("wholeText").String()
}

// SplitText Splits a text node into two text nodes at the specified offset.
func (t *Text) SplitText(offset int) *Text {
	newNode := t.Js().Call("splitText", offset)
	return NewText(newNode)
}
