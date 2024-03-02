package render

import (
    "go-webgl/game"
    "math"
)

func RenderSector(r Renderer, p game.Player, s game.Sector) {
    for _, wall := range s.Walls {
        RenderWall(r, p, wall)
    }
}

func RenderWall(r Renderer, p game.Player, w game.Wall) {
    screenW, screenH := r.Size()

    halfW := float64(screenW / 2)
    halfH := float64(screenH / 2)

    // transform point to player's perspective
    tx1, ty1, _ := w.Points[0].RelativePosition(p)
    tx2, ty2, _ := w.Points[1].RelativePosition(p)

    // rotate the point around the player's view
    angle := p.Position.Rad()
    tz1 := tx1*math.Cos(angle) + ty1*math.Sin(angle)
    tz2 := tx2*math.Cos(angle) + ty2*math.Sin(angle)
    tx1 = tx1*math.Sin(angle) - ty1*math.Cos(angle)
    tx2 = tx2*math.Sin(angle) - ty2*math.Cos(angle)

    // apply perspective
    x1 := -tx1 * 16 / tz1
    y1a := -halfW / tz1
    y1b := halfW / tz1

    x2 := -tx2 * 16 / tz2
    y2a := -halfW / tz2
    y2b := halfW / tz2

    // draw the wall
    r.DrawLine(halfW+x1, halfH+y1a, halfW+x2, halfH+y2a, 1)
    r.DrawLine(halfW+x2, halfH+y2a, halfW+x2, halfH+y2b, 1)
    r.DrawLine(halfW+x2, halfH+y2b, halfW+x1, halfH+y1b, 1)
    r.DrawLine(halfW+x1, halfH+y1b, halfW+x1, halfH+y1a, 1)
}
