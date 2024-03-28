package element

import "go-webgl/dom/wasm"

// DeviceMotionEvent https://developer.mozilla.org/en-US/docs/Web/API/DeviceMotionEvent
type DeviceMotionEvent struct {
	*Event
}

// Acceleration Returns an DeviceAcceleration object representing the acceleration of the device.
func (e *DeviceMotionEvent) Acceleration() *DeviceAcceleration {
	return &DeviceAcceleration{
		Entity: wasm.New(e.Js().Get("acceleration")),
	}
}

// AccelerationIncludingGravity Returns an DeviceAcceleration object representing the acceleration of the device including gravity.
func (e *DeviceMotionEvent) AccelerationIncludingGravity() *DeviceAcceleration {
	return &DeviceAcceleration{
		Entity: wasm.New(e.Js().Get("accelerationIncludingGravity")),
	}
}

// RotationRate Returns a DeviceRotationRate object representing the rate of change of the device's orientation.
func (e *DeviceMotionEvent) RotationRate() *DeviceRotationRate {
	return &DeviceRotationRate{
		Entity: wasm.New(e.Js().Get("rotationRate")),
	}

}

// Interval Returns a double representing the interval of the event.
func (e *DeviceMotionEvent) Interval() float64 {
	return e.Js().Get("interval").Float()
}
