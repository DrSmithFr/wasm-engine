package element

import "syscall/js"

// FormDataEvent https://developer.mozilla.org/en-US/docs/Web/API/FormDataEvent
type FormDataEvent struct {
	*Event
}

// FormData Returns a FormData object representing the data contained in the form.
// todo
func (e *FormDataEvent) FormData() js.Value {
	return e.Js().Get("formData")
}
