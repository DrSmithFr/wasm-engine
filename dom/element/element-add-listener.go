package element

import "syscall/js"

// AddBeforeMatch All Adds a beforematch event listener to the element.
func (e *Element) AddBeforeMatchAllListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "beforematchall", listener)
}

// AddScrollListener Adds a scroll event listener to the element.
func (e *Element) AddScrollListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "scroll", listener)
}

// AddScrollEndListener Adds a scrollend event listener to the element.
func (e *Element) AddScrollEndListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "scrollend", listener)
}

// AddSecurityPolicyViolationListener Adds a securitypolicyviolation event listener to the element.
func (e *Element) AddSecurityPolicyViolationListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "securitypolicyviolation", listener)
}

// AddWheelListener Adds a wheel event listener to the element.
func (e *Element) AddWheelListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "wheel", listener)
}

// AddAnimationCancelListener Adds an animationcancel event listener to the element.
func (e *Element) AddAnimationCancelListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "animationcancel", listener)
}

// AddAnimationEndListener Adds an animationend event listener to the element.
func (e *Element) AddAnimationEndListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "animationend", listener)
}

// AddAnimationIterationListener Adds an animationiteration event listener to the element.
func (e *Element) AddAnimationIterationListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "animationiteration", listener)
}

// AddAnimationStartListener Adds an animationstart event listener to the element.
func (e *Element) AddAnimationStartListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "animationstart", listener)
}

// AddCopyListener Adds a copy event listener to the element.
func (e *Element) AddCopyListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "copy", listener)
}

// AddCutListener Adds a cut event listener to the element.
func (e *Element) AddCutListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "cut", listener)
}

// AddPasteListener Adds a paste event listener to the element.
func (e *Element) AddPasteListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "paste", listener)
}

// AddCompositionEndListener Adds a compositionend event listener to the element.
func (e *Element) AddCompositionEndListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "compositionend", listener)
}

// AddCompositionStartListener Adds a compositionstart event listener to the element.
func (e *Element) AddCompositionStartListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "compositionstart", listener)
}

// AddCompositionUpdateListener Adds a compositionupdate event listener to the element.
func (e *Element) AddCompositionUpdateListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "compositionupdate", listener)
}

// AddBlurListener Adds a blur event listener to the element.
func (e *Element) AddBlurListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "blur", listener)
}

// AddFocusListener Adds a focus event listener to the element.
func (e *Element) AddFocusListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "focus", listener)
}

// AddFocusInListener Adds a focusin event listener to the element.
func (e *Element) AddFocusInListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "focusin", listener)
}

// AddFocusOutListener Adds a focusout event listener to the element.
func (e *Element) AddFocusOutListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "focusout", listener)
}

// AddFullscreenChangeListener Adds a fullscreenchange event listener to the element.
func (e *Element) AddFullscreenChangeListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "fullscreenchange", listener)
}

// AddFullscreenErrorListener Adds a fullscreenerror event listener to the element.
func (e *Element) AddFullscreenErrorListener(callback func(e *Event)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewEvent(args[0])
		callback(NewEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "fullscreenerror", listener)
}

// AddKeyDownListener Adds a keydown event listener to the element.
func (e *Element) AddKeyDownListener(callback func(e *KeyboardEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewKeyboardEvent(args[0])
		callback(NewKeyboardEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "keydown", listener)
}

// AddKeyPressListener Adds a keypress event listener to the element.
func (e *Element) AddKeyPressListener(callback func(e *KeyboardEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewKeyboardEvent(args[0])
		callback(NewKeyboardEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "keypress", listener)
}

// AddKeyUpListener Adds a keyup event listener to the element.
func (e *Element) AddKeyUpListener(callback func(e *KeyboardEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewKeyboardEvent(args[0])
		callback(NewKeyboardEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "keyup", listener)
}

// AddAuxClickListener Adds an auxclick event listener to the element.
func (e *Element) AddAuxClickListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "auxclick", listener)
}

// AddClickListener Adds a click event listener to the element.
func (e *Element) AddClickListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "click", listener)
}

// AddContextMenuListener Adds a contextmenu event listener to the element.
func (e *Element) AddContextMenuListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "contextmenu", listener)
}

// AddDblClickListener Adds a dblclick event listener to the element.
func (e *Element) AddDblClickListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "dblclick", listener)
}

// AddMouseDownListener Adds a mousedown event listener to the element.
func (e *Element) AddMouseDownListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "mousedown", listener)
}

// AddMouseEnterListener Adds a mouseenter event listener to the element.
func (e *Element) AddMouseEnterListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "mouseenter", listener)
}

// AddMouseLeaveListener Adds a mouseleave event listener to the element.
func (e *Element) AddMouseLeaveListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "mouseleave", listener)
}

// AddMouseMoveListener Adds a mousemove event listener to the element.
func (e *Element) AddMouseMoveListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "mousemove", listener)
}

// AddMouseOutListener Adds a mouseout event listener to the element.
func (e *Element) AddMouseOutListener(callback func(e *MouseEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewMouseEvent(args[0])
		callback(NewMouseEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "mouseout", listener)
}

// AddGotPointerCaptureListener Adds a gotpointercapture event listener to the element.
func (e *Element) AddGotPointerCaptureListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "gotpointercapture", listener)
}

// AddLostPointerCaptureListener Adds a lostpointercapture event listener to the element.
func (e *Element) AddLostPointerCaptureListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "lostpointercapture", listener)
}

// AddPointerCancelListener Adds a pointercancel event listener to the element.
func (e *Element) AddPointerCancelListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointercancel", listener)
}

// AddPointerDownListener Adds a pointerdown event listener to the element.
func (e *Element) AddPointerDownListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointerdown", listener)
}

// AddPointerEnterListener Adds a pointerenter event listener to the element.
func (e *Element) AddPointerEnterListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointerenter", listener)
}

// AddPointerLeaveListener Adds a pointerleave event listener to the element.
func (e *Element) AddPointerLeaveListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointerleave", listener)
}

// AddPointerMoveListener Adds a pointermove event listener to the element.
func (e *Element) AddPointerMoveListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointermove", listener)
}

// AddPointerOutListener Adds a pointerout event listener to the element.
func (e *Element) AddPointerOutListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointerout", listener)
}

// AddPointUpListener Adds a pointerup event listener to the element.
func (e *Element) AddPointerUpListener(callback func(e *PointerEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewPointerEvent(args[0])
		callback(NewPointerEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "pointerup", listener)
}

// AddTouchCancelListener Adds a touchcancel event listener to the element.
func (e *Element) AddTouchCancelListener(callback func(e *TouchEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTouchEvent(args[0])
		callback(NewTouchEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "touchcancel", listener)
}

// AddTouchEndListener Adds a touchend event listener to the element.
func (e *Element) AddTouchEndListener(callback func(e *TouchEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTouchEvent(args[0])
		callback(NewTouchEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "touchend", listener)
}

// AddTouchMoveListener Adds a touchmove event listener to the element.
func (e *Element) AddTouchMoveListener(callback func(e *TouchEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTouchEvent(args[0])
		callback(NewTouchEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "touchmove", listener)
}

// AddTouchStartListener Adds a touchstart event listener to the element.
func (e *Element) AddTouchStartListener(callback func(e *TouchEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTouchEvent(args[0])
		callback(NewTouchEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "touchstart", listener)
}

// AddTransitionCancelListener Adds a transitioncancel event listener to the element.
func (e *Element) AddTransitionCancelListener(callback func(e *TransitionEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTransitionEvent(args[0])
		callback(NewTransitionEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "transitioncancel", listener)
}

// AddTransitionEndListener Adds a transitionend event listener to the element.
func (e *Element) AddTransitionEndListener(callback func(e *TransitionEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTransitionEvent(args[0])
		callback(NewTransitionEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "transitionend", listener)
}

// AddTransitionRunListener Adds a transitionrun event listener to the element.
func (e *Element) AddTransitionRunListener(callback func(e *TransitionEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTransitionEvent(args[0])
		callback(NewTransitionEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "transitionrun", listener)
}

// AddTransitionStartListener Adds a transitionstart event listener to the element.
func (e *Element) AddTransitionStartListener(callback func(e *TransitionEvent)) {
	listener := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := NewTransitionEvent(args[0])
		callback(NewTransitionEvent(event.Js()))
		return nil
	})

	e.Js().Call("addEventListener", "transitionstart", listener)
}
