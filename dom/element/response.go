package element

import (
	"go-webgl/dom/promise"
	"go-webgl/dom/wasm"
	"syscall/js"
)

// Response https://developer.mozilla.org/en-US/docs/Web/API/Response

type ResponseStatus int

const (
	ResponseStatusUnknown ResponseStatus = 0

	ResponseStatusOk             ResponseStatus = 200
	ResponseStatusCreated        ResponseStatus = 201
	ResponseStatusAccepted       ResponseStatus = 202
	ResponseStatusNoContent      ResponseStatus = 204
	ResponseStatusResetContent   ResponseStatus = 205
	ResponseStatusPartialContent ResponseStatus = 206

	ResponseStatusMultipleChoices   ResponseStatus = 300
	ResponseStatusMovedPermanently  ResponseStatus = 301
	ResponseStatusFound             ResponseStatus = 302
	ResponseStatusSeeOther          ResponseStatus = 303
	ResponseStatusNotModified       ResponseStatus = 304
	ResponseStatusUseProxy          ResponseStatus = 305
	ResponseStatusTemporaryRedirect ResponseStatus = 307
	ResponseStatusPermanentRedirect ResponseStatus = 308

	ResponseStatusBadRequest                  ResponseStatus = 400
	ResponseStatusUnauthorized                ResponseStatus = 401
	ResponseStatusForbidden                   ResponseStatus = 403
	ResponseStatusNotFound                    ResponseStatus = 404
	ResponseStatusMethodNotAllowed            ResponseStatus = 405
	ResponseStatusNotAcceptable               ResponseStatus = 406
	ResponseStatusProxyAuthRequired           ResponseStatus = 407
	ResponseStatusRequestTimeout              ResponseStatus = 408
	ResponseStatusConflict                    ResponseStatus = 409
	ResponseStatusGone                        ResponseStatus = 410
	ResponseStatusLengthRequired              ResponseStatus = 411
	ResponseStatusPreconditionFailed          ResponseStatus = 412
	ResponseStatusPayloadTooLarge             ResponseStatus = 413
	ResponseStatusURITooLong                  ResponseStatus = 414
	ResponseStatusUnsupportedMediaType        ResponseStatus = 415
	ResponseStatusRangeNotSatisfiable         ResponseStatus = 416
	ResponseStatusExpectationFailed           ResponseStatus = 417
	ResponseStatusTeapot                      ResponseStatus = 418
	ResponseStatusMisdirectedRequest          ResponseStatus = 421
	ResponseStatusUnprocessableEntity         ResponseStatus = 422
	ResponseStatusLocked                      ResponseStatus = 423
	ResponseStatusFailedDependency            ResponseStatus = 424
	ResponseStatusTooEarly                    ResponseStatus = 425
	ResponseStatusUpgradeRequired             ResponseStatus = 426
	ResponseStatusPreconditionRequired        ResponseStatus = 428
	ResponseStatusTooManyRequests             ResponseStatus = 429
	ResponseStatusRequestHeaderFieldsTooLarge ResponseStatus = 431
	ResponseStatusUnavailableForLegalReasons  ResponseStatus = 451

	ResponseStatusInternalServerError           ResponseStatus = 500
	ResponseStatusBadGateway                    ResponseStatus = 502
	ResponseStatusServiceUnavailable            ResponseStatus = 503
	ResponseStatusGatewayTimeout                ResponseStatus = 504
	ResponseStatusHTTPVersionNotSupported       ResponseStatus = 505
	ResponseStatusVariantAlsoNegotiates         ResponseStatus = 506
	ResponseStatusInsufficientStorage           ResponseStatus = 507
	ResponseStatusLoopDetected                  ResponseStatus = 508
	ResponseStatusNotExtended                   ResponseStatus = 510
	ResponseStatusNetworkAuthenticationRequired ResponseStatus = 511
)

type Response struct {
	*wasm.Entity
}

// NewResponse returns a new Request object.
func NewResponse(request js.Value) *Response {
	return &Response{
		Entity: wasm.New(request),
	}
}

// Body Returns a ReadableStream object representing the body of the response.
func (r *Response) Body() js.Value {
	return r.Js().Get("body")
}

// BodyUsed Returns a boolean that is true if the body has been read.
func (r *Response) BodyUsed() bool {
	return r.Js().Get("bodyUsed").Bool()
}

// Headers Returns a Headers object associated with the response.
func (r *Response) Headers() *HeadersReadOnly {
	return NewHeadersReadOnly(r.Js().Get("headers"))
}

// Ok Returns a boolean indicating if the response was successful (status in the range 200-299) or not.
func (r *Response) Ok() bool {
	return r.Js().Get("ok").Bool()
}

// Redirected Returns a boolean indicating if the response is the result of a redirect.
func (r *Response) Redirected() bool {
	return r.Js().Get("redirected").Bool()
}

// Status Returns the status code of the response.
func (r *Response) Status() ResponseStatus {
	return ResponseStatus(r.Js().Get("status").Int())
}

// StatusText Returns the status message corresponding to the status code.
func (r *Response) StatusText() string {
	return r.Js().Get("statusText").String()
}

type ResponseType string

const (
	ResponseTypeBasic ResponseType = "basic"
	ResponseTypeCors  ResponseType = "cors"
)

// Type Returns the type of the response.
func (r *Response) Type() ResponseType {
	return ResponseType(r.Js().Get("type").String())
}

// URL Returns the URL of the response.
func (r *Response) URL() string {
	return r.Js().Get("url").String()
}

// ArrayBuffer Returns a promise that resolves with an ArrayBuffer representation of the response body.
func (r *Response) ArrayBuffer() *promise.Promise {
	return promise.New(r.Js().Call("arrayBuffer"))
}

// Blob Returns a promise that resolves with a Blob representation of the response body.
func (r *Response) Blob() *promise.Promise {
	return promise.New(r.Js().Call("blob"))
}

// Clone Creates a clone of a Response object.
func (r *Response) Clone() *Response {
	return NewResponse(r.Js().Call("clone"))
}

// FormData Returns a promise that resolves with a FormData representation of the response body.
func (r *Response) FormData() *promise.Promise {
	return promise.New(r.Js().Call("formData"))
}

// JSON Returns a promise that resolves with a JSON representation of the response body.
func (r *Response) JSON() *promise.Promise {
	return promise.New(r.Js().Call("json"))
}

// Text Returns a promise that resolves with a text representation of the response body.
func (r *Response) Text() *promise.Promise {
	return promise.New(r.Js().Call("text"))
}
