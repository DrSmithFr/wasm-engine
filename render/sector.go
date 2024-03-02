package render

import (
    "go-webgl/game"
)

func RenderSector(r Renderer, p game.Player, s game.Sector) {
    for _, wall := range s.Walls {
        RenderWall(r, p, wall)
    }
}

func RenderWall(r Renderer, p game.Player, w game.Wall) {
    screenW, screenH := r.Size()

    x1, y1, z1 := w.Points[0].RelativePosition(p)
    x2, y2, z2 := w.Points[1].RelativePosition(p)

    // set origin to the center of the screen
    x1 += float64(screenW) / 2
    y1 += float64(screenH) / 2
    x2 += float64(screenW) / 2
    y2 += float64(screenH) / 2

    // render wall
    r.DrawLine(x1, y1, x2, y2, 1)
    r.DrawLine(x2, y2, x2, z2, 1)
    r.DrawLine(x2, z2, x1, z1, 1)
    r.DrawLine(x1, z1, x1, y1, 1)
}
