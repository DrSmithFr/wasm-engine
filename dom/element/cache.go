package element

import (
	"go-webgl/dom/promise"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Cache https://developer.mozilla.org/en-US/docs/Web/API/Cache
type Cache struct {
	*wasm.Entity
}

func NewCache(raw js.Value) *Cache {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &Cache{
		Entity: wasm.New(raw),
	}
}

type MatchOptions struct {
	IgnoreSearch bool
	IgnoreMethod bool
	IgnoreVary   bool
}

// Match Returns a Promise that resolves to the response associated with the first matching request in the Cache object.
func (c *Cache) Match(request js.Value, option *MatchOptions) *promise.Promise {
	if option == nil {
		return promise.New(c.Js().Call("match", request))
	}

	return promise.New(c.Js().Call("match", request, js.ValueOf(option)))
}

// MatchAll Returns a Promise that resolves to an array of all matching responses in the Cache object.
func (c *Cache) MatchAll(request js.Value, option *MatchOptions) *promise.Promise {
	if option == nil {
		return promise.New(c.Js().Call("matchAll", request))
	}

	return promise.New(c.Js().Call("matchAll", request, js.ValueOf(option)))
}

// Add Takes a URL, fetches it, and adds the response to the given cache.
func (c *Cache) Add(request js.Value) *promise.Promise {
	return promise.New(c.Js().Call("add", request))
}

// AddAll Takes an array of URL objects and adds them to the given cache.
func (c *Cache) AddAll(requests []js.Value) *promise.Promise {
	return promise.New(c.Js().Call("addAll", requests))
}

// Put Takes both a request and its response and adds it to the given cache.
func (c *Cache) Put(request js.Value, response js.Value) *promise.Promise {
	return promise.New(c.Js().Call("put", request, response))
}

// Delete Finds the Cache entry whose key is the request, and if found, deletes the Cache entry and returns a Promise that resolves to true.
func (c *Cache) Delete(request js.Value) *promise.Promise {
	return promise.New(c.Js().Call("delete", request))
}

type MatchKeysOptions struct {
	IgnoreSearch bool
	IgnoreMethod bool
	IgnoreVary   bool
}

// Keys Returns a Promise that resolves to an array of Cache keys.
func (c *Cache) Keys(request js.Value, option *MatchKeysOptions) *promise.Promise {
	if option == nil {
		return promise.New(c.Js().Call("keys", request))
	}

	return promise.New(c.Js().Call("keys", request, js.ValueOf(option)))
}
