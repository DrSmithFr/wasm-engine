package element

import "go-webgl/dom/wasm"

// DeviceRotationRate
type DeviceRotationRate struct {
	*wasm.Entity
}

// Alpha Returns a double representing the rotation of the device around the z axis.
func (r *DeviceRotationRate) Alpha() float64 {
	return r.Js().Get("alpha").Float()
}

// SetAlpha Sets a double representing the rotation of the device around the z axis.
func (r *DeviceRotationRate) SetAlpha(alpha float64) {
	r.Js().Set("alpha", alpha)
}

// Beta Returns a double representing the rotation of the device around the x axis.
func (r *DeviceRotationRate) Beta() float64 {
	return r.Js().Get("beta").Float()
}

// SetBeta Sets a double representing the rotation of the device around the x axis.
func (r *DeviceRotationRate) SetBeta(beta float64) {
	r.Js().Set("beta", beta)
}

// Gamma Returns a double representing the rotation of the device around the y axis.
func (r *DeviceRotationRate) Gamma() float64 {
	return r.Js().Get("gamma").Float()
}

// SetGamma Sets a double representing the rotation of the device around the y axis.
func (r *DeviceRotationRate) SetGamma(gamma float64) {
	r.Js().Set("gamma", gamma)
}
