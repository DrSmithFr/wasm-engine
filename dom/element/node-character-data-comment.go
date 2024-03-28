package element

import "syscall/js"

// Comment https://developer.mozilla.org/en-US/docs/Web/API/Comment
type Comment struct {
	*CharacterData
}

func NewComment(raw js.Value) *Comment {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Comment{
		CharacterData: NewCharacterData(raw),
	}
}
