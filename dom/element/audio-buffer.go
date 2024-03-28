package element

import (
	"go-webgl/dom/wasm"
	"syscall/js"
)

// AudioBuffer https://developer.mozilla.org/en-US/docs/Web/API/AudioBuffer
type AudioBuffer struct {
	*wasm.Entity
}

func NewAudioBuffer(raw js.Value) *AudioBuffer {
	if raw.IsNull() || raw.IsUndefined() {
		return nil
	}

	return &AudioBuffer{
		Entity: wasm.New(raw),
	}
}

// SampleRate Returns a float representing the sample rate, in samples per second, of the PCM data stored in the buffer.
func (a *AudioBuffer) SampleRate() float64 {
	return a.Js().Get("sampleRate").Float()
}

// Length Returns an integer representing the length, in sample-frames, of the PCM data stored in the buffer.
func (a *AudioBuffer) Length() int {
	return a.Js().Get("length").Int()
}

// Duration Returns a float representing the duration, in seconds, of the PCM data stored in the buffer.
func (a *AudioBuffer) Duration() float64 {
	return a.Js().Get("duration").Float()
}

// NumberOfChannels Returns an integer representing the number of discrete audio channels described by the PCM data stored in the buffer.
func (a *AudioBuffer) NumberOfChannels() int {
	return a.Js().Get("numberOfChannels").Int()
}

// GetChannelData Returns a Float32Array containing the PCM data associated with the channel, defined by the channel parameter (with 0 representing the first channel).
func (a *AudioBuffer) GetChannelData(channel int) js.Value {
	return a.Js().Call("getChannelData", channel)
}

// CopyFromChannel Copies the samples from the specified channel of the AudioBuffer to the destination array.
func (a *AudioBuffer) CopyFromChannel(destination js.Value, channelNumber int, startInChannel int) {
	a.Js().Call("copyFromChannel", destination, channelNumber, startInChannel)
}

// CopyToChannel Copies the samples to the specified channel of the AudioBuffer, from the source array.
func (a *AudioBuffer) CopyToChannel(source js.Value, channelNumber int, startInChannel int) {
	a.Js().Call("copyToChannel", source, channelNumber, startInChannel)
}
