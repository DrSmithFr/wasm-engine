package main

import (
    "go-webgl/browser"
    "go-webgl/controller"
    "go-webgl/game"
    "go-webgl/render"
    "image/color"
    "log"
    "time"
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

    engine := render.NewDirectCtx(width, height, 0, 0)

    controls = controller.NewKeyboardOnly()

    // setting up everything
    log.Println("Binding engine and controls to DOM elements and events")
    engine.Init(DOM)
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
                Ceiling: 100,
                Floor:   0,
            }},
        }},
    }

    log.Printf("Loading level: %+v", level)
    gs.LoadLevel(level)

    // start the game loop
    log.Println("Starting the game loop")
    engine.Start(30, GameLoop)

    // avoid WebAssembly to exit the program
    log.Println("WASM is keep running forever, waiting for a signal to stop it.")
    run := make(chan bool)
    <-run

    log.Println("Game engine stopped")
}

func GameLoop(r render.Renderer) bool {
    start := time.Now()
    r.Clear()

    // render background
    log.Println("Rendering background")
    r.SetColor(color.RGBA{0, 0, 0, 255})
    width, height = r.Size()
    r.DrawRect(0, 0, float64(width), float64(height))

    // render sectors
    log.Println("Rendering sectors")
    r.SetColor(color.RGBA{255, 255, 255, 255})
    for _, sector := range gs.GetLevel().Sectors {
        log.Println("Rendering sector")
        render.RenderSector(r, *gs.GetPlayer(), sector)
    }

    // render minimap
    log.Println("Rendering minimap")
    render.RenderMinimap(r, *gs.GetPlayer(), *gs.GetLevel())

    // handle move
    log.Println("Handling moves")
    handleMove()

    // flush the buffer
    log.Println("Flushing frame to rendering canvas")
    r.Flush()

    // calculate time elapsed
    elapsed := time.Since(start)
    log.Printf("Frame took %s", elapsed.String())

    return true
}

func handleMove() {
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
}
