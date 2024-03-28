package element

import (
	"go-webgl/dom/wasm"
	"log"
	"syscall/js"
)

// CaretPosition https://developer.mozilla.org/en-US/docs/Web/API/CaretPosition
type CaretPosition struct {
	*wasm.Entity
}

func NewCaretPosition(raw js.Value) *CaretPosition {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &CaretPosition{
		Entity: wasm.New(raw),
	}
}

// OffsetNode Returns the node.
func (c *CaretPosition) OffsetNode() *Node {
	return NewNode(c.Js().Get("offsetNode"))
}

// Offset Returns the offset.
func (c *CaretPosition) Offset() int {
	return c.Js().Get("offset").Int()
}

// GetClientRect Returns the client rectangle.
func (c *CaretPosition) GetClientRect() js.Value {
	log.Println("CaretPosition.GetClientRect() is not implemented.")
	return c.Js().Call("getClientRect")
}
