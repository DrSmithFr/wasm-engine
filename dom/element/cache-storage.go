package element

import (
	"go-webgl/dom/promise"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// CacheStorage https://developer.mozilla.org/en-US/docs/Web/API/CacheStorage
type CacheStorage struct {
	*wasm.Entity
}

func NewCacheStorage(raw js.Value) *CacheStorage {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &CacheStorage{
		Entity: wasm.New(raw),
	}
}

// Open Returns a Promise that resolves to the Cache object matching the cacheName.
func (cs *CacheStorage) Open(cacheName string) *promise.Promise {
	return promise.New(cs.Js().Call("open", cacheName))
}

// Match Returns a Promise that resolves to the response associated with the first matching request in the Cache object.
func (cs *CacheStorage) Match(request js.Value, option *MatchOptions) *promise.Promise {
	if option == nil {
		return promise.New(cs.Js().Call("match", request))
	}

	return promise.New(cs.Js().Call("match", request, js.ValueOf(option)))
}

// Has Returns a Promise that resolves to true if a Cache object matching the cacheName exists.
func (cs *CacheStorage) Has(cacheName string) *promise.Promise {
	return promise.New(cs.Js().Call("has", cacheName))
}

// Delete Returns a Promise that resolves to true if the Cache object matching the cacheName was successfully deleted.
func (cs *CacheStorage) Delete(cacheName string) *promise.Promise {
	return promise.New(cs.Js().Call("delete", cacheName))
}

// Keys Returns a Promise that resolves to an array of Cache keys.
func (cs *CacheStorage) Keys() *promise.Promise {
	return promise.New(cs.Js().Call("keys"))
}
