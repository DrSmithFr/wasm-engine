package element

import "syscall/js"

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
