package game

import "math"

type Angle struct {
    Alfa int8
}

func FixAngle(alpha int) int {
    if alpha < 0 {
        return alpha + 360
    }

    if alpha > 359 {
        return alpha - 360
    }

    return alpha
}

func DegToRad(alpha int) float64 {
    return float64(alpha) * math.Pi / 180.
}
