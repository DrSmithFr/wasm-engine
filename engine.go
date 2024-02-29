package main

import (
    "go-webgl/browser"
    "go-webgl/controller"
    "go-webgl/game"
    "go-webgl/render"
    "image/color"
    "math"
)

var DOM browser.DOM
var controls controller.Interface
var gs *game.GameState

var width int
var height int

func main() {
    // loading DOM to memory
    DOM = browser.LoadDOM()
    width, height = DOM.GetScreenSize()

    engine := render.NewDirectCtx(width, height, 0, 0)
    controls = controller.NewKeyboardOnly()

    // setting up everything
    engine.Init(DOM)
    controls.Init(DOM)

    // creating game state
    var err error
    if gs, err = game.NewGameState(width, height); err != nil {
        panic(err)
    }

    // start the game loop
    engine.Start(30, GameLoop)

    // avoid WebAssembly to exit the program
    run := make(chan bool)
    <-run
}

func GameLoop(r render.Renderer) bool {
    // render default color
    r.Clear()

    rt := game.NewRayTracer(gs)
    rays := rt.ComputeRays()

    renderGameView(r, rays)
    renderMiniMap(r, rays)

    handleMove()

    return true
}

func renderMiniMap(r render.Renderer, rays []game.Ray) {
    renderLevel(r)
    renderRayCasting(r, rays)
    renderPlayer(r)
}

func renderRayCasting(r render.Renderer, rays []game.Ray) {
    r.SetColor(color.RGBA{0xff, 0x00, 0x00, 0xff})

    for _, ray := range rays {
        r.DrawLine(
            ray.Origin.X,
            ray.Origin.Y,
            ray.Impact.X,
            ray.Impact.Y,
            1,
        )
    }
}

func renderGameView(r render.Renderer, rays []game.Ray) {
    const screenHeight = 320
    const lineWidth = 8

    screenWidth := len(rays) * lineWidth

    up := &game.Upscale{
        Source: game.Resolution{screenWidth, screenHeight},
        Target: game.Resolution{width, height},
    }

    w, h := r.Size()
    renderSky(r, w, h, up)
    renderGround(r, w, h, up)

    for rayN, ray := range rays {
        angle := game.FixAngle(ray.Origin.Angle + 30 - rayN*60/len(rays))
        ca := game.DegToRad(float64(angle))

        // fix fisheye
        distT := ray.Distance * math.Cos(ca)

        lineH := float64(gs.GetMapSize()*screenHeight) / distT
        if lineH > screenHeight {
            lineH = screenHeight
        }

        lineOffset := (screenHeight / 2) - lineH/2

        var c color.RGBA

        if ray.Impact.Type == game.Horizontal {
            c = color.RGBA{0xcc, 0xcc, 0xcc, 0xff}
        } else {
            c = color.RGBA{0xff, 0xff, 0xff, 0xff}
        }

        if ray.Impact.CellType == game.Door {
            c.R = 0
            c.G = 0
        }

        if ray.Impact.CellType == game.Window {
            c.R = 0
        }

        if ray.Impact.CellType == game.Checkerboard {
            c.R = 0
            c.B = 0
        }

        r.SetColor(c)
        r.DrawRect(
            up.ScaleWidth(float64(rayN*lineWidth)),
            up.ScaleHeight(lineOffset),
            up.ScaleWidth(float64(rayN*lineWidth)+lineWidth),
            up.ScaleHeight(lineH+lineOffset),
        )
    }
}

func renderSky(r render.Renderer, screenWidth, screenHeight int, up *game.Upscale) {
    skyColor := color.RGBA{0x00, 0x00, 0x99, 0xff}

    r.SetColor(skyColor)
    r.DrawRect(0, 0, float64(screenWidth), float64(screenHeight/2))
}

func renderGround(r render.Renderer, screenWidth, screenHeight int, up *game.Upscale) {
    groundColor := color.RGBA{0x00, 0x99, 0x99, 0xff}

    r.SetColor(groundColor)
    r.DrawRect(0, float64(screenHeight/2), float64(screenWidth), float64(screenHeight))
}

func handleMove() {
    actions := controls.GetState()

    if actions.Action {
        gs.Action()
    }

    if actions.Up {
        gs.MoveUp()
    } else if actions.Down {
        gs.MoveDown()
    }

    if actions.TurnRight {
        gs.TurnRight()
    } else if actions.TurnLeft {
        gs.TurnLeft()
    }
}

func renderLevel(r render.Renderer) {
    level := gs.GetLevel()
    blockSize := gs.GetBlockSize()
    mapSize := gs.GetMapSize()

    // minimap background
    r.SetColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
    r.DrawRect(0, 0, float64(mapSize*blockSize), float64(mapSize*blockSize))

    // draw walls
    r.SetColor(color.RGBA{0xff, 0xff, 0xff, 0xcc})

    for y := 0; y < mapSize; y++ {
        for x := 0; x < mapSize; x++ {
            if level.Walls[x+y*mapSize] == 0 {
                // avoid useless rendering
                continue
            }

            c := color.RGBA{0xff, 0xff, 0xff, 0xcc}

            if level.Walls[x+y*mapSize] == game.Door {
                c.R = 0
                c.G = 0
            }

            if level.Walls[x+y*mapSize] == game.Window {
                c.R = 0
            }

            if level.Walls[x+y*mapSize] == game.Checkerboard {
                c.R = 0
                c.B = 0
            }

            r.SetColor(c)
            r.DrawRect(
                float64(x*blockSize),
                float64(y*blockSize),
                float64(x*blockSize+blockSize),
                float64(y*blockSize+blockSize),
            )
        }
    }
}

func renderPlayer(r render.Renderer) {
    // draw player on screen
    r.SetColor(color.RGBA{0xff, 0xff, 0x00, 0xcc})

    // draw player on screen
    playerX, playerY, playerDeltaX, playerDeltaY := gs.GetPlayerPosition()
    r.DrawCircle(playerX, playerY, 5)

    // draw player direction
    r.DrawLine(
        playerX,
        playerY,
        playerX+playerDeltaX*5,
        playerY+playerDeltaY*5,
        1,
    )
}
