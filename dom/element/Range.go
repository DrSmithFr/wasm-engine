package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type Range struct {
	*wasm.Entity
}

func NewRange(raw js.Value) *Range {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Range{
		Entity: wasm.New(raw),
	}
}

// Collapsed Returns true if the range is collapsed.
func (r *Range) Collapsed() bool {
	return r.Js().Get("collapsed").Bool()
}

// CommonAncestorContainer Returns the common ancestor container.
func (r *Range) CommonAncestorContainer() *Node {
	return NewNode(r.Js().Get("commonAncestorContainer"))
}

// EndContainer Returns the end container.
func (r *Range) EndContainer() *Node {
	return NewNode(r.Js().Get("endContainer"))
}

// EndOffset Returns the end offset.
func (r *Range) EndOffset() int {
	return r.Js().Get("endOffset").Int()
}

// StartContainer Returns the start container.
func (r *Range) StartContainer() *Node {
	return NewNode(r.Js().Get("startContainer"))
}

// StartOffset Returns the start offset.
func (r *Range) StartOffset() int {
	return r.Js().Get("startOffset").Int()
}

// Collapse Collapses the range.
func (r *Range) Collapse(toStart bool) {
	r.Js().Call("collapse", toStart)
}

type CompareBoundaryMode int

const (
	CompareBoundaryModeStartToStart CompareBoundaryMode = 0
	CompareBoundaryModeStartToEnd                       = 1
	CompareBoundaryModeEndToEnd                         = 2
	CompareBoundaryModeEndToStart                       = 3
)

type BoundaryType int

const (
	BoundaryBefore BoundaryType = -1
	BoundaryEqual               = 0
	BoundaryInside              = 0
	BoundaryAfter               = 1
)

// CompareBoundaryPoints Compares the boundary points.
func (r *Range) CompareBoundaryPoints(how CompareBoundaryMode, sourceRange *Range) BoundaryType {
	boundrary := r.Js().Call("compareBoundaryPoints", int(how), sourceRange.Js()).Int()
	return BoundaryType(boundrary)
}

// ComparePoint Compares the point.
func (r *Range) ComparePoint(node *Node, offset int) BoundaryType {
	boundrary := r.Js().Call("comparePoint", node.Js(), offset).Int()
	return BoundaryType(boundrary)
}

// CloneContents Clones the contents.
func (r *Range) CloneContents() *DocumentFragment {
	return NewDocumentFragment(r.Js().Call("cloneContents"))
}

// CloneRange Clones the range.
func (r *Range) CloneRange() *Range {
	return NewRange(r.Js().Call("cloneRange"))
}

// CreateContextualFragment Creates a contextual fragment.
func (r *Range) CreateContextualFragment(fragment string) *DocumentFragment {
	return NewDocumentFragment(r.Js().Call("createContextualFragment", fragment))
}

// DeleteContents Deletes the contents.
func (r *Range) DeleteContents() {
	r.Js().Call("deleteContents")
}

// Detach Does nothing. Kept for compatibility.
func (r *Range) Detach() {
	r.Js().Call("detach")
}

// ExtractContents Extracts the contents.
func (r *Range) ExtractContents() *DocumentFragment {
	return NewDocumentFragment(r.Js().Call("extractContents"))
}

// GetBoundingClientRect Returns the client rectangle.
func (r *Range) GetBoundingClientRect() *DOMRect {
	return NewDOMRect(r.Js().Call("getBoundingClientRect"))
}
