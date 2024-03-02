package render

import (
    "go-webgl/game"
    "image/color"
)

func RenderMinimap(r Renderer, p game.Player, m game.Map) {
    r.Clear()
    r.SetColor(color.RGBA{255, 255, 255, 255})

    screenW, screenH := r.Size()

    halfW := float64(screenW / 2)
    halfH := float64(screenH / 2)

    // render player position
    r.DrawCircle(
        halfW,
        halfH,
        5,
    )

    // draw a line in the direction the player is facing
    const drawLineLength = 20
    r.DrawLine(halfW, halfH, halfW, halfH-drawLineLength, 1)

    // render sectors
    for _, s := range m.Sectors {
        for _, w := range s.Walls {
            x1, y1, _ := TransformPointToPlayer(p, w.Points[0])
            x2, y2, _ := TransformPointToPlayer(p, w.Points[1])

            tx1, tz1, tx2, tz2 := RotateAroundPlayer(p, x1, y1, x2, y2)

            r.DrawLine(
                halfW-tx1,
                halfH-tz1,
                halfW-tx2,
                halfH-tz2,
                1,
            )
        }
    }

    r.Flush()
}
