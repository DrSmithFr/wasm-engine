package render

import (
    "go-webgl/game"
    "image/color"
    "math"
)

func RenderSector(r Renderer, p game.Player, s game.Sector) {
    for _, wall := range s.Walls {
        RenderWall(r, p, wall)
    }
}

func RenderWall(r Renderer, p game.Player, w game.Wall) {
    // transform point to player's perspective
    x1, y1, _ := TransformPointToPlayer(p, w.Points[0])
    x2, y2, _ := TransformPointToPlayer(p, w.Points[1])

    // rotate the point around the player's view
    var z1, z2 float64
    x1, z1, x2, z2 = RotateAroundPlayer(p, x1, y1, x2, y2)

    // draw the wall
    if z1 > 0 || z2 > 0 {
        x1, z1, x2, z2 = ClipWallCoordinates(r, x1, z1, x2, z2)
        DrawWall(r, w, x1, z1, x2, z2)
    }
}

func TransformPointToPlayer(p game.Player, pts game.Point3D) (float64, float64, float64) {
    // transform point to player's perspective
    tx := pts.X - p.Position.X
    ty := pts.Y - p.Position.Y
    tz := pts.Z - p.Position.Z

    return tx, ty, tz
}

func RotateAroundPlayer(p game.Player, tx1, ty1, tx2, ty2 float64) (float64, float64, float64, float64) {
    angle := p.Position.Rad()

    // rotate the point around the player's view
    tz1 := tx1*math.Cos(angle) + ty1*math.Sin(angle)
    tz2 := tx2*math.Cos(angle) + ty2*math.Sin(angle)
    tx1 = tx1*math.Sin(angle) - ty1*math.Cos(angle)
    tx2 = tx2*math.Sin(angle) - ty2*math.Cos(angle)

    return tx1, tz1, tx2, tz2
}

func ClipWallCoordinates(r Renderer, tx1, tz1, tx2, tz2 float64) (float64, float64, float64, float64) {
    const zeroPos = 0.0001

    screenW, _ := r.Size()
    halfW := float64(screenW) / 2.0

    // if line crosses the player's view plane, clip it
    ix1, iz1 := Intersect(tx1, tz1, tx2, tz2, zeroPos, zeroPos, -halfW, 0)
    ix2, iz2 := Intersect(tx1, tz1, tx2, tz2, zeroPos, zeroPos, halfW, 0)

    if tz1 <= 0 {
        if iz1 > 0 {
            tx1 = ix1
            tz1 = iz1
        } else {
            tx1 = ix2
            tz1 = iz2
        }
    }

    if tz2 <= 0 {
        if iz1 > 0 {
            tx2 = ix1
            tz2 = iz1
        } else {
            tx2 = ix2
            tz2 = iz2
        }
    }

    return tx1, tz1, tx2, tz2
}

func Intersect(x1, y1, x2, y2, x3, y3, x4, y4 float64) (float64, float64) {
    x := CrossProduct(x1, y1, x2, y2)
    y := CrossProduct(x3, y3, x4, y4)

    det := CrossProduct(x1-x2, y1-y2, x3-x4, y3-y4)

    x = CrossProduct(x, x1-x2, y, x3-x4) / det
    y = CrossProduct(x, y1-y2, y, y3-y4) / det

    return x, y
}

func CrossProduct(x1, y1, x2, y2 float64) float64 {
    return x1*y2 - y1*x2
}

func DrawWall(r Renderer, w game.Wall, tx1, tz1, tx2, tz2 float64) {
    const fieldOfView = 60
    screenW, screenH := r.Size()

    halfW := float64(screenW / 2)
    halfH := float64(screenH / 2)

    // perspective projection top of the wall
    x1 := -tx1 * fieldOfView / tz1
    y1a := -halfW / tz1
    y1b := halfW / tz1

    // perspective projection bottom of the wall
    x2 := -tx2 * fieldOfView / tz2
    y2a := -halfW / tz2
    y2b := halfW / tz2

    // Mesh rendering
    for x := x1; x <= x2; x++ {
        if x < -halfW {
            x = -halfW
            continue
        }

        if x > halfW {
            break
        }

        divider := x2 - x1

        if divider == 0 {
            divider = 1
        }

        ya := y1a + (x-x1)*(y2a-y1a)/divider
        yb := y1b + (x-x1)*(y2b-y1b)/divider

        // draw ceiling
        r.SetColor(color.RGBA{150, 150, 150, 100})
        r.DrawLine(halfW+x, 0, halfW+x, halfH+ya, 1)

        // draw wall
        r.SetColor(w.Color)
        r.DrawLine(halfW+x, halfH+ya, halfW+x, halfH+yb, 1)

        // draw floor
        r.SetColor(color.RGBA{0, 0, 255, 100})
        r.DrawLine(halfW+x, halfH+yb, halfW+x, float64(screenH), 1)
    }

    // draw wall edges
    r.SetColor(color.RGBA{255, 0, 255, 100})
    r.DrawLine(halfW+x1, halfH+y1a, halfW+x1, halfH+y1b, 1)
    r.DrawLine(halfW+x2, halfH+y2a, halfW+x2, halfH+y2b, 1)
}
