package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Location https://developer.mozilla.org/en-US/docs/Web/API/Location
type Location struct {
	*wasm.Entity
}

func NewLocation(raw js.Value) *Location {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Location{
		Entity: wasm.New(raw),
	}
}

// AncestorOrigins Returns the ancestor origins.
func (l *Location) AncestorOrigins() *DOMStringList {
	return NewDOMStringList(l.Js().Get("ancestorOrigins"))
}

// Href Returns the entire URL.
func (l *Location) Href() string {
	return l.Js().Get("href").String()
}

// SetHref Sets the entire URL.
func (l *Location) SetHref(href string) *Location {
	l.Js().Set("href", href)
	return l
}

type ProtocolType string

const (
	ProtocolTypeHTTP  ProtocolType = "http:"
	ProtocolTypeHTTPS ProtocolType = "https:"
)

// Protocol Returns the protocol. (http: or https:)
func (l *Location) Protocol() ProtocolType {
	return ProtocolType(l.Js().Get("protocol").String())
}

// SetProtocol Sets the protocol. (http: or https:)
func (l *Location) SetProtocol(protocol ProtocolType) *Location {
	l.Js().Set("protocol", string(protocol))
	return l
}

// Host Returns the host and port number.
func (l *Location) Host() string {
	return l.Js().Get("host").String()
}

// SetHost Sets the host and port number.
func (l *Location) SetHost(host string) *Location {
	l.Js().Set("host", host)
	return l
}

// Hostname Returns the host name.
func (l *Location) Hostname() string {
	return l.Js().Get("hostname").String()
}

// SetHostname Sets the host name.
func (l *Location) SetHostname(hostname string) *Location {
	l.Js().Set("hostname", hostname)
	return l
}

// Port Returns the port number.
func (l *Location) Port() int {
	return l.Js().Get("port").Int()
}

// SetPort Sets the port number.
func (l *Location) SetPort(port int) *Location {
	l.Js().Set("port", port)
	return l
}

// Pathname Returns the path name.
func (l *Location) Pathname() string {
	return l.Js().Get("pathname").String()
}

// SetPathname Sets the path name.
func (l *Location) SetPathname(pathname string) *Location {
	l.Js().Set("pathname", pathname)
	return l
}

// Search Returns the query string.
func (l *Location) Search() string {
	return l.Js().Get("search").String()
}

// SetSearch Sets the query string.
func (l *Location) SetSearch(search string) *Location {
	l.Js().Set("search", search)
	return l
}

// Hash Returns the anchor name.
func (l *Location) Hash() string {
	return l.Js().Get("hash").String()
}

// SetHash Sets the anchor name.
func (l *Location) SetHash(hash string) *Location {
	l.Js().Set("hash", hash)
	return l
}

// Origin Returns the origin.
func (l *Location) Origin() string {
	return l.Js().Get("origin").String()
}

// Assign Loads a new document.
func (l *Location) Assign(url string) {
	l.Js().Call("assign", url)
}

// Reload Reloads the current document.
func (l *Location) Reload() {
	l.Js().Call("reload")
}

// Replace Replaces the current document.
func (l *Location) Replace(url string) {
	l.Js().Call("replace", url)
}

// ToString Returns the string representation of the object.
func (l *Location) ToString() string {
	return l.Js().Call("toString").String()
}
