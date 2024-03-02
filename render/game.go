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
    // transform point to player's perspective
    tx1, ty1, _ := TransformPointToPlayer(p, w.Points[0])
    tx2, ty2, _ := TransformPointToPlayer(p, w.Points[1])

    // rotate the point around the player's view
    var tz1, tz2 float64
    tx1, tz1, tx2, tz2 = RotateAroundPlayer(p, tx1, ty1, tx2, ty2)

    // draw the wall
    if tz1 > 0 || tz2 > 0 {
        //tx1, tz1, tx2, tz2 = ClipWallCoordinates(r, tx1, tz1, tx2, tz2)
        DrawWall(r, tx1, tz1, tx2, tz2)
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
    const clippingOffset = 0.

    screenW, screenH := r.Size()

    halfW := float64(screenW / 2)
    halfH := float64(screenH / 2)

    // get player's view plane coordinates
    px1 := zeroPos
    py1 := zeroPos
    px2 := halfW / 2
    py2 := halfH / 2

    // if line crosses the player's view plane, clip it
    ix1, iz1 := Intersect(tx1, tz1, tx2, tz2, -px1, py1, px2, py2)
    ix2, iz2 := Intersect(tx1, tz1, tx2, tz2, px1, py1, px2, py2)

    if iz1 <= 0 {
        if iz1 > 0 {
            tx1 = ix1
            tz1 = iz1
        } else {
            tx1 = ix2
            tz1 = iz2
        }
    }

    if iz2 <= 0 {
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
    return x1*y2 - x2*y1
}

func DrawWall(r Renderer, tx1, tz1, tx2, tz2 float64) {
    screenW, screenH := r.Size()

    halfW := float64(screenW / 2)
    halfH := float64(screenH / 2)

    x1 := -tx1 * 16 / tz1
    y1a := -halfW / tz1
    y1b := halfW / tz1

    x2 := -tx2 * 16 / tz2
    y2a := -halfW / tz2
    y2b := halfW / tz2

    r.DrawLine(halfW+x1, halfH+y1a, halfW+x2, halfH+y2a, 1)
    r.DrawLine(halfW+x2, halfH+y2a, halfW+x2, halfH+y2b, 1)
    r.DrawLine(halfW+x2, halfH+y2b, halfW+x1, halfH+y1b, 1)
    r.DrawLine(halfW+x1, halfH+y1b, halfW+x1, halfH+y1a, 1)
}
