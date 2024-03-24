package dom

import (
	"go-webgl/dom/element"
)

func Window() *element.Window {
	return element.LoadWindow()
}

func Document() *element.Document {
	return element.LoadDocument()
}
