package main

import (
    "fmt"
    "github.com/llgcode/draw2d/draw2dimg"
    "github.com/llgcode/draw2d/draw2dkit"
    "go-webgl/canvas"
    "go-webgl/wolfenstein"
    "image/color"
    "math"
    "runtime"
    "syscall/js"
)

var DOM *browser.DOM
var cvs *browser.Canvas2d
var gs *wolfenstein.GameState

type move struct {
    up     bool
    down   bool
    left   bool
    right  bool
    action bool
}

var keyboard = move{false, false, false, false, false}

var width float64
var height float64

func main() {
    // loading DOM to memory
    DOM = browser.LoadDOM()

    // setting up everything
    bindEvents(*DOM)

    // create canvas
    cvs, _ = browser.NewCanvas2d(false)
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

    // let's handle key down
    var keydownEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        keydownEvent(DOM, args[0])
        return nil
    })

    var keyupEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        keyupEvent(DOM, args[0])
        return nil
    })

    DOM.Document.Call("addEventListener", "keydown", keydownEventHandler)
    DOM.Document.Call("addEventListener", "keyup", keyupEventHandler)

    // let's handle that mouse pointer down
    var mouseEventHandler = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        clickEvent(DOM, args[0])
        return nil
    })

    DOM.Window.Call("addEventListener", "pointerdown", mouseEventHandler)
}

func resizeEvent(DOM browser.DOM, event js.Value) {
    windowsWidth := js.Global().Get("innerWidth").Int()
    windowsHeight := js.Global().Get("innerHeight").Int()

    cvs.SetSize(windowsWidth, windowsHeight)

    width = float64(windowsWidth)
    height = float64(windowsHeight)

    go DOM.Log(fmt.Sprintf("resizeEvent x:%d y:%d", windowsWidth, windowsHeight))
}

func keydownEvent(DOM browser.DOM, event js.Value) {
    code := event.Get("code").String()

    switch code {
    case "ArrowUp", "KeyW":
        keyboard.up = true
    case "ArrowDown", "KeyS":
        keyboard.down = true
    case "ArrowRight", "KeyD":
        keyboard.right = true
    case "ArrowLeft", "KeyA":
        keyboard.left = true
    case "KeyE":
        keyboard.action = true
    }

    //go DOM.Log(fmt.Sprintf("key down:%s", code))
}

func keyupEvent(DOM browser.DOM, event js.Value) {
    code := event.Get("code").String()

    switch code {
    case "ArrowUp", "KeyW":
        keyboard.up = false
    case "ArrowDown", "KeyS":
        keyboard.down = false
    case "ArrowRight", "KeyD":
        keyboard.right = false
    case "ArrowLeft", "KeyA":
        keyboard.left = false
    case "KeyE":
        keyboard.action = false
    }

    //go DOM.Log(fmt.Sprintf("key up:%s", code))
}

func clickEvent(DOM browser.DOM, event js.Value) {
    mouseX := event.Get("clientX").Int()
    mouseY := event.Get("clientY").Int()

    go DOM.Log(fmt.Sprintf("mouseEvent x:%d y:%d", mouseX, mouseY))
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
        // render 3D walls
        ca := ray.Origin.Angle - gs.GetPlayerAngle()
        if ca < 0 {
            ca += 2 * math.Pi
        }
        if ca > 2*math.Pi {
            ca -= 2 * math.Pi
        }

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

    if keyboard.action {
        gs.Action()
    }

    if keyboard.up {
        gs.MoveUp()
    } else if keyboard.down {
        gs.MoveDown()
    }

    if keyboard.right {
        gs.MoveRight()
    } else if keyboard.left {
        gs.MoveLeft()
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
