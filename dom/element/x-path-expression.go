package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

type XPathExpression struct {
	*wasm.Entity
}

func NewXPathExpression(raw js.Value) *XPathExpression {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &XPathExpression{
		Entity: wasm.New(raw),
	}
}

// Evaluate Evaluates an XPath expression string and returns a result of the specified type.
func (x *XPathExpression) Evaluate(contextNode *Node, resultType XPathResultType, result *js.Value) *XPathResult {
	var r js.Value
	if result == nil {
		r = x.Js().Call("evaluate", contextNode.Js())
	} else {
		r = x.Js().Call("evaluate", contextNode.Js(), int(resultType), result)
	}

	return NewXPathResult(r)

}
