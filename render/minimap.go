package render

import (
    "go-webgl/game"
    "image/color"
)

func RenderMinimap(r Renderer, p game.Player, m game.Map) {
    screenW, screenH := r.Size()
    r.SetColor(color.RGBA{255, 255, 255, 255})

    // render player position
    r.DrawCircle(
        p.Position.X+float64(screenW/2),
        p.Position.Y+float64(screenH/2),
        5,
    )

    // render player direction
    x1 := p.Position.X
    y1 := p.Position.Y

    // draw a line in the direction the player is facing
    dx, dy := p.Position.GetDelta()
    x2 := x1 + dx*game.MovementMomentum
    y2 := y1 + dy*game.MovementMomentum

    r.DrawLine(
        x1+float64(screenW/2),
        y1+float64(screenH/2),
        x2+float64(screenW/2),
        y2+float64(screenH/2),
        1,
    )

    for _, s := range m.Sectors {
        for _, wall := range s.Walls {
            x1, y1, _ := wall.Points[0].Position()
            x2, y2, _ := wall.Points[1].Position()

            r.DrawLine(
                x1+float64(screenW/2),
                y1+float64(screenH/2),
                x2+float64(screenW/2),
                y2+float64(screenH/2),
                1,
            )
        }
    }
}
