package element

import "syscall/js"

// PaymentRequest https://developer.mozilla.org/en-US/docs/Web/API/PaymentRequest
type PaymentRequest struct {
	*EventTarget
}

func NewPaymentRequest(entity js.Value) *PaymentRequest {
	if entity.IsNull() || entity.IsUndefined() {
		return nil
	}

	return &PaymentRequest{
		EventTarget: NewEventTarget(entity),
	}
}

// Id Returns a DOMString containing the ID of the PaymentRequest object.
func (p *PaymentRequest) Id() string {
	return p.Js().Get("id").String()
}

// CanMakePayment Returns a Promise that resolves with a Boolean value indicating whether or not the PaymentRequest can be fulfilled with any of the given payment methods.
func (p *PaymentRequest) CanMakePayment() js.Value {
	return p.Js().Call("canMakePayment")
}

// Show Shows the user agent's default user interface for requesting payment.
func (p *PaymentRequest) Show() js.Value {
	return p.Js().Call("show")
}

// Abort Cancels the payment request.
func (p *PaymentRequest) Abort() {
	p.Js().Call("abort")
}
