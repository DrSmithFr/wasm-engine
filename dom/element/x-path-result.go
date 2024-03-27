package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type XPathResultType int

const (
	XPathResultAnyType XPathResultType = iota
	XPathResultNumberType
	XPathResultStringType
	XPathResultBooleanType
	XPathResultUnorderedNodeIteratorType
	XPathResultOrderedNodeIteratorType
	XPathResultUnorderedNodeSnapshotType
	XPathResultOrderedNodeSnapshotType
	XPathResultAnyUnorderedNodeType
	XPathResultFirstOrderedNodeType
)

type XPathResult struct {
	*wasm.Entity
}

func NewXPathResult(raw js.Value) *XPathResult {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &XPathResult{
		Entity: wasm.New(raw),
	}
}

// BooleanValue Returns the result as a boolean value.
func (x *XPathResult) BooleanValue() bool {
	return x.Js().Get("booleanValue").Bool()
}

// InvalidIteratorState Returns true if the iterator has been invalidated.
func (x *XPathResult) InvalidIteratorState() bool {
	return x.Js().Get("invalidIteratorState").Bool()
}

// NumberValue Returns the result as a number value.
func (x *XPathResult) NumberValue() float64 {
	return x.Js().Get("numberValue").Float()
}

// ResultType Returns the result type.
func (x *XPathResult) ResultType() XPathResultType {
	return XPathResultType(x.Js().Get("resultType").Int())
}

// SingleNodeValue Returns the result as a single node value.
func (x *XPathResult) SingleNodeValue() *Node {
	return NewNode(x.Js().Get("singleNodeValue"))
}

// SnapshotLength Returns the number of nodes in the result snapshot.
func (x *XPathResult) SnapshotLength() int {
	return x.Js().Get("snapshotLength").Int()
}

// StringValue Returns the result as a string value.
func (x *XPathResult) StringValue() string {
	return x.Js().Get("stringValue").String()
}

// IterateNext Returns the next node in the result set.
func (x *XPathResult) IterateNext() *Node {
	return NewNode(x.Js().Call("iterateNext"))
}

// SnapshotItem Returns the node at the specified index in the result snapshot.
func (x *XPathResult) SnapshotItem(index int) *Node {
	return NewNode(x.Js().Call("snapshotItem", index))
}
