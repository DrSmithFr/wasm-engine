package dom

import (
	"go-webgl/dom/element"
	"go-webgl/dom/wasm"
	"syscall/js"
)

func Window() *element.Window {
	return element.LoadWindow()
}

func Document() *element.Document {
	return element.LoadDocument()
}

func Location() *element.Location {
	return Document().Location()
}

type cacheStorage struct {
	*wasm.Entity
}

func CacheStorage() *cacheStorage {
	return &cacheStorage{
		Entity: wasm.New(js.Global().Get("CacheStorage")),
	}
}
