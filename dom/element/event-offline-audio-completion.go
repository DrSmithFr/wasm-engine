package element

// OfflineAudioCompletionEvent https://developer.mozilla.org/en-US/docs/Web/API/OfflineAudioCompletionEvent
type OfflineAudioCompletionEvent struct {
	*Event
}

// RenderedBuffer Returns an AudioBuffer containing the rendered audio data.
func (e *OfflineAudioCompletionEvent) RenderedBuffer() *AudioBuffer {
	return NewAudioBuffer(e.Js().Get("renderedBuffer"))
}
