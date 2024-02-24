package main

import (
    "fmt"
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
    "go-webgl/browser"
    "go-webgl/canvas"
    "go-webgl/controller"
    "go-webgl/wolfenstein"
    "image/color"
    "math"
    "runtime"
    "syscall/js"
)

var DOM *browser.DOM
var ctl controller.Interface
var cvs *canvas.Canvas2d
var gs *wolfenstein.GameState

var width float64
var height float64

func main() {
    // loading DOM to memory
    DOM = browser.LoadDOM()
    ctl = controller.NewKeyboardOnly()

    // setting up everything
    ctl.Init(*DOM)
    bindEvents(*DOM)

    // create canvas
    cvs, _ = canvas.NewCanvas2d(false)
    cvs.Create(
        js.Global().Get("innerWidth").Int(),
        js.Global().Get("innerHeight").Int(),
    )

    DOM.Log(fmt.Sprintf("number of thread: %d", runtime.NumCPU()))

    // create gameState
    gs, _ = wolfenstein.NewGameState(cvs.Width(), cvs.Height())

    height = float64(cvs.Height())
    width = float64(cvs.Width())

    // starting rendering
    cvs.Start(30, GameLoop)

    // allow daemon style process
    emptyChanToKeepAppRunning := make(chan bool)
    <-emptyChanToKeepAppRunning
}

func bindEvents(DOM browser.DOM) {
    // let's handle windows resize
    var resizeEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        resizeEvent(DOM, args[0])
        return nil
    })

    DOM.Window.Call("addEventListener", "resize", resizeEventHandler)
}

func resizeEvent(DOM browser.DOM, event js.Value) {
    windowsWidth := js.Global().Get("innerWidth").Int()
    windowsHeight := js.Global().Get("innerHeight").Int()

    cvs.SetSize(windowsWidth, windowsHeight)
    go DOM.Log(fmt.Sprintf("resizeEvent x:%d y:%d", windowsWidth, windowsHeight))
}

func GameLoop(gc *draw2dimg.GraphicContext) bool {
    // render default color
    gc.SetFillColor(color.RGBA{0x18, 0x18, 0x18, 0xff})
    gc.Clear()

    rt := wolfenstein.NewRayTracer(gs)
    rays := rt.ComputeRays()

    renderGameView(gc, rays)
    renderMiniMap(gc, rays)

    handleMove()

    return true
}

func renderMiniMap(gc *draw2dimg.GraphicContext, rays []wolfenstein.Ray) {
    renderLevel(gc)
    renderRayCasting(gc, rays)
    renderPlayer(gc)
}

func renderRayCasting(gc *draw2dimg.GraphicContext, rays []wolfenstein.Ray) {
    gc.SetFillColor(color.RGBA{0xff, 0x00, 0x00, 0xff})
    gc.SetStrokeColor(color.RGBA{0xff, 0x00, 0x00, 0xff})

    for _, ray := range rays {
        // render raycast
        gc.BeginPath()
        gc.MoveTo(ray.Origin.X, ray.Origin.Y)
        gc.LineTo(ray.Impact.X, ray.Impact.Y)
        gc.Close()
        gc.FillStroke()
    }
}

func renderGameView(gc *draw2dimg.GraphicContext, rays []wolfenstein.Ray) {
    const screenHeight = 320
    const lineWidth = 8

    screenWidth := len(rays) * lineWidth

    up := &wolfenstein.Upscale{
        Source: wolfenstein.Resolution{screenWidth, screenHeight},
        Target: wolfenstein.Resolution{int(width), int(height)},
    }

    renderSky(gc, screenWidth, screenHeight, up)
    renderGround(gc, screenWidth, screenHeight, up)

    for rayN, ray := range rays {
        angle := wolfenstein.FixAngle(ray.Origin.Angle + 30 - rayN*60/len(rays))
        ca := wolfenstein.DegToRad(float64(angle))

        // fix fisheye
        distT := ray.Distance * math.Cos(ca)

        lineH := float64(gs.GetMapSize()*screenHeight) / distT
        if lineH > screenHeight {
            lineH = screenHeight
        }

        lineOffset := (screenHeight / 2) - lineH/2

        var c color.RGBA

        if ray.Impact.Type == wolfenstein.Horizontal {
            c = color.RGBA{0xcc, 0xcc, 0xcc, 0xff}
        } else {
            c = color.RGBA{0xff, 0xff, 0xff, 0xff}
        }

        if ray.Impact.CellType == wolfenstein.Door {
            c.R = 0
            c.G = 0
        }

        if ray.Impact.CellType == wolfenstein.Window {
            c.R = 0
        }

        if ray.Impact.CellType == wolfenstein.Checkerboard {
            c.R = 0
            c.B = 0
        }

        gc.SetFillColor(c)
        gc.SetStrokeColor(c)

        draw2dkit.Rectangle(
            gc,
            up.ScaleWidth(float64(rayN*lineWidth)),
            up.ScaleHeight(lineOffset),
            up.ScaleWidth(float64(rayN*lineWidth)+lineWidth),
            up.ScaleHeight(lineH+lineOffset),
        )
        gc.FillStroke()
    }
}

func renderSky(gc *draw2dimg.GraphicContext, screenWidth, screenHeight int, up *wolfenstein.Upscale) {
    skyColor := color.RGBA{0x00, 0x00, 0x99, 0xff}

    gc.SetFillColor(skyColor)
    gc.SetStrokeColor(skyColor)

    draw2dkit.Rectangle(
        gc,
        0,
        0,
        up.ScaleWidth(float64(screenWidth)),
        up.ScaleHeight(float64(screenHeight/2)),
    )
    gc.FillStroke()
}

func renderGround(gc *draw2dimg.GraphicContext, screenWidth, screenHeight int, up *wolfenstein.Upscale) {
    groundColor := color.RGBA{0x00, 0x99, 0x99, 0xff}

    gc.SetFillColor(groundColor)
    gc.SetStrokeColor(groundColor)

    draw2dkit.Rectangle(
        gc,
        0,
        up.ScaleHeight(float64(screenHeight/2)),
        up.ScaleWidth(float64(screenWidth)),
        up.ScaleHeight(float64(screenHeight)),
    )
    gc.FillStroke()
}

func handleMove() {
    actions := ctl.GetState()

    if actions.Action {
        gs.Action()
    }

    if actions.Up {
        gs.MoveUp()
    } else if actions.Down {
        gs.MoveDown()
    }

    if actions.Right {
        gs.TurnRight()
    } else if actions.Left {
        gs.TurnLeft()
    }
}

func renderLevel(gc *draw2dimg.GraphicContext) {
    gc.BeginPath()

    level := gs.GetLevel()
    blockSize := gs.GetBlockSize()
    mapSize := gs.GetMapSize()

    // minimap background
    gc.SetFillColor(color.RGBA{0x00, 0x00, 0x00, 0xff})
    gc.SetStrokeColor(color.RGBA{0x00, 0x00, 0x00, 0xff})

    draw2dkit.Rectangle(
        gc,
        0,
        0,
        float64(mapSize*blockSize),
        float64(mapSize*blockSize),
    )
    gc.FillStroke()

    gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xcc})
    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xcc})

    for y := 0; y < mapSize; y++ {
        for x := 0; x < mapSize; x++ {
            if level.Walls[x+y*mapSize] == 0 {
                // avoid useless rendering
                continue
            }

            c := color.RGBA{0xff, 0xff, 0xff, 0xcc}

            if level.Walls[x+y*mapSize] == wolfenstein.Door {
                c.R = 0
                c.G = 0
            }

            if level.Walls[x+y*mapSize] == wolfenstein.Window {
                c.R = 0
            }

            if level.Walls[x+y*mapSize] == wolfenstein.Checkerboard {
                c.R = 0
                c.B = 0
            }

            gc.SetFillColor(c)
            gc.SetStrokeColor(c)

            draw2dkit.Rectangle(
                gc,
                float64(x*blockSize+1),
                float64(y*blockSize+1),
                float64(x*blockSize+blockSize-2),
                float64(y*blockSize+blockSize-2),
            )
            gc.FillStroke()
        }
    }
}

func renderPlayer(gc *draw2dimg.GraphicContext) {
    // draw player on screen
    gc.SetFillColor(color.RGBA{0xff, 0xff, 0x00, 0xcc})
    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0x00, 0xcc})
    gc.BeginPath()

    // draw player on screen
    playerX, playerY, playerDeltaX, playerDeltaY := gs.GetPlayerPosition()
    draw2dkit.Circle(gc, playerX, playerY, 5)
    gc.FillStroke()

    // draw player direction
    gc.BeginPath()
    gc.MoveTo(playerX, playerY)
    gc.LineTo(playerX+playerDeltaX*5, playerY+playerDeltaY*5)
    gc.Close()
    gc.FillStroke()
}
