package main

import (
    "fmt"
    "go-webgl/browser"
    "go-webgl/controller"
    "go-webgl/game"
    "go-webgl/render"
    "image/color"
    "log"
)

var DOM browser.DOM
var controls controller.Interface
var gs *game.GameState

var width int
var height int

func main() {
    log.Println("Starting the game engine")

    // loading DOM to memory
    DOM = browser.LoadDOM()
    width, height = DOM.GetScreenSize()
    log.Println("Screen size is", width, "x", height)

    // creating render in order to render the game
    gameView := render.NewDirectCtx(width, height)
    minimapView := render.NewDirectCtx(200, 200)
    mapView := render.NewDirectCtx(width*80/100, height*60/100)

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
    log.Println("Binding engine and controls to DOM elements and events")
    gameView.Init(DOM)
    minimapView.Init(DOM)
    controls.Init(DOM)

    // creating game state
    log.Println("Creating game state")
    gs = game.NewGameState(width, height)

    // load level
    level := game.Map{
        Spawn:      game.Point3D{50, 50, 0},
        SpawnAngle: 0,
        Sectors: []game.Sector{{
            Walls: []game.Wall{{
                Points:  [2]game.Point3D{{70, 20, 0}, {70, 70, 0}},
                Ceiling: 10,
                Floor:   0,
            }},
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
        if false {
            rotate := fmt.Sprintf("rotate(%ddeg)", gs.GetPlayer().Position.Angle+180)
            r.GetCanvas().SetCssProperty("transform", rotate)
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
