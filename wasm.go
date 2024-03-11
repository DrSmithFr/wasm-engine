package main

import (
    "fmt"
    "go-webgl/browser"
    "go-webgl/browser/video"
    "go-webgl/controller"
    "go-webgl/game"
    "go-webgl/render"
    "image/color"
    "log"
    "syscall/js"
)

var DOM browser.Document
var controls controller.Interface
var gs *game.GameState

var width int
var height int

func main() {
    log.Println("Starting the game engine")

    // loading Document to memory
    DOM = browser.Load()
    width, height = DOM.GetScreenSize()
    log.Println("Screen size is", width, "x", height)

    // creating render in order to render the game
    gameView := render.NewDirectCtx(608, 480)
    minimapView := render.NewDirectCtx(200, 200)
    mapView := render.NewDirectCtx(width*80/100, height*60/100)

    // setting up the game position
    gameView.Canvas.SetCssProperty("position", "absolute")
    gameView.Canvas.SetCssProperty("top", "50%")
    gameView.Canvas.SetCssProperty("left", "50%")
    gameView.Canvas.SetCssProperty("translate", "-50% -50%")

    // setting up the map position
    mapView.Canvas.SetCssProperty("position", "absolute")
    mapView.Canvas.SetCssProperty("top", "50%")
    mapView.Canvas.SetCssProperty("left", "50%")
    mapView.Canvas.SetCssProperty("translate", "-50% -50%")
    mapView.Canvas.SetCssProperty("border", "5px solid #666")
    mapView.Canvas.SetCssProperty("border-radius", "30px")
    mapView.Canvas.SetCssProperty("background", "#252525")

    // setting up the minimap position
    minimapView.Canvas.SetCssProperty("position", "absolute")
    minimapView.Canvas.SetCssProperty("top", "30px")
    minimapView.Canvas.SetCssProperty("right", "30px")
    minimapView.Canvas.SetCssProperty("border", "5px solid #666")
    minimapView.Canvas.SetCssProperty("border-radius", "50%")
    minimapView.Canvas.SetCssProperty("background", "#25252575")

    controls = controller.NewKeyboardOnly()

    // setting up everything
    log.Println("Binding engine and controls to Document elements and events")
    gameView.Init(DOM)
    mapView.Init(DOM)
    minimapView.Init(DOM)
    controls.Init(DOM)

    // handle the window resize event
    DOM.Window.Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
        width, height = DOM.GetScreenSize()
        gameView.SetSize(width, height)
        mapView.SetSize(width*80/100, height*60/100)
        return nil
    }))

    // creating game state
    log.Println("Creating game state")
    gs = game.NewGameState(width, height)

    // load level
    level := game.Map{
        Spawn:      game.Point3D{90, 90, 0},
        SpawnAngle: 0,
        Sectors: []game.Sector{{
            Walls: []game.Wall{
                {
                    Points:  [2]game.Point3D{{40, 40, 0}, {120, 40, 0}},
                    Color:   color.RGBA{255, 0, 0, 255},
                    Ceiling: 10,
                    Floor:   0,
                },
                {
                    Points:  [2]game.Point3D{{120, 40, 0}, {160, 90, 0}},
                    Color:   color.RGBA{0, 255, 0, 255},
                    Ceiling: 10,
                    Floor:   0,
                },
                {
                    Points:  [2]game.Point3D{{160, 90, 0}, {120, 140, 0}},
                    Color:   color.RGBA{0, 255, 255, 255},
                    Ceiling: 10,
                    Floor:   0,
                },
                {
                    Points:  [2]game.Point3D{{120, 140, 0}, {40, 120, 0}},
                    Color:   color.RGBA{255, 0, 255, 255},
                    Ceiling: 10,
                    Floor:   0,
                },
                {
                    Points:  [2]game.Point3D{{40, 120, 0}, {40, 40, 0}},
                    Color:   color.RGBA{255, 50, 255, 255},
                    Ceiling: 10,
                    Floor:   0,
                },
            },
        }},
    }

    log.Printf("Loading level: %+v", level)
    gs.LoadLevel(level)

    // start the game loop
    log.Println("Starting the game loop")

    // Abusing the requestAnimationFrame to handle moves
    gameView.Start(120, handleMove)

    // Rendering the game
    gameView.Start(30, GameLoop)

    // Rendering the map
    mapView.Start(30, func(r render.Renderer) bool {
        actions := controls.GetState()

        if actions.ShowMap {
            r.GetCanvas().SetCssProperty("display", "block")
            render.RenderMap(r, *gs.GetPlayer(), *gs.GetLevel())
            return true
        } else {
            r.GetCanvas().SetCssProperty("display", "none")
        }

        return false
    })

    // Rendering the minimap
    minimapView.Start(25, func(r render.Renderer) bool {
        render.RenderMinimap(r, *gs.GetPlayer(), *gs.GetLevel())

        // Make the minimap canvas rotate with the player direction
        // just for fun and to show how to manipulate the canvas
        if controls.GetState().LockMap {
            rotate := fmt.Sprintf("rotate(%ddeg)", gs.GetPlayer().Position.Angle+180)
            r.GetCanvas().SetCssProperty("transform", rotate)
            r.GetCanvas().SetCssProperty("border-color", "red !important")
        } else {
            r.GetCanvas().SetCssProperty("transform", "rotate(0deg)")
            r.GetCanvas().SetCssProperty("border-color", "#666 !important")
        }

        return true
    })

    // avoid WebAssembly to exit the program
    log.Println("WASM is keep running forever, waiting for a signal to stop it.")
    run := make(chan bool)
    <-run

    log.Println("Game engine stopped")
}

func GameLoop(r render.Renderer) bool {
    // clear the buffer
    r.Clear()

    // render background
    r.SetColor(color.RGBA{0, 0, 0, 255})
    width, height = r.Size()
    r.DrawRect(0, 0, float64(width), float64(height))

    // render sectors
    r.SetColor(color.RGBA{255, 255, 255, 255})
    for _, sector := range gs.GetLevel().Sectors {
        render.RenderSector(r, *gs.GetPlayer(), sector)
    }

    // flush the buffer
    r.Flush()

    // testing cinematic
    player := video.BindById("cinematic")
    player.Play()

    return true
}

func handleMove(r render.Renderer) bool {
    actions := controls.GetState()
    player := gs.GetPlayer()

    // Handle actions first
    if actions.Action {
        player.Action()
    }

    // Handle movement allowing only one direction at a time
    if actions.Up {
        player.MoveUp()
    } else if actions.Down {
        player.MoveDown()
    } else if actions.Right {
        player.MoveRight()
    } else if actions.Left {
        player.MoveLeft()
    }

    // Handle rotation allowing only one direction at a time
    if actions.TurnRight {
        player.TurnRight()
    } else if actions.TurnLeft {
        player.TurnLeft()
    }

    return true
}
