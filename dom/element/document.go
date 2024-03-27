package element

import (
	"go-webgl/dom/element/list"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Document

type HTMLDocument Document

type Document struct {
	*wasm.Entity
}

// CurrentDocument enforce document Singleton
var CurrentDocument *Document

func LoadDocument() *Document {
	if CurrentDocument != nil {
		return CurrentDocument
	}

	CurrentDocument = &Document{
		Entity: wasm.New(js.Global().Get("document")),
	}

	return CurrentDocument
}

func NewDocument(raw js.Value) *Document {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Document{
		Entity: wasm.New(raw),
	}
}

//
// Document methods
//

// ActiveElement Returns the HTMLElement that currently has focus.
func (d *Document) ActiveElement() js.Value {
	return d.Js().Get("activeElement")
}

// AdoptedStyleSheets Add an array of constructed stylesheets to be used by the document.
// These stylesheets may also be shared with shadow DOM subtrees of the same document.
func (d *Document) AdoptedStyleSheets() []string {
	rawStyles := d.Js().Get("adoptedStyleSheets")
	styles := make([]string, rawStyles.Length())

	for i := 0; i < rawStyles.Length(); i++ {
		styles[i] = rawStyles.Index(i).String()
	}

	return styles
}

func (d *Document) SetAdoptedStyleSheets(styles []string) {
	d.Js().Set("adoptedStyleSheets", styles)
}

// Title Returns the title of the current document.
func (d *Document) Title() string {
	return d.Js().Get("title").String()
}

func (d *Document) SetTitle(title string) {
	d.Js().Set("title", title)
}

// URL Returns the URL of the current document.
func (d *Document) URL() string {
	return d.Js().Get("URL").String()
}

// Body Returns the <body> or <frameset> node of the current document, or null if no such element exists.
func (d *Document) Body() *HTMLElement {
	rawBody := d.Js().Get("body")
	return NewHTMLElement(rawBody)
}

// Head Returns the <head> element of the current document.
func (d *Document) Head() *HTMLElement {
	rawHead := d.Js().Get("head")
	return NewHTMLElement(rawHead)
}

//
// Document HTMLElement methods
//

// CreateElement Creates the HTML element specified by tagName.
func (d *Document) CreateElement(tagName list.Type) *HTMLElement {
	rawElement := d.Js().Call("createElement", string(tagName))
	return NewHTMLElement(rawElement)
}

func (d *Document) CreateCanvasElement() *CanvasElement {
	c := d.CreateElement("canvas")
	return NewCanvasElement(c)
}

// GetElementById Returns a reference to the element by its ID.
func (d *Document) GetElementById(id string) *HTMLElement {
	rawElement := d.Js().Call("getElementById", id)
	return NewHTMLElement(rawElement)
}

// GetElementsByClassName Returns an array-like object of all child elements which have all of the given class names.
func (d *Document) GetElementsByClassName(classNames string) []*HTMLElement {
	rawElements := d.Js().Call("getElementsByClassName", classNames)
	return NewHTMLElementList(rawElements)
}

// GetElementsByTagName Returns an HTMLCollection of elements with the given tag name.
func (d *Document) GetElementsByTagName(tagName string) []*HTMLElement {
	rawElements := d.Js().Call("getElementsByTagName", tagName)
	return NewHTMLElementList(rawElements)
}

// GetElementsByTagNameNS Returns an HTMLCollection of elements with the given tag name in the namespace.
func (d *Document) GetElementsByTagNameNS(namespace, tagName string) []*HTMLElement {
	rawElements := d.Js().Call("getElementsByTagNameNS", namespace, tagName)
	return NewHTMLElementList(rawElements)
}
