package game

import (
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
    "image/color"
    "math"
)

type RayTracer struct {
    gameState *GameState
}

func NewRayTracer(gs *GameState) *RayTracer {
    return &RayTracer{gs}
}

func DegToRad(deg float64) float64 {
    return deg * (math.Pi / 180.)
}

func FixAngle(angle int) int {
    if angle > 359 {
        return angle - 360
    }

    if angle < 0 {
        return angle + 360
    }

    return angle
}

func (rt *RayTracer) ComputeRays() []Ray {
    const maxDepth = 8

    gs := rt.gameState

    var rayAngle int
    var rayX, rayY float64
    var rayTargetX, rayTargetY float64
    var mapX, mapY int
    var dof int
    var distT float64
    var TargetTypeT Cell

    blocSize := gs.GetBlockSize()
    playerX, playerY, _, _ := gs.GetPlayerPosition()

    const FieldOfViewsAngle = 60
    startAngle := FixAngle(gs.GetPlayer().Position.Angle - 30)

    var rays []Ray

    for rayN := 0; rayN <= FieldOfViewsAngle; rayN++ {
        rayAngleDeg := FixAngle(startAngle + rayAngle)
        rayAngleRad := DegToRad(float64(rayAngleDeg))

        // check Horizontal
        dof = 0
        aTan := -1 / math.Tan(rayAngleRad)

        distH := 1000000.0
        TargetTypeH := EmptyCell
        hx := playerX
        hy := playerY

        if rayAngleDeg > 180 {
            // looking up
            rayY = math.Trunc(playerY/float64(blocSize))*float64(blocSize) - 1
            rayX = (playerY-rayY)*aTan + playerX

            rayTargetY = -float64(blocSize)
            rayTargetX = -rayTargetY * aTan
        } else if rayAngleDeg < 180 {
            // looking down
            rayY = math.Trunc(playerY/float64(blocSize))*float64(blocSize) + float64(blocSize)
            rayX = (playerY-rayY)*aTan + playerX

            rayTargetY = float64(blocSize)
            rayTargetX = -rayTargetY * aTan
        }

        if rayAngleDeg == 0 || rayAngleDeg == 180 {
            rayX = playerX
            rayY = playerY
            dof = maxDepth
        }

        for ; dof < maxDepth; {
            cell := gs.GetMapValueAt(rayX, rayY)

            // hit wall
            if cell > EmptyCell {
                TargetTypeH = cell
                dof = maxDepth
                hx = rayX
                hy = rayY
                distH = dist(playerX, playerY, hx, hy, rayAngleRad)
            } else {
                rayX += rayTargetX
                rayY += rayTargetY
                dof++
            }
        }

        // check Vertical
        dof = 0
        nTan := -math.Tan(rayAngleRad)

        distV := 1000000.0
        TargetTypeV := EmptyCell
        vx := playerX
        vy := playerY

        if rayAngleDeg > 90 && rayAngleDeg < 270 {
            // looking left
            rayX = math.Trunc(playerX/float64(blocSize))*float64(blocSize) - 1
            rayY = (playerX-rayX)*nTan + playerY

            rayTargetX = -float64(blocSize)
            rayTargetY = -rayTargetX * nTan
        } else if rayAngleRad < 90 || rayAngleRad > 270 {
            // looking right
            rayX = math.Trunc(playerX/float64(blocSize))*float64(blocSize) + float64(blocSize)
            rayY = (playerX-rayX)*nTan + playerY

            rayTargetX = float64(blocSize)
            rayTargetY = -rayTargetX * nTan
        }

        if rayAngleDeg == 0 || rayAngleDeg == 180 {
            rayX = playerX
            rayY = playerY
            dof = maxDepth
        }

        for ; dof < maxDepth; {
            cell := gs.GetMapValueAt(rayX, rayY)

            // hit wall
            if cell > EmptyCell {
                TargetTypeV = cell
                vx = rayX
                vy = rayY
                distV = dist(playerX, playerY, vx, vy, rayAngleRad)
                dof = maxDepth
            } else {
                rayX += rayTargetX
                rayY += rayTargetY
                dof++
            }
        }

        var impactType ImpactType

        // vertical wall
        if distV < distH {
            rayX = vx
            rayY = vy
            distT = distV
            TargetTypeT = TargetTypeV
            impactType = Vertical
        }

        // horizontal wall
        if distH < distV {
            rayX = hx
            rayY = hy
            distT = distH
            TargetTypeT = TargetTypeH
            impactType = Horizontal
        }

        ray := Ray{
            Origin: Point{
                X:     playerX,
                Y:     playerY,
                Angle: rayAngle,
            },
            Impact: Impact{
                X:        rayX,
                Y:        rayY,
                CellX:    mapX,
                CellY:    mapY,
                CellType: TargetTypeT,
                Type:     impactType,
            },
            Distance: distT,
        }

        rays = append(rays, ray)
        rayAngle++
    }

    return rays
}

func (rt *RayTracer) DrawRay(
    gc *draw2dimg.GraphicContext,
    color color.RGBA,
    ray Ray,
) {
    gc.SetFillColor(color)
    gc.SetStrokeColor(color)

    drawLine(gc, ray.Origin.X, ray.Origin.Y, ray.Impact.X, ray.Impact.Y)
    drawPoint(gc, ray.Origin.X, ray.Origin.Y, 5)

    rt.drawCellHighlight(gc, ray.Impact.CellX, ray.Impact.CellY, color)
}

func (rt *RayTracer) drawCellHighlight(gc *draw2dimg.GraphicContext, x int, y int, rgba color.RGBA) {
    blockSize := rt.gameState.GetBlockSize()

    gc.SetFillColor(rgba)
    gc.SetStrokeColor(rgba)

    draw2dkit.Rectangle(
        gc,
        float64(x*blockSize),
        float64(y*blockSize),
        float64(x*blockSize+blockSize),
        float64(y*blockSize+blockSize),
    )
    gc.FillStroke()

}

func dist(aX, aY, bX, bY, angle float64) float64 {
    return math.Sqrt((bX-aX)*(bX-aX) + (bY-aY)*(bY-aY))
}

func drawPoint(gc *draw2dimg.GraphicContext, x, y, radius float64) {
    draw2dkit.Circle(gc, x, y, radius)
    gc.FillStroke()
}

func drawLine(gc *draw2dimg.GraphicContext, x1, y1, x2, y2 float64) {
    gc.BeginPath()
    gc.MoveTo(x1, y1)
    gc.LineTo(x2, y2)
    gc.Close()
    gc.FillStroke()
}
