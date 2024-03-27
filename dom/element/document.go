package element

import (
	"go-webgl/dom/element/list"
	"go-webgl/dom/promise"
	"go-webgl/dom/wasm"
	"log"
	"syscall/js"
	"time"
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

// Body Returns the <body> or <frameset> node of the current document, or null if no such element exists.
func (d *Document) Body() *HTMLElement {
	rawBody := d.Js().Get("body")
	return NewHTMLElement(rawBody)
}

// CharacterSet Returns the character encoding of the current document.
func (d *Document) CharacterSet() string {
	return d.Js().Get("characterSet").String()
}

// ChildElementCount Returns the number of child elements of the current element.
func (d *Document) ChildElementCount() int {
	return d.Js().Get("childElementCount").Int()
}

// Children Returns a live HTMLCollection containing all of the children of this node.
func (d *Document) Children() *HTMLCollection {
	rawChildren := d.Js().Get("children")
	return NewHTMLCollection(rawChildren)
}

type DocumentCompatMode string

const (
	CompatModeBackCompat DocumentCompatMode = "BackCompat"
	CompatModeCSS1Compat DocumentCompatMode = "CSS1Compat"
)

// CompatMode Returns the mode used by the browser to render the document.
func (d *Document) CompatMode() DocumentCompatMode {
	mode := d.Js().Get("compatMode").String()
	return DocumentCompatMode(mode)
}

// ContentType Returns the Content-Type from the MIME Header of the current document.
func (d *Document) ContentType() string {
	return d.Js().Get("contentType").String()
}

// CurrentScript Returns the script element, that is currently executing, or null if the script is global or not currently executing.
func (d *Document) CurrentScript() js.Value {
	log.Println("CurrentScript is not implemented")
	return d.Js().Get("currentScript")
}

// DocType Returns the Document Type Declaration associated with the document.
func (d *Document) DocType() js.Value {
	log.Println("DocType is not implemented")
	return d.Js().Get("doctype")
}

// DocumentElement Returns the Document Element of the document (the <html> element).
func (d *Document) DocumentElement() *HTMLElement {
	rawElement := d.Js().Get("documentElement")
	return NewHTMLElement(rawElement)
}

// DocumentURI Returns the URI of the current document.
func (d *Document) DocumentURI() string {
	return d.Js().Get("documentURI").String()
}

// Embeds Returns a list of the embedded <embed> elements within the current document.
func (d *Document) Embeds() *HTMLCollection {
	embeds := d.Js().Get("embeds")
	return NewHTMLCollection(embeds)
}

// FeaturePolicy Returns the FeaturePolicy object associated with the document.
func (d *Document) FeaturePolicy() js.Value {
	log.Println("FeaturePolicy is not implemented")
	return d.Js().Get("featurePolicy")
}

// FirstElementChild Returns the first child element of the current element.
func (d *Document) FirstElementChild() *HTMLElement {
	rawElement := d.Js().Get("firstElementChild")
	return NewHTMLElement(rawElement)
}

// Fonts Returns a FontFaceSet object which is a set-like object that represents a list of FontFace objects.
func (d *Document) Fonts() js.Value {
	log.Println("Fonts is not implemented")
	return d.Js().Get("fonts")
}

// Forms Returns a live HTMLCollection containing all the forms of the current document.
func (d *Document) Forms() *HTMLCollection {
	rawForms := d.Js().Get("forms")
	return NewHTMLCollection(rawForms)
}

// FragmentDirective Returns true is supported
func (d *Document) FragmentDirective() bool {
	if d.Js().Get("fragmentDirective").IsUndefined() {
		return false
	}

	return true
}

// FullscreenElement Returns the element that is currently being presented in fullscreen mode in this document, or null if full-screen mode is not currently in use.
func (d *Document) FullscreenElement() *HTMLElement {
	rawElement := d.Js().Get("fullscreenElement")
	return NewHTMLElement(rawElement)
}

// Head Returns the <head> element of the current document.
func (d *Document) Head() *HTMLElement {
	rawHead := d.Js().Get("head")
	return NewHTMLElement(rawHead)
}

// Hidden Returns a Boolean value indicating whether the document is hidden or not.
func (d *Document) Hidden() bool {
	return d.Js().Get("hidden").Bool()
}

// Images Returns a live HTMLCollection containing all the images of the current document.
func (d *Document) Images() *HTMLCollection {
	rawImages := d.Js().Get("images")
	return NewHTMLCollection(rawImages)
}

// Implementation Returns the DOMImplementation object that handles this document.
func (d *Document) Implementation() js.Value {
	log.Println("Implementation is not implemented")
	return d.Js().Get("implementation")
}

// LastElementChild Returns the last child element of the current element.
func (d *Document) LastElementChild() *HTMLElement {
	rawElement := d.Js().Get("lastElementChild")
	return NewHTMLElement(rawElement)
}

// Links Returns a live HTMLCollection containing all the hyperlinks of the current document.
func (d *Document) Links() *HTMLCollection {
	rawLinks := d.Js().Get("links")
	return NewHTMLCollection(rawLinks)
}

// PictureInPictureElement Returns the Picture-in-Picture (PiP) window for the document.
func (d *Document) PictureInPictureElement() *HTMLElement {
	rawElement := d.Js().Get("pictureInPictureElement")
	return NewHTMLElement(rawElement)
}

// PictureInPictureEnabled Returns a Boolean value indicating whether the document is allowed to use the Picture-in-Picture API.
func (d *Document) PictureInPictureEnabled() bool {
	return d.Js().Get("pictureInPictureEnabled").Bool()
}

// Plugins Returns a list of the <embed> elements within the current document.
func (d *Document) Plugins() *HTMLCollection {
	rawPlugins := d.Js().Get("plugins")
	return NewHTMLCollection(rawPlugins)
}

// PointerLockElement Returns the element that is currently being pointed at by the mouse pointer, or null if the mouse is not being used.
func (d *Document) PointerLockElement() *HTMLElement {
	rawElement := d.Js().Get("pointerLockElement")
	return NewHTMLElement(rawElement)
}

// Scripts Returns a live HTMLCollection containing all the scripts of the current document.
func (d *Document) Scripts() *HTMLCollection {
	rawScripts := d.Js().Get("scripts")
	return NewHTMLCollection(rawScripts)
}

// ScrollingElement Returns the Element that is the root element of the document (for example, the <html> element for HTML documents).
func (d *Document) ScrollingElement() *HTMLElement {
	rawElement := d.Js().Get("scrollingElement")
	return NewHTMLElement(rawElement)
}

// StyleSheets Returns a StyleSheetList of CSSStyleSheet objects for stylesheets explicitly linked into or embedded in a document.
func (d *Document) StyleSheets() js.Value {
	log.Println("StyleSheets is not implemented")
	return d.Js().Get("styleSheets")
}

// Timeline Returns the timeline associated with the document.
func (d *Document) Timeline() js.Value {
	log.Println("Timeline is not implemented")
	return d.Js().Get("timeline")
}

type DocumentVisibilityState string

const (
	VisibilityStateHidden    DocumentVisibilityState = "hidden"
	VisibilityStateVisible   DocumentVisibilityState = "visible"
	VisibilityStatePrerender DocumentVisibilityState = "prerender"
	VisibilityStateUnloaded  DocumentVisibilityState = "unloaded"
)

// VisibilityState Returns the visibility state of the document.
func (d *Document) VisibilityState() DocumentVisibilityState {
	state := d.Js().Get("visibilityState").String()
	return DocumentVisibilityState(state)
}

// Cookie Returns the cookie of the current document.
func (d *Document) Cookie() string {
	return d.Js().Get("cookie").String()
}

func (d *Document) SetCookie(cookie string) {
	d.Js().Set("cookie", cookie)
}

// DefaultView Returns the window object associated with the document or null if none is available.
func (d *Document) DefaultView() *Window {
	rawWindow := d.Js().Get("defaultView")
	return NewWindow(rawWindow)
}

type DocumentDesignMode string

const (
	DesignModeOn  DocumentDesignMode = "on"
	DesignModeOff DocumentDesignMode = "off"
)

// DesignMode Returns or sets a string that indicates whether the document is in design mode.
func (d *Document) DesignMode() DocumentDesignMode {
	return DocumentDesignMode(d.Js().Get("designMode").String())
}

func (d *Document) SetDesignMode(mode DocumentDesignMode) {
	d.Js().Set("designMode", string(mode))
}

type DocumentDir string

const (
	DirLTR DocumentDir = "ltr"
	DirRTL DocumentDir = "rtl"
)

// Dir Returns or sets the directionality of the text of the current element.
func (d *Document) Dir() DocumentDir {
	return DocumentDir(d.Js().Get("dir").String())
}

func (d *Document) SetDir(dir DocumentDir) {
	d.Js().Set("dir", string(dir))
}

// FullscreenEnabled Returns a Boolean value indicating whether or not the document can be viewed in fullscreen mode.
func (d *Document) FullscreenEnabled() bool {
	return d.Js().Get("fullscreenEnabled").Bool()
}

// LastModified Returns the date on which the document was last modified.
func (d *Document) LastModified() time.Time {
	date := d.Js().Get("lastModified").String()
	t, err := time.Parse(time.RFC3339, date)

	if err != nil {
		panic("browser time format is not RFC3339")
	}

	return t
}

// Location Returns the Location object, which contains information about the URL of the document and provides methods for changing that URL.
func (d *Document) Location() *Location {
	rawLocation := d.Js().Get("location")
	return NewLocation(rawLocation)
}

func (d *Document) SetLocation(location string) {
	d.Js().Set("location", location)
}

type ReadyState string

const (
	ReadyStateLoading     ReadyState = "loading"
	ReadyStateInteractive ReadyState = "interactive"
	ReadyStateComplete    ReadyState = "complete"
)

// ReadyState Returns the loading status of the document.
func (d *Document) ReadyState() ReadyState {
	return ReadyState(d.Js().Get("readyState").String())
}

// Referrer Returns the URI of the page that linked to this page.
func (d *Document) Referrer() string {
	return d.Js().Get("referrer").String()
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

//
// Document HTMLElement methods
//

// AdoptNode Adopts a node from an external document.
func (d *Document) AdoptNode(node *HTMLElement) *HTMLElement {
	rawNode := d.Js().Call("adoptNode", node.Js())
	return NewHTMLElement(rawNode)
}

// Append Adds a node to the end of the list of children of a specified parent node.
func (d *Document) Append(node *HTMLElement) *Document {
	d.Js().Call("append", node.Js())
	return d
}

// BrowsingTopics Returns a list of the browsing topics.
func (d *Document) BrowsingTopics() js.Value {
	log.Println("BrowsingTopics is not implemented")
	return d.Js().Get("browsingTopics")
}

// CaretPositionFromPoint Returns the CaretPosition object representing the caret position at the given point.
func (d *Document) CaretPositionFromPoint(x, y int) *CaretPosition {
	rawPosition := d.Js().Call("caretPositionFromPoint", x, y)
	return NewCaretPosition(rawPosition)
}

// CaretRangeFromPoint Returns the Range object representing the caret position at the given point.
func (d *Document) CaretRangeFromPoint(x, y int) *Range {
	rawRange := d.Js().Call("caretRangeFromPoint", x, y)
	return NewRange(rawRange)
}

// CreateAttribute Creates an attribute node.
func (d *Document) CreateAttribute(name string) *Attr {
	rawAttr := d.Js().Call("createAttribute", name)
	return NewAttr(rawAttr)
}

// CreateAttributeNS Creates an attribute node with a specified namespace and name.
func (d *Document) CreateAttributeNS(namespaceURI, qualifiedName string) *Attr {
	rawAttr := d.Js().Call("createAttributeNS", namespaceURI, qualifiedName)
	return NewAttr(rawAttr)
}

// CreateCDATASection Creates a CDATASection node.
func (d *Document) CreateCDATASection(data string) js.Value {
	log.Println("CreateCDATASection is not implemented")
	return d.Js().Call("createCDATASection", data)
}

// CreateComment Creates a Comment node with the specified text.
func (d *Document) CreateComment(data string) *Comment {
	rawComment := d.Js().Call("createComment", data)
	return NewComment(rawComment)
}

// CreateDocumentFragment Creates a new DocumentFragment.
func (d *Document) CreateDocumentFragment() *DocumentFragment {
	rawFragment := d.Js().Call("createDocumentFragment")
	return NewDocumentFragment(rawFragment)
}

// CreateElement Creates the HTML element specified by tagName.
func (d *Document) CreateElement(tagName list.Type) *HTMLElement {
	rawElement := d.Js().Call("createElement", string(tagName))
	return NewHTMLElement(rawElement)
}

// CreateElementNS Creates an element with the specified namespace URI and qualified name.
func (d *Document) CreateElementNS(namespaceURI, qualifiedName string) *HTMLElement {
	rawElement := d.Js().Call("createElementNS", namespaceURI, qualifiedName)
	return NewHTMLElement(rawElement)
}

// CreateEvent Creates an event of the type specified.
func (d *Document) CreateEvent(eventType string) js.Value {
	log.Println("CreateEvent is not implemented")
	return d.Js().Call("createEvent", eventType)
}

type NodeFilter int

const (
	NodeFilterShowAll                   NodeFilter = 0xFFFFFFFF
	NodeFilterShowAttribute                        = 0x2
	NodeFilterShowCDataSection                     = 0x8
	NodeFilterShowComment                          = 0x80
	NodeFilterShowDocument                         = 0x100
	NodeFilterShowDocumentFragment                 = 0x400
	NodeFilterShowDocumentType                     = 0x200
	NodeFilterShowElement                          = 0x1
	NodeFilterShowEntity                           = 0x20
	NodeFilterShowEntityReference                  = 0x10
	NodeFilterShowNotation                         = 0x800
	NodeFilterShowProcessingInstruction            = 0x40
	NodeFilterShowText                             = 0x4
)

// CreateNodeIterator Creates a NodeIterator object.
func (d *Document) CreateNodeIterator(root *Node, whatToShow NodeFilter, filter js.Func) *NodeIterator {
	rawIterator := d.Js().Call("createNodeIterator", root.Js(), int(whatToShow), filter)
	return NewNodeIterator(rawIterator)
}

// CreateProcessingInstruction Creates a ProcessingInstruction node.
func (d *Document) CreateProcessingInstruction(target, data string) js.Value {
	log.Println("CreateProcessingInstruction is not implemented")
	return d.Js().Call("createProcessingInstruction", target, data)
}

// CreateRange Creates a Range object.
func (d *Document) CreateRange() *Range {
	rawRange := d.Js().Call("createRange")
	return NewRange(rawRange)
}

// CreateTextNode Creates a Text node.
func (d *Document) CreateTextNode(data string) *Text {
	rawText := d.Js().Call("createTextNode", data)
	return NewText(rawText)
}

// createTreeWalker Creates a TreeWalker object.
func (d *Document) createTreeWalker(root *Node, whatToShow NodeFilter, filter *js.Func) *TreeWalker {
	var rawWalker js.Value

	if filter == nil {
		rawWalker = d.Js().Call("createTreeWalker", root.Js(), int(whatToShow))
	} else {
		rawWalker = d.Js().Call("createTreeWalker", root.Js(), int(whatToShow), &filter)
	}

	return NewTreeWalker(rawWalker)
}

// ElementFromPoint Returns the element from the document whose elementFromPoint method is being called which is the topmost element which lies under the given point.
func (d *Document) ElementFromPoint(x, y int) *Element {
	rawElement := d.Js().Call("elementFromPoint", x, y)
	return NewElement(rawElement)
}

// ElementsFromPoint Returns a list of all elements under the given point.
func (d *Document) ElementsFromPoint(x, y int) []*Element {
	rawElements := d.Js().Call("elementsFromPoint", x, y)

	if rawElements.IsNull() || rawElements.IsUndefined() {
		return nil
	}

	elements := make([]*Element, rawElements.Length())

	for i := 0; i < rawElements.Length(); i++ {
		elements[i] = NewElement(rawElements.Index(i))
	}

	return elements
}

// ExitFullscreen Exits full-screen mode.
func (d *Document) ExitFullscreen() *promise.Promise {
	return promise.New(d.Js().Call("exitFullscreen"))
}

// ExitPictureInPicture Exits Picture-in-Picture mode.
func (d *Document) ExitPictureInPicture() *promise.Promise {
	return promise.New(d.Js().Call("exitPictureInPicture"))
}

// ExitPointerLock Exits pointer lock.
func (d *Document) ExitPointerLock() {
	d.Js().Call("exitPointerLock")
}

// GetAnimations Returns a list of all Animation objects currently in effect on the element.
func (d *Document) GetAnimations() js.Value {
	log.Println("GetAnimations is not implemented")
	return d.Js().Call("getAnimations")
}

// GetElementById Returns a reference to the element by its ID.
func (d *Document) GetElementById(id string) *Element {
	rawElement := d.Js().Call("getElementById", id)
	return NewElement(rawElement)
}

// GetElementsByClassName Returns an array-like object of all child elements which have all of the given class names.
func (d *Document) GetElementsByClassName(classNames string) []*Element {
	rawElements := d.Js().Call("getElementsByClassName", classNames)
	return NewElementList(rawElements)
}

// GetElementsByTagName Returns an HTMLCollection of elements with the given tag name.
func (d *Document) GetElementsByTagName(tagName string) []*Element {
	rawElements := d.Js().Call("getElementsByTagName", tagName)
	return NewElementList(rawElements)
}

// GetElementsByTagNameNS Returns an HTMLCollection of elements with the given tag name in the namespace.
func (d *Document) GetElementsByTagNameNS(namespace, tagName string) []*Element {
	rawElements := d.Js().Call("getElementsByTagNameNS", namespace, tagName)
	return NewElementList(rawElements)
}

// GetSelection Returns a Selection object representing the range of text selected by the user, or the current position of the caret.
func (d *Document) GetSelection() *Selection {
	rawSelection := d.Js().Call("getSelection")
	return NewSelection(rawSelection)
}

// HasStorageAccess Returns a Promise that resolves with a boolean indicating whether the document has access to the storage area.
func (d *Document) HasStorageAccess() *promise.Promise {
	return promise.New(d.Js().Call("hasStorageAccess"))
}

// ImportNode Imports a node from another document.
func (d *Document) ImportNode(node *Node, deep bool) *Node {
	rawNode := d.Js().Call("importNode", node.Js(), deep)
	return NewNode(rawNode)
}

// Prepend Inserts a set of Node objects or DOMString objects before the first child of the ParentNode.
func (d *Document) Prepend(nodes ...*Node) {
	for _, node := range nodes {
		d.Js().Call("prepend", node.Js())
	}
}

// QuerySelector Returns the first element that is a descendant of the element on which it is invoked that matches the specified group of selectors.
func (d *Document) QuerySelector(selectors string) *Element {
	rawElement := d.Js().Call("querySelector", selectors)
	return NewElement(rawElement)
}

// QuerySelectorAll Returns a static (not live) NodeList representing a list of elements that match the specified group of selectors.
func (d *Document) QuerySelectorAll(selectors string) []*Element {
	rawElements := d.Js().Call("querySelectorAll", selectors)
	return NewElementList(rawElements)
}

// ReplaceChildren Replaces the existing children of a node with a specified new set of children.
func (d *Document) ReplaceChildren(nodes ...*Node) {
	for _, node := range nodes {
		d.Js().Call("replaceChildren", node.Js())
	}
}

// RequestStorageAccess Requests permission from the user for the page to access a storage quota.
func (d *Document) RequestStorageAccess() *promise.Promise {
	return promise.New(d.Js().Call("requestStorageAccess"))
}

// CreateExpression Creates a new XPathExpression.
func (d *Document) CreateExpression(xpathText, namespaceURLMapper string) *XPathExpression {
	var raw js.Value

	if namespaceURLMapper == "" {
		raw = d.Js().Call("createExpression", xpathText)
	} else {
		raw = d.Js().Call("createExpression", xpathText, namespaceURLMapper)
	}

	return NewXPathExpression(raw)
}

// CreateNSResolver Creates a new XPathNSResolver.
func (d *Document) CreateNSResolver(nodeResolver *Node) js.Value {
	log.Println("CreateNSResolver is not implemented")
	return d.Js().Call("createNSResolver", nodeResolver.Js())
}

// Evaluate Evaluates an XPath expression string and returns the result.
func (d *Document) Evaluate(xpathExpression *XPathExpression, contextNode *Node, namespaceResolver *js.Func, resultType XPathResultType, result *js.Value) *XPathResult {
	var raw js.Value

	if result == nil {
		raw = d.Js().Call("evaluate", xpathExpression.Js(), contextNode.Js(), &namespaceResolver, int(resultType))
	} else {
		raw = d.Js().Call("evaluate", xpathExpression.Js(), contextNode.Js(), &namespaceResolver, int(resultType), result)
	}

	return NewXPathResult(raw)
}

// Clear Clears the document's content.
func (d *Document) Clear() {
	d.Js().Call("clear")
}

// GetElementsByNames Returns a list of elements with the specified name.
func (d *Document) GetElementsByNames(name string) []*Element {
	rawElements := d.Js().Call("getElementsByName", name)
	return NewElementList(rawElements)
}

// HasFocus Returns a Boolean value indicating whether the document or any element inside the document has focus.
func (d *Document) HasFocus() bool {
	return d.Js().Call("hasFocus").Bool()
}

// Open Opens document stream for writing.
func (d *Document) Open() {
	d.Js().Call("open")
}

// Write Writes text to a document stream.
func (d *Document) Write(text string) {
	d.Js().Call("write", text)
}

// Writeln Writes a line of text to a document stream.
func (d *Document) Writeln(text string) {
	d.Js().Call("writeln", text)
}

//
// Custom CreateElement methods
//

func (d *Document) CreateCanvasElement() *CanvasElement {
	c := d.CreateElement("canvas")
	return NewCanvasElement(c)
}
