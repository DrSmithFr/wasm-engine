package game

import "image/color"

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
    Color   color.RGBA
    Floor   float64
    Ceiling float64
    Portal  *Sector
}

func (p *Wall) Position(player Player) (float64, float64, float64, float64, float64, float64) {
    x1, y1, z1 := p.Points[0].Position()
    x2, y2, z2 := p.Points[1].Position()

    return x1, x2, y1, y2, z1, z2
}

func (p *Wall) RelativePosition(player Player) (float64, float64, float64, float64, float64, float64) {
    tx1, tz1, z1 := p.Points[0].RelativePosition(player)
    tx2, tz2, z2 := p.Points[1].RelativePosition(player)

    return tx1, tx2, tz1, tz2, z1, z2
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
    x1 := p.X - player.Position.X
    y1 := p.Y - player.Position.Y
    z1 := p.Z - player.Position.Z

    return x1, y1, z1
}
