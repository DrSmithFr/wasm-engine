package element

// DeviceOrientationEvent https://developer.mozilla.org/en-US/docs/Web/API/DeviceOrientationEvent
type DeviceOrientationEvent struct {
	*Event
}

// Absolute Returns a Boolean that indicates whether or not the device is providing orientation data absolutely.
func (e *DeviceOrientationEvent) Absolute() bool {
	return e.Js().Get("absolute").Bool()
}

// Alpha Returns a double representing the motion of the device around the z axis.
func (e *DeviceOrientationEvent) Alpha() float64 {
	return e.Js().Get("alpha").Float()
}

// Beta Returns a double representing the motion of the device around the x axis.
func (e *DeviceOrientationEvent) Beta() float64 {
	return e.Js().Get("beta").Float()
}

// Gamma Returns a double representing the motion of the device around the y axis.
func (e *DeviceOrientationEvent) Gamma() float64 {
	return e.Js().Get("gamma").Float()
}
