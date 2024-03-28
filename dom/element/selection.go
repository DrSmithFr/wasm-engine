package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Selection https://developer.mozilla.org/en-US/docs/Web/API/Selection
type Selection struct {
	*wasm.Entity
}

func NewSelection(raw js.Value) *Selection {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Selection{
		Entity: wasm.New(raw),
	}
}

// AnchorNode Returns the node in which the selection begins.
func (s *Selection) AnchorNode() *Node {
	return NewNode(s.Js().Get("anchorNode"))
}

// AnchorOffset Returns the number of characters that the selection's anchor is offset within the anchorNode.
func (s *Selection) AnchorOffset() int {
	return s.Js().Get("anchorOffset").Int()
}

// FocusNode Returns the node in which the selection ends.
func (s *Selection) FocusNode() *Node {
	return NewNode(s.Js().Get("focusNode"))
}

// FocusOffset Returns the number of characters that the selection's focus is offset within the focusNode.
func (s *Selection) FocusOffset() int {
	return s.Js().Get("focusOffset").Int()
}

// IsCollapsed Returns a Boolean indicating whether the selection's start and end points are at the same position.
func (s *Selection) IsCollapsed() bool {
	return s.Js().Get("isCollapsed").Bool()
}

// RangeCount Returns the number of ranges in the selection.
func (s *Selection) RangeCount() int {
	return s.Js().Get("rangeCount").Int()
}

type SelectionType string

const (
	SelectionTypeNone  SelectionType = "none"
	SelectionTypeCaret SelectionType = "caret"
	SelectionTypeRange SelectionType = "range"
)

// Type Returns the type of the selection.
func (s *Selection) Type() SelectionType {
	return SelectionType(s.Js().Get("type").String())
}

// AddRange Adds a range to the selection.
func (s *Selection) AddRange(r *Range) {
	s.Js().Call("addRange", r.Js())
}

// Collapse Collapses the selection to the end.
func (s *Selection) Collapse(node *Node) {
	s.Js().Call("collapse", node.Js())
}

func (s *Selection) CollapseWithOffset(node *Node, offset int) {
	s.Js().Call("collapse", node.Js(), offset)
}

// CollapseToEnd Collapses the selection to the end.
func (s *Selection) CollapseToEnd() {
	s.Js().Call("collapseToEnd")
}

// CollapseToStart Collapses the selection to the start.
func (s *Selection) CollapseToStart() {
	s.Js().Call("collapseToStart")
}

// ContainsNode Returns a Boolean indicating whether the selection contains a specified node.
func (s *Selection) ContainsNode(node *Node, allowPartial bool) bool {
	return s.Js().Call("containsNode", node.Js(), allowPartial).Bool()
}

// DeleteFromDocument Removes the selection from the document.
func (s *Selection) DeleteFromDocument() {
	s.Js().Call("deleteFromDocument")
}

// Extend Extends the selection.
func (s *Selection) Extend(node *Node) {
	s.Js().Call("extend", node.Js())
}

func (s *Selection) ExtendWithOffset(node *Node, offset int) {
	s.Js().Call("extend", node.Js(), offset)
}

// GetRangeAt Returns a range at a specified index.
func (s *Selection) GetRangeAt(index int) *Range {
	return NewRange(s.Js().Call("getRangeAt", index))
}

type SelectionAlter string

const (
	SelectionAlterMove   SelectionAlter = "move"
	SelectionAlterExtend SelectionAlter = "extend"
)

type SelectionDirection string

const (
	SelectionDirectionForward  SelectionDirection = "forward"
	SelectionDirectionBackward SelectionDirection = "backward"
	SelectionDirectionLeft     SelectionDirection = "left"
	SelectionDirectionRight    SelectionDirection = "right"
)

type SelectionGranularity string

const (
	SelectionGranularityCharacter         SelectionGranularity = "character"
	SelectionGranularityWord              SelectionGranularity = "word"
	SelectionGranularitySentence          SelectionGranularity = "sentence"
	SelectionGranularityLine              SelectionGranularity = "line"
	SelectionGranularityParagraph         SelectionGranularity = "paragraph"
	SelectionGranularityLineboundary      SelectionGranularity = "lineboundary"
	SelectionGranularitySentenceboundary  SelectionGranularity = "sentenceboundary"
	SelectionGranularityParagraphboundary SelectionGranularity = "paragraphboundary"
	SelectionGranularityDocumentBoundary  SelectionGranularity = "documentboundary"
)

// Modify Modifies the selection.
func (s *Selection) Modify(alter SelectionAlter, direction SelectionDirection, granularity SelectionGranularity) {
	s.Js().Call("modify", string(alter), string(direction), string(granularity))
}

// RemoveRange Removes a range from the selection.
func (s *Selection) RemoveRange(r *Range) {
	s.Js().Call("removeRange", r.Js())
}

// RemoveAllRanges Removes all ranges from the selection.
func (s *Selection) RemoveAllRanges() {
	s.Js().Call("removeAllRanges")
}

// SelectAllChildren Selects all the children of the specified node.
func (s *Selection) SelectAllChildren(node *Node) {
	s.Js().Call("selectAllChildren", node.Js())
}

// SetBaseAndExtent Sets the selection to be a range including all or parts of two specified DOM elements.
func (s *Selection) SetBaseAndExtent(baseNode *Node, baseOffset int, extentNode *Node, extentOffset int) {
	s.Js().Call("setBaseAndExtent", baseNode.Js(), baseOffset, extentNode.Js(), extentOffset)
}

// ToString Returns the text of the selection.
func (s *Selection) ToString() string {
	return s.Js().Call("toString").String()
}
