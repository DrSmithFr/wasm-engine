package render

import (
    "go-webgl/game"
    "image/color"
)

func RenderMap(r Renderer, p game.Player, m game.Map) {
    r.Clear()

    screenW, screenH := r.Size()

    // render grid lines
    r.SetColor(color.RGBA{60, 60, 60, 255})

    for x := 0; x < screenW; x += 10 {
        r.DrawLine(float64(x), 0, float64(x), float64(screenH), 1)
    }

    for y := 0; y < screenH; y += 10 {
        r.DrawLine(0, float64(y), float64(screenW), float64(y), 1)
    }

    // draw grid axis
    r.SetColor(color.RGBA{100, 100, 100, 255})
    r.DrawLine(float64(screenW/2), 0, float64(screenW/2), float64(screenH), 1)
    r.DrawLine(0, float64(screenH/2), float64(screenW), float64(screenH/2), 1)

    px, py := p.Position.X, p.Position.Y

    // render player position
    r.SetColor(color.RGBA{255, 255, 255, 255})
    r.DrawCircle(
        px+float64(screenW/2),
        py+float64(screenH/2),
        5,
    )

    // draw a line in the direction the player is facing
    dx, dy := p.Position.GetDelta()

    r.DrawLine(
        px+float64(screenW/2),
        py+float64(screenH/2),
        px+dx*game.MovementMomentum+float64(screenW/2),
        py+dy*game.MovementMomentum+float64(screenH/2),
        1,
    )

    // render sectors
    for _, s := range m.Sectors {
        for _, wall := range s.Walls {
            x1, x2, y1, y2, _, _ := wall.Position(p)

            r.SetColor(color.RGBA{255, 255, 255, 255})

            // set color to red if wall is a portal
            if wall.Portal != nil {
                r.SetColor(color.RGBA{255, 0, 0, 255})
            }

            // draw the wall
            r.DrawLine(
                x1+float64(screenW/2),
                y1+float64(screenH/2),
                x2+float64(screenW/2),
                y2+float64(screenH/2),
                1,
            )

            // draw wall points
            r.DrawCircle(x1+float64(screenW/2), y1+float64(screenH/2), 2)
            r.DrawCircle(x2+float64(screenW/2), y2+float64(screenH/2), 2)
        }
    }

    r.Flush()
}
