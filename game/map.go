package game

import (
    "math"
)

type Collision int

const (
    None Collision = iota
    Window
    Door
)

type Map struct {
    Sectors    []Sector
    Spawn      Point3D
    SpawnAngle int
}

type Sector struct {
    Walls []Wall
}

type Wall struct {
    Points  [2]Point3D
    Floor   float64
    Ceiling float64
    Portal  *Sector
}

func (p *Wall) RelativePosition(player Player) (float64, float64, float64, float64, float64, float64) {
    x1, y1, z1 := p.Points[0].RelativePosition(player)
    x2, y2, z2 := p.Points[1].RelativePosition(player)

    return x1, y1, z1, x2, y2, z2
}

type Point3D struct {
    X float64
    Y float64
    Z float64
}

func (p *Point3D) Position() (float64, float64, float64) {
    return p.X, p.Y, p.Z
}

func (p *Point3D) RelativePosition(player Player) (float64, float64, float64) {
    // transform point to player's perspective
    x := p.X - player.Position.X
    y := p.Y - player.Position.Y

    // rotate point to player's perspective
    z := x*math.Cos(player.Position.Rad()) + y*math.Sin(player.Position.Rad())
    x = x*math.Sin(player.Position.Rad()) - y*math.Cos(player.Position.Rad())

    return x, y, z
}
