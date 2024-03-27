package element

import "syscall/js"

type CharacterData struct {
	*Node
}

func NewCharacterData(raw js.Value) *CharacterData {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &CharacterData{
		Node: NewNode(raw),
	}
}

// Data Returns the character data of the node.
func (c *CharacterData) Data() string {
	return c.Js().Get("data").String()
}

// SetData Sets the character data of the node.
func (c *CharacterData) SetData(data string) {
	c.Js().Set("data", data)
}

// Length Returns the length of the character data.
func (c *CharacterData) Length() int {
	return c.Js().Get("length").Int()
}

// NextSibling Returns the next sibling node.
func (c *CharacterData) NextSibling() *Element {
	return NewElement(c.Js().Get("nextSibling"))
}

// PreviousSibling Returns the previous sibling node.
func (c *CharacterData) PreviousSibling() *Element {
	return NewElement(c.Js().Get("previousSibling"))
}

// After Inserts a set of Node or DOMString objects in the children list of this CharacterData's parent, just after this CharacterData.
func (c *CharacterData) After(nodes ...*Node) {
	elems := make([]interface{}, len(nodes))
	for i, node := range nodes {
		elems[i] = node.Js()
	}

	c.Js().Call("after", elems...)
}

// AppendData Appends the given DOMString to the CharacterData.data string; when this method returns, data contains the concatenated DOMString.
func (c *CharacterData) AppendData(data string) {
	c.Js().Call("appendData", data)
}

// Before Inserts a set of Node or DOMString objects in the children list of this CharacterData's parent, just before this CharacterData.
func (c *CharacterData) Before(nodes ...*Node) {
	elems := make([]interface{}, len(nodes))
	for i, node := range nodes {
		elems[i] = node.Js()
	}

	c.Js().Call("before", elems...)
}

// DeleteData Removes the specified amount of characters, starting at the specified offset, from the CharacterData.data string; when this method returns, data contains the shortened DOMString.
func (c *CharacterData) DeleteData(offset, count int) {
	c.Js().Call("deleteData", offset, count)
}

// InsertData Inserts the specified DOMString at the specified offset; when this method returns, data contains the inserted DOMString.
func (c *CharacterData) InsertData(offset int, data string) {
	c.Js().Call("insertData", offset, data)
}

// Remove Removes the object from its parent children list.
func (c *CharacterData) Remove() {
	c.Js().Call("remove")
}

// ReplaceData Replaces the specified amount of characters, starting at the specified offset, with the specified DOMString; when this method returns, data contains the modified DOMString.
func (c *CharacterData) ReplaceData(offset, count int, data string) {
	c.Js().Call("replaceData", offset, count, data)
}

// ReplaceWith Replaces this CharacterData in the children list of its parent with a set of Node or DOMString objects.
func (c *CharacterData) ReplaceWith(nodes ...*Node) {
	elems := make([]interface{}, len(nodes))
	for i, node := range nodes {
		elems[i] = node.Js()
	}

	c.Js().Call("replaceWith", elems...)
}

// SubstringData Returns a DOMString containing the part of CharacterData.data of the specified length and starting at the specified offset.
func (c *CharacterData) SubstringData(offset, count int) string {
	return c.Js().Call("substringData", offset, count).String()
}
