package render

import (
    "go-webgl/game"
    "image/color"
)

func RenderMinimap(r Renderer, p game.Player, m game.Map) {
    r.Clear()

    screenW, screenH := r.Size()

    // render player position
    r.SetColor(color.RGBA{255, 255, 255, 255})
    r.DrawCircle(
        float64(screenW/2),
        float64(screenH/2),
        5,
    )

    // draw a line in the direction the player is facing
    dx, dy := p.Position.GetDelta()

    r.DrawLine(
        float64(screenW/2),
        float64(screenH/2),
        float64(screenW/2)+dx*game.MovementMomentum,
        float64(screenH/2)+dy*game.MovementMomentum,
        1,
    )

    // render player direction
    px := p.Position.X
    py := p.Position.Y

    for _, s := range m.Sectors {
        for _, wall := range s.Walls {
            x1, y1, _ := wall.Points[0].Position()
            x2, y2, _ := wall.Points[1].Position()

            // translate the wall to the player position
            x1 -= px
            y1 -= py
            x2 -= px
            y2 -= py

            r.DrawLine(
                x1+float64(screenW/2),
                y1+float64(screenH/2),
                x2+float64(screenW/2),
                y2+float64(screenH/2),
                1,
            )
        }
    }

    r.Flush()
}
