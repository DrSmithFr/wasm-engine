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
    up    bool
    down  bool
    left  bool
    right bool
}

var keyboard = move{false, false, false, false}

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
    cvs.Start(30, Render)

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
    }

    //go DOM.Log(fmt.Sprintf("key up:%s", code))
}

func clickEvent(DOM browser.DOM, event js.Value) {
    mouseX := event.Get("clientX").Int()
    mouseY := event.Get("clientY").Int()

    go DOM.Log(fmt.Sprintf("mouseEvent x:%d y:%d", mouseX, mouseY))
}

func Render(gc *draw2dimg.GraphicContext) bool {
    // render default color
    gc.SetFillColor(color.RGBA{0x18, 0x18, 0x18, 0xff})
    gc.Clear()

    rt := wolfenstein.NewRayTracer(gs)
    rays := rt.ComputeRays()

    render2D(gc, rays)
    render3D(gc, rays)

    handleMove()

    return true
}

func render2D(gc *draw2dimg.GraphicContext, rays []*wolfenstein.Ray) {
    renderLevel(gc)
    renderRayCasting(gc, rays)
    renderPlayer(gc)
}

func renderRayCasting(gc *draw2dimg.GraphicContext, rays []*wolfenstein.Ray) {
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

func render3D(gc *draw2dimg.GraphicContext, rays []*wolfenstein.Ray) {
    const lineHeight = 320
    const viewOffset = 530

    for rayN, ray := range rays {
        // render 3D walls
        ca := gs.GetPlayerAngle() - ray.Origin.Angle
        if ca < 0 {
            ca += 2 * math.Pi
        }
        if ca > 2*math.Pi {
            ca -= 2 * math.Pi
        }

        // fix fisheye
        distT := ray.Distance * math.Cos(ca)

        lineH := float64(gs.GetMapSize()*lineHeight) / distT
        if lineH > lineHeight {
            lineH = lineHeight
        }

        lineOffset := (lineHeight / 2) - lineH/2

        draw2dkit.Rectangle(
            gc,
            float64(rayN*8+viewOffset),
            lineOffset,
            float64(rayN*8+viewOffset)+8,
            lineH+lineOffset,
        )
        gc.FillStroke()
    }
}

func handleMove() {
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
    gc.SetFillColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0xff, 0xff})
    gc.BeginPath()

    level := gs.GetLevel()
    blockSize := gs.GetBlockSize()
    mapSize := gs.GetMapSize()

    for y := 0; y < mapSize; y++ {
        for x := 0; x < mapSize; x++ {
            if level[x+y*mapSize] == 0 {
                // avoid useless rendering
                continue
            }

            draw2dkit.Rectangle(
                gc,
                float64(x*blockSize+1),
                float64(y*blockSize+1),
                float64(x*blockSize+blockSize-1),
                float64(y*blockSize+blockSize-1),
            )
            gc.FillStroke()
        }
    }
}

func renderPlayer(gc *draw2dimg.GraphicContext) {
    // draw player on screen
    gc.SetFillColor(color.RGBA{0xff, 0xff, 0x00, 0xff})
    gc.SetStrokeColor(color.RGBA{0xff, 0xff, 0x00, 0xff})
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
