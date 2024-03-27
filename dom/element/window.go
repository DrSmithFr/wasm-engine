package element

import (
	"go-webgl/dom/promise"
	"syscall/js"
)

// for reference: https://developer.mozilla.org/en-US/docs/Web/API/Window

type Window struct {
	*EventTarget
	document *Document
}

// CurrentWindow enforce document Singleton
var CurrentWindow *Window

func LoadWindow() *Window {
	if CurrentWindow != nil {
		return CurrentWindow
	}

	CurrentWindow = &Window{
		EventTarget: NewEventTarget(js.Global()),
	}
	CurrentWindow.document = LoadDocument()

	return CurrentWindow
}

func NewWindow(raw js.Value) *Window {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Window{
		EventTarget: NewEventTarget(raw),
	}
}

//
// Window methods
//

// Caches Returns the CacheStorage object associated with the current context.
func (d *Window) Caches() *CacheStorage {
	return NewCacheStorage(d.Js().Get("caches"))
}

// ClientInformation Returns the Navigator object that represents the browser.
func (d *Window) ClientInformation() js.Value {
	// @TODO: return Navigator object
	return d.Js().Get("clientInformation")
}

// Closed Returns a Boolean value indicating whether a window has been closed or not.
func (d *Window) Closed() bool {
	return d.Js().Get("closed").Bool()
}

// Console Returns a reference to the console object, which provides methods for logging information to the browser's console.
func (d *Window) Console() js.Value {
	// @TODO: return Console object
	return d.Js().Get("console")
}

// CookieStore Returns the CookieStore object associated with the current context.
func (d *Window) CookieStore() js.Value {
	// @TODO: return CookieStore object
	return d.Js().Get("cookieStore")
}

// Credentialless Returns a boolean that indicates whether the current document was loaded inside a credentialless
func (d *Window) Credentialless() bool {
	return d.Js().Get("credentialless").Bool()
}

// CrossOriginIsolated Returns a boolean that indicates whether the current context is cross-origin isolated.
func (d *Window) CrossOriginIsolated() bool {
	return d.Js().Get("crossOriginIsolated").Bool()
}

// Crypto Returns the Crypto object associated with the global object.
func (d *Window) Crypto() js.Value {
	// @TODO: return Crypto object
	return d.Js().Get("crypto")
}

// CustomElements Returns the CustomElementRegistry object, which can be used to register new custom elements and get information about previously registered custom elements.
func (d *Window) CustomElements() js.Value {
	// @TODO: return CustomElementRegistry object
	return d.Js().Get("customElements")
}

// DevicePixelRatio Returns the ratio of the resolution in physical pixels to the resolution in CSS pixels for the current display device.
func (d *Window) DevicePixelRatio() float64 {
	return d.Js().Get("devicePixelRatio").Float()
}

// Document Returns the Document object that represents the document in the specified window.
func (d *Window) Document() *Document {
	return d.document
}

// DocumentPictureInPicture Returns the PictureInPictureWindow object associated with the current context.
func (d *Window) DocumentPictureInPicture() js.Value {
	// @TODO: return PictureInPictureWindow object
	return d.Js().Get("documentPictureInPicture")
}

// Fence Returns the Fence object associated with the current context.
func (d *Window) Fence() js.Value {
	// @TODO: return Fence object
	return d.Js().Get("fence")
}

// FrameElement Returns the element (such as an <iframe> element) in which the window is embedded, or null if the window is top-level.
func (d *Window) FrameElement() js.Value {
	// @TODO: return Element object
	return d.Js().Get("frameElement")
}

// Frames Returns the WindowProxy object for the window that contains the frame in which the current window is nested.
func (d *Window) Frames() js.Value {
	// @TODO: return WindowProxy object
	return d.Js().Get("frames")
}

// FullScreen Returns a Boolean value indicating whether the window is displayed in full screen mode.
func (d *Window) FullScreen() bool {
	return d.Js().Get("fullScreen").Bool()
}

// History Returns a reference to the History object, which provides an interface for manipulating the browser session history (pages visited in the tab or frame that the current page is loaded in).
func (d *Window) History() js.Value {
	// @TODO: return History object
	return d.Js().Get("history")
}

// IndexedDB Returns the IDBFactory object for the window.
func (d *Window) IndexedDB() js.Value {
	// @TODO: return IDBFactory object
	return d.Js().Get("indexedDB")
}

// InnerHeight Returns the height of the content area of the browser window including, if rendered, the horizontal scrollbar.
func (d *Window) InnerHeight() int {
	return d.Js().Get("innerHeight").Int()
}

// InnerWidth Returns the width of the content area of the browser window including, if rendered, the vertical scrollbar.
func (d *Window) InnerWidth() int {
	return d.Js().Get("innerWidth").Int()
}

// InnerSize Returns the width and height of the content area of the browser window including, if rendered, the vertical scrollbar.
func (d *Window) InnerSize() (int, int) {
	return d.InnerWidth(), d.InnerHeight()
}

// IsSecureContext Returns a boolean value indicating whether the current context is secure (true) or not (false).
func (d *Window) IsSecureContext() bool {
	return d.Js().Get("isSecureContext").Bool()
}

// LaunchQueue Returns the LaunchQueue object associated with the current context.
func (d *Window) LaunchQueue() js.Value {
	// @TODO: return LaunchQueue object
	return d.Js().Get("launchQueue")
}

// Length Returns the number of frames (either <iframe> or <frame> elements) in the window.
func (d *Window) Length() int {
	return d.Js().Get("length").Int()
}

// LocalStorage Returns a reference to the local storage object used to store data that may only be accessed by the origin that created it.
func (d *Window) LocalStorage() js.Value {
	// @TODO: return Storage object
	return d.Js().Get("localStorage")
}

// Location Returns a Location object, which contains information about the URL of the document and provides methods for changing that URL.
func (d *Window) Location() *Location {
	return NewLocation(d.Js().Get("location"))
}

// SetLocation Sets the Location string
func (d *Window) SetLocation(url string) {
	d.Js().Set("location", url)
}

// Locationbar Returns the locationbar object, which contains the locationbar object for the window.
func (d *Window) Locationbar() js.Value {
	// @TODO: return locationbar object
	return d.Js().Get("locationbar")
}

// Menubar Returns the menubar object, which contains the menubar object for the window.
func (d *Window) Menubar() js.Value {
	//@TODO: return menubar object
	return d.Js().Get("menubar")
}

// Name Returns the name of the window.
func (d *Window) Name() string {
	return d.Js().Get("name").String()
}

// SetName Sets the name of the window.
func (d *Window) SetName(name string) {
	d.Js().Set("name", name)
}

// Navigation Returns the Navigation object associated with the current context.
func (d *Window) Navigation() js.Value {
	//@TODO: return Navigation object
	return d.Js().Get("navigation")
}

// Navigator Returns the Navigator object that represents the browser.
func (d *Window) Navigator() js.Value {
	//@TODO: return Navigator object
	return d.Js().Get("navigator")
}

// Opener Returns a reference to the window that opened this current window.
func (d *Window) Opener() *Window {
	return NewWindow(d.Js().Get("opener"))
}

// Origin Returns the origin of the global object, serialized as a string.
func (d *Window) Origin() string {
	return d.Js().Get("origin").String()
}

// OriginAgentCluster Returns the OriginAgentCluster object associated with the current context.
func (d *Window) OriginAgentCluster() bool {
	return d.Js().Get("originAgentCluster").Bool()
}

// OuterHeight Returns the height of the outside of the browser window.
func (d *Window) OuterHeight() int {
	return d.Js().Get("outerHeight").Int()
}

// OuterWidth Returns the width of the outside of the browser window.
func (d *Window) OuterWidth() int {
	return d.Js().Get("outerWidth").Int()
}

// OuterSize Returns the width and height of the outside of the browser window.
func (d *Window) OuterSize() (int, int) {
	return d.OuterWidth(), d.OuterHeight()
}

// PageXOffset Returns the number of pixels that the document has already been scrolled horizontally.
func (d *Window) PageXOffset() int {
	return d.Js().Get("pageXOffset").Int()
}

// PageYOffset Returns the number of pixels that the document has already been scrolled vertically.
func (d *Window) PageYOffset() int {
	return d.Js().Get("pageYOffset").Int()
}

func (d *Window) PageOffset() (int, int) {
	return d.PageXOffset(), d.PageYOffset()
}

// Parent Returns a reference to the parent of the current window or a reference to the current window if there is no parent.
func (d *Window) Parent() js.Value {
	//@TODO: return Window or subframe object
	return d.Js().Get("parent")
}

// Performance Returns the Performance object associated with the current context.
func (d *Window) Performance() js.Value {
	//@TODO: return Performance object
	return d.Js().Get("performance")
}

// Personalbar Returns the personalbar object, which contains the personalbar object for the window.
func (d *Window) Personalbar() js.Value {
	//@ TODO: return personalbar object
	return d.Js().Get("personalbar")
}

// Scheduler Returns the Scheduler object associated with the current context.
func (d *Window) Scheduler() js.Value {
	//@TODO: return Scheduler object
	return d.Js().Get("scheduler")
}

// Screen Returns the Screen object associated with the window.
func (d *Window) Screen() *Screen {
	return NewScreen(d.Js().Get("screen"))
}

// ScreenX Returns the horizontal distance of the left border of the user's browser from the left side of the screen.
func (d *Window) ScreenX() int {
	return d.Js().Get("screenX").Int()
}

// ScreenLeft Returns the horizontal distance of the left border of the user's browser from the left side of the screen.
func (d *Window) ScreenLeft() int {
	return d.Js().Get("screenLeft").Int()
}

// ScreenY Returns the vertical distance of the top border of the user's browser from the top side of the screen.
func (d *Window) ScreenY() int {
	return d.Js().Get("screenY").Int()
}

// ScreenTop Returns the vertical distance of the top border of the user's browser from the top side of the screen.
func (d *Window) ScreenTop() int {
	return d.Js().Get("screenTop").Int()
}

// Scrollbars Returns the scrollbars object, which contains the scrollbars object for the window.
func (d *Window) Scrollbars() js.Value {
	//@TODO: return scrollbars object
	return d.Js().Get("scrollbars")
}

// ScrollX Returns the number of pixels that the document has already been scrolled horizontally.
func (d *Window) ScrollX() int {
	return d.Js().Get("scrollX").Int()
}

// ScrollY Returns the number of pixels that the document has already been scrolled vertically.
func (d *Window) ScrollY() int {
	return d.Js().Get("scrollY").Int()
}

// Self Returns the Window object itself.
func (d *Window) Self() *Window {
	return d
}

// SessionStorage Returns a reference to the session storage object used to store data that may only be accessed by the origin that created it.
func (d *Window) SessionStorage() js.Value {
	//@TODO: return Storage object
	return d.Js().Get("sessionStorage")
}

// SpeechSynthesis Returns the SpeechSynthesis object, which is the entry point into using Web Speech API speech synthesis functionality.
func (d *Window) SpeechSynthesis() js.Value {
	//@TODO: return SpeechSynthesis object
	return d.Js().Get("speechSynthesis")
}

// StatusBar Returns the statusbar object, which contains the statusbar object for the window.
func (d *Window) StatusBar() js.Value {
	//@TODO: return statusbar object
	return d.Js().Get("statusbar")
}

// Top Returns a reference to the topmost window in the window hierarchy.
func (d *Window) Top() js.Value {
	//@TODO: return Window object
	return d.Js().Get("top")
}

// VisualViewport Returns the VisualViewport object associated with the current context.
func (d *Window) VisualViewport() js.Value {
	//@ TODO: return VisualViewport object
	return d.Js().Get("visualViewport")
}

// Window Returns a reference to the window that opened this current window.
func (d *Window) Window() *Window {
	return d
}

// Atob Decodes a string of data which has been encoded using base-64 encoding.
func (d *Window) Atob(data string) string {
	return d.Js().Call("atob", data).String()
}

// Alert Displays an alert dialog with the specified content and an OK button.
func (d *Window) Alert(message string) {
	d.Js().Call("alert", message)
}

// Blur Removes focus from the current window.
func (d *Window) Blur() {
	d.Js().Call("blur")
}

// Btoa Creates a base-64 encoded ASCII string from a string of binary data.
func (d *Window) Btoa(data string) string {
	return d.Js().Call("btoa", data).String()
}

// CancelAnimationFrame Cancels an animation frame request previously scheduled through a call to window.requestAnimationFrame().
func (d *Window) CancelAnimationFrame(requestID int) {
	d.Js().Call("cancelAnimationFrame", requestID)
}

// CancelIdleCallback Cancels a callback previously scheduled with requestIdleCallback().
func (d *Window) CancelIdleCallback(handle int) {
	d.Js().Call("cancelIdleCallback", handle)
}

// ClearInterval Cancels the repeated execution set using setInterval().
func (d *Window) ClearInterval(handle int) {
	d.Js().Call("clearInterval", handle)
}

// ClearTimeout Cancels the repeated execution set using setTimeout().
func (d *Window) ClearTimeout(handle int) {
	d.Js().Call("clearTimeout", handle)
}

// Close Closes the current window, or a window with a specified name.
func (d *Window) Close() {
	d.Js().Call("close")
}

// Confirm Displays a dialog with an optional message prompting the user to input some text.
func (d *Window) Confirm(message string) bool {
	return d.Js().Call("confirm", message).Bool()
}

// CreateImageBitmap Creates a new ImageBitmap object from the given image data.
// TODO

// Fetch Sends a request to the server and receives a response.
// TODO

// Focus Sets focus to the current window.
func (d *Window) Focus() {
	d.Js().Call("focus")
}

// GetComputedStyle Returns an object containing the values of all CSS properties of an element, after applying active stylesheets and resolving any basic computation those values may contain.
func (d *Window) GetComputedStyle(element *Element, pseudoElement string) *CSSStyleDeclaration {
	if pseudoElement == "" {
		return NewCSSStyleDeclaration(d.Js().Call("getComputedStyle", element.Js()))
	}

	return NewCSSStyleDeclaration(d.Js().Call("getComputedStyle", element.Js(), pseudoElement))
}

// GetScreenDetails Returns the screen details.
func (d *Window) GetScreenDetails() *promise.Promise {
	return promise.New(d.Js().Call("getScreenDetails"))
}

// GetSelection Returns a Selection object representing the range of text selected by the user, or the current position of the caret.
func (d *Window) GetSelection() *Selection {
	return NewSelection(d.Js().Call("getSelection"))
}

// MatchMedia Returns a MediaQueryList object representing the specified media query string.
func (d *Window) MatchMedia(query string) *MediaQueryList {
	return NewMediaQueryList(d.Js().Call("matchMedia", query))
}

// MoveBy Moves the current window by a specified amount.
func (d *Window) MoveBy(x, y int) {
	d.Js().Call("moveBy", x, y)
}

// MoveTo Moves the current window to the specified position.
func (d *Window) MoveTo(x, y int) {
	d.Js().Call("moveTo", x, y)
}

type UrlTarget string

const (
	UrlTargetSelf        UrlTarget = "_self"
	UrlTargetBlank       UrlTarget = "_blank"
	UrlTargetParent      UrlTarget = "_parent"
	UrlTargetTop         UrlTarget = "_top"
	UrlTargetUnfencedTop UrlTarget = "_unfenced-top"
)

type UrlFeatures struct {
	Popup bool

	Width  int
	Height int

	ScreenX int
	ScreenY int

	NoOpener   bool
	NoReferrer bool
}

// Open Opens a new browser window.
func (d *Window) Open(url string, target UrlTarget, features *UrlFeatures) *Window {
	if features == nil {
		return NewWindow(d.Js().Call("open", url, string(target)))
	}

	return NewWindow(d.Js().Call("open", url, target, js.ValueOf(features)))
}

// PostMessage Sends a cross-origin message to a window.
// TODO

// Print Prints the current window.
func (d *Window) Print() {
	d.Js().Call("print")
}

// Prompt Displays a dialog box that prompts the visitor for input.
func (d *Window) Prompt(message, defaultValue string) string {
	return d.Js().Call("prompt", message, defaultValue).String()
}

// QueueMicrotask Adds a callback to the microtask queue.
func (d *Window) QueueMicrotask(callback func() interface{}) {
	task := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return callback()
	})

	d.Js().Call("queueMicrotask", task)
}

// RequestAnimationFrame Requests an animation frame for a function before the next repaint.
func (d *Window) RequestAnimationFrame(callback func() interface{}) int {
	task := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return callback()
	})

	return d.Js().Call("requestAnimationFrame", task).Int()
}

// RequestIdleCallback Schedules a function to run when the main thread is idle.
func (d *Window) RequestIdleCallback(callback func() interface{}) int {
	task := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return callback()
	})

	return d.Js().Call("requestIdleCallback", task).Int()
}

// ResizeBy Resizes the window by a specified amount.
func (d *Window) ResizeBy(x, y int) {
	d.Js().Call("resizeBy", x, y)
}

// ResizeTo Resizes the window to the specified width and height.
func (d *Window) ResizeTo(width, height int) {
	d.Js().Call("resizeTo", width, height)
}

// Scroll Scrolls the window to a particular place in the document.
func (d *Window) Scroll(x, y int) {
	d.Js().Call("scroll", x, y)
}

type ScrollBehavior string

const (
	ScrollBehaviorAuto    ScrollBehavior = "auto"
	ScrollBehaviorSmooth  ScrollBehavior = "smooth"
	ScrollBehaviorInstant ScrollBehavior = "instant"
)

type ScrollOptions struct {
	top      int
	left     int
	behavior ScrollBehavior
}

// ScrollOpt Scrolls the window to a particular place in the document.
func (d *Window) ScrollOpt(opt ScrollOptions) {
	d.Js().Call("scroll", js.ValueOf(opt))
}

// ScrollBy Scrolls the window by a specified number of pixels.
func (d *Window) ScrollBy(x, y int) {
	d.Js().Call("scrollBy", x, y)
}

// ScrollByOpt Scrolls the window by a specified number of pixels.
func (d *Window) ScrollByOpt(opt ScrollOptions) {
	d.Js().Call("scrollBy", js.ValueOf(opt))
}

// ScrollTo Scrolls the window to a particular place in the document.
func (d *Window) ScrollTo(x, y int) {
	d.Js().Call("scrollTo", x, y)
}

// ScrollToOpt Scrolls the window to a particular place in the document.
func (d *Window) ScrollToOpt(opt ScrollOptions) {
	d.Js().Call("scrollTo", js.ValueOf(opt))
}

// SetInterval Calls a function or evaluates an expression at specified intervals (in milliseconds).
func (d *Window) SetInterval(callback func() interface{}, ms int) int {
	task := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return callback()
	})

	return d.Js().Call("setInterval", task, ms).Int()
}

// SetTimeout Calls a function or evaluates an expression after a specified number of milliseconds.
func (d *Window) SetTimeout(callback func() interface{}, ms int) int {
	task := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		return callback()
	})

	return d.Js().Call("setTimeout", task, ms).Int()
}

// Stop Stops the window from loading.
func (d *Window) Stop() {
	d.Js().Call("stop")
}

// StrucuredClone Creates a structured clone of a value.
// TODO
