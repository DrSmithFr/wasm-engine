package element

import "go-webgl/dom/wasm"

type DeviceAcceleration struct {
	*wasm.Entity
}

// X Returns a double representing the acceleration of the device along the x axis.
func (a *DeviceAcceleration) X() float64 {
	return a.Js().Get("x").Float()
}

// SetX Sets a double representing the acceleration of the device along the x axis.
func (a *DeviceAcceleration) SetX(x float64) {
	a.Js().Set("x", x)
}

// Y Returns a double representing the acceleration of the device along the y axis.
func (a *DeviceAcceleration) Y() float64 {
	return a.Js().Get("y").Float()
}

// SetY Sets a double representing the acceleration of the device along the y axis.
func (a *DeviceAcceleration) SetY(y float64) {
	a.Js().Set("y", y)
}

// Z Returns a double representing the acceleration of the device along the z axis.
func (a *DeviceAcceleration) Z() float64 {
	return a.Js().Get("z").Float()
}

// SetZ Sets a double representing the acceleration of the device along the z axis.
func (a *DeviceAcceleration) SetZ(z float64) {
	a.Js().Set("z", z)
}
