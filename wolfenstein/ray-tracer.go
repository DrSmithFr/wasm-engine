package wolfenstein

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

func (rt *RayTracer) ComputeRays() []Ray {
    const precision = 2
    const maxDepth = 8
    gs := rt.gameState

    var rayX, rayY, rayAngle float64
    var rayTargetX, rayTargetY float64
    var mapX, mapY, mapIndex int
    var dof int
    var distT float64
    var TargetTypeT Cell

    level := gs.GetLevel()
    blocSize := gs.GetBlockSize()
    mapSize := gs.GetMapSize()
    playerX, playerY, _, _ := gs.GetPlayerPosition()

    oneRadian := 0.0174533 / precision
    rayAngle = gs.GetPlayerAngle()

    const FieldOfViewsAngle = 60 * precision
    rayAngle -= oneRadian * FieldOfViewsAngle / 2

    if rayAngle < 0 {
        rayAngle += 2 * math.Pi
    } else if rayAngle > 2*math.Pi {
        rayAngle -= 2 * math.Pi
    }

    var rays []Ray

    for rayN := 0; rayN <= FieldOfViewsAngle; rayN++ {
        // check Horizontal
        dof = 0
        aTan := -1 / math.Tan(rayAngle)

        distH := 1000000.0
        TargetTypeH := EmptyCell
        hx := playerX
        hy := playerY

        if rayAngle > math.Pi {
            // looking up
            rayY = math.Trunc(playerY/float64(blocSize))*float64(blocSize) - 1
            rayX = (playerY-rayY)*aTan + playerX

            rayTargetY = -float64(blocSize)
            rayTargetX = -rayTargetY * aTan
        } else if rayAngle < math.Pi {
            // looking down
            rayY = math.Trunc(playerY/float64(blocSize))*float64(blocSize) + float64(blocSize)
            rayX = (playerY-rayY)*aTan + playerX

            rayTargetY = float64(blocSize)
            rayTargetX = -rayTargetY * aTan
        }

        if rayAngle == 0 || rayAngle == math.Pi {
            rayX = playerX
            rayY = playerY
            dof = maxDepth
        }

        for ; dof < maxDepth; {
            mapX = int(math.Trunc(rayX / float64(blocSize)))
            mapY = int(math.Trunc((rayY) / float64(blocSize)))

            mapIndex = mapY*mapSize + mapX

            // hit wall
            if mapIndex > 0 && mapIndex < mapSize*mapSize && level.Walls[mapIndex] > EmptyCell {
                TargetTypeH = level.Walls[mapIndex]
                dof = maxDepth
                hx = rayX
                hy = rayY
                distH = dist(playerX, playerY, hx, hy, rayAngle)
            } else {
                rayX += rayTargetX
                rayY += rayTargetY
                dof++
            }
        }

        // check Vertical
        dof = 0
        nTan := -math.Tan(rayAngle)
        P2 := math.Pi / 2
        P3 := 3 * P2

        distV := 1000000.0
        TargetTypeV := EmptyCell
        vx := playerX
        vy := playerY

        if rayAngle > P2 && rayAngle < P3 {
            // looking left
            rayX = math.Trunc(playerX/float64(blocSize))*float64(blocSize) - 1
            rayY = (playerX-rayX)*nTan + playerY

            rayTargetX = -float64(blocSize)
            rayTargetY = -rayTargetX * nTan
        } else if rayAngle < P2 || rayAngle > P3 {
            // looking right
            rayX = math.Trunc(playerX/float64(blocSize))*float64(blocSize) + float64(blocSize)
            rayY = (playerX-rayX)*nTan + playerY

            rayTargetX = float64(blocSize)
            rayTargetY = -rayTargetX * nTan
        }

        if rayAngle == 0 || rayAngle == math.Pi {
            rayX = playerX
            rayY = playerY
            dof = maxDepth
        }

        for ; dof < maxDepth; {
            mapX = int(math.Trunc(rayX / float64(blocSize)))
            mapY = int(math.Trunc((rayY) / float64(blocSize)))

            mapIndex = mapY*mapSize + mapX

            // hit wall
            if mapIndex > 0 && mapIndex < mapSize*mapSize && level.Walls[mapIndex] > EmptyCell {
                TargetTypeV = level.Walls[mapIndex]
                vx = rayX
                vy = rayY
                distV = dist(playerX, playerY, vx, vy, rayAngle)
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

        // updating Angle for next ray
        rayAngle += oneRadian

        if rayAngle < 0 {
            rayAngle += 2 * math.Pi
        } else if rayAngle > 2*math.Pi {
            rayAngle -= 2 * math.Pi
        }
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
