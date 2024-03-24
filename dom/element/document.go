package element

import (
	"go-webgl/dom/element/list"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Document

type HTMLDocument Document

type Document struct {
	raw js.Value
}

// CurrentDocument enforce document Singleton
var CurrentDocument *Document

func LoadDocument() *Document {
	if CurrentDocument != nil {
		return CurrentDocument
	}

	CurrentDocument = &Document{raw: js.Global().Get("document")}

	return CurrentDocument
}

// enforce interface compliance
var _ wasm.WASM = (*Document)(nil)

func (d *Document) Bind(e js.Value) {
	d.raw = e
}

func (d *Document) Js() js.Value {
	return d.raw
}

//
// Document methods
//

// ActiveElement Returns the HTMLElement that currently has focus.
func (d *Document) ActiveElement() js.Value {
	return d.raw.Get("activeElement")
}

// AdoptedStyleSheets Add an array of constructed stylesheets to be used by the document.
// These stylesheets may also be shared with shadow DOM subtrees of the same document.
func (d *Document) AdoptedStyleSheets() []string {
	rawStyles := d.raw.Get("adoptedStyleSheets")
	styles := make([]string, rawStyles.Length())

	for i := 0; i < rawStyles.Length(); i++ {
		styles[i] = rawStyles.Index(i).String()
	}

	return styles
}

func (d *Document) SetAdoptedStyleSheets(styles []string) {
	d.raw.Set("adoptedStyleSheets", styles)
}

// Title Returns the title of the current document.
func (d *Document) Title() string {
	return d.raw.Get("title").String()
}

func (d *Document) SetTitle(title string) {
	d.raw.Set("title", title)
}

// URL Returns the URL of the current document.
func (d *Document) URL() string {
	return d.raw.Get("URL").String()
}

// Body Returns the <body> or <frameset> node of the current document, or null if no such element exists.
func (d *Document) Body() *HTMLElement {
	rawBody := d.raw.Get("body")
	return NewHTMLElement(rawBody)
}

// Head Returns the <head> element of the current document.
func (d *Document) Head() *HTMLElement {
	rawHead := d.raw.Get("head")
	return NewHTMLElement(rawHead)
}

//
// Document HTMLElement methods
//

// CreateElement Creates the HTML element specified by tagName.
func (d *Document) CreateElement(tagName list.Type) *HTMLElement {
	rawElement := d.raw.Call("createElement", string(tagName))
	return NewHTMLElement(rawElement)
}

func (d *Document) CreateCanvasElement() *CanvasElement {
	c := d.CreateElement("canvas")
	return NewCanvasElement(c)
}

// GetElementById Returns a reference to the element by its ID.
func (d *Document) GetElementById(id string) *HTMLElement {
	rawElement := d.raw.Call("getElementById", id)
	return NewHTMLElement(rawElement)
}

// GetElementsByClassName Returns an array-like object of all child elements which have all of the given class names.
func (d *Document) GetElementsByClassName(classNames string) []*HTMLElement {
	rawElements := d.raw.Call("getElementsByClassName", classNames)
	return NewHTMLElementList(rawElements)
}

// GetElementsByTagName Returns an HTMLCollection of elements with the given tag name.
func (d *Document) GetElementsByTagName(tagName string) []*HTMLElement {
	rawElements := d.raw.Call("getElementsByTagName", tagName)
	return NewHTMLElementList(rawElements)
}

// GetElementsByTagNameNS Returns an HTMLCollection of elements with the given tag name in the namespace.
func (d *Document) GetElementsByTagNameNS(namespace, tagName string) []*HTMLElement {
	rawElements := d.raw.Call("getElementsByTagNameNS", namespace, tagName)
	return NewHTMLElementList(rawElements)
}
