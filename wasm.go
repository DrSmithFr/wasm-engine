package main

import (
	"fmt"
	"go-webgl/controller"
	"go-webgl/dom"
	"go-webgl/game"
	"go-webgl/render"
	"image/color"
	"log"
	"syscall/js"
)

var controls controller.Interface
var gs *game.GameState

var width int
var height int

func main() {
	window := dom.Window()
	document := window.Document()

	document.Body().Style().
		SetProperty("background", "#000", false).
		SetProperty("margin", "0", false).
		SetProperty("padding", "0", false).
		SetProperty("overflow", "hidden", false).
		SetProperty("width", "100dvw", false).
		SetProperty("height", "100dvh", false)

	// loading Document to memory
	width, height = window.InnerSize()
	log.Println("Screen size is", width, "x", height)

	// creating canvas
	gameCanvas := document.CreateCanvasElement()
	mapCanvas := document.CreateCanvasElement()
	minimapCanvas := document.CreateCanvasElement()

	// setting up the game position
	gameCanvas.
		HTMLElement().
		Style().
		SetProperty("background", "#333", false).
		SetProperty("position", "absolute", false).
		SetProperty("top", "50%", false).
		SetProperty("left", "50%", false).
		SetProperty("translate", "-50% -50%", false).
		SetProperty("aspect-ratio", "4 / 3", false).
		SetProperty("height", "100dvh", false)

	// setting up the map position
	mapCanvas.
		HTMLElement().
		Style().
		SetProperty("background", "#333", false).
		SetProperty("position", "absolute", false).
		SetProperty("top", "50%", false).
		SetProperty("left", "50%", false).
		SetProperty("translate", "-50% -50%", false).
		SetProperty("border", "5px solid #666", false).
		SetProperty("border-radius", "30px", false).
		SetProperty("background", "#252525", false)

	// setting up the minimap position
	minimapCanvas.
		HTMLElement().
		Style().
		SetProperty("background", "#333", false).
		SetProperty("position", "absolute", false).
		SetProperty("top", "30px", false).
		SetProperty("right", "30px", false).
		SetProperty("border", "5px solid #666", false).
		SetProperty("border-radius", "50%", false).
		SetProperty("background", "#25252575", false)

	// appending canvas to the document
	document.Body().Element().AppendChild(gameCanvas)
	document.Body().Element().AppendChild(mapCanvas)
	document.Body().Element().AppendChild(minimapCanvas)

	// creating render in order to render the game
	gameView := render.NewDirectCtx(gameCanvas)
	mapView := render.NewDirectCtx(mapCanvas)
	minimapView := render.NewDirectCtx(minimapCanvas)

	// fix render size
	width, height = window.InnerSize()
	gameView.SetSize(width, height)
	mapView.SetSize(width*80/100, height*60/100)
	minimapView.SetSize(200, 200)

	controls = controller.NewKeyboardOnly()

	// setting up everything
	log.Println("Binding engine and controls to Document elements and events")
	gameView.Init(window)
	mapView.Init(window)
	minimapView.Init(window)

	controls.Init(window)

	// handle the window resize event
	window.Js().Call("addEventListener", "resize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		width, height = window.InnerSize()
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
			r.GetCanvas().HTMLElement().Style().SetProperty("display", "block", false)
			render.RenderMap(r, *gs.GetPlayer(), *gs.GetLevel())
			return true
		} else {
			r.GetCanvas().HTMLElement().Style().SetProperty("display", "none", false)
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
			r.GetCanvas().
				HTMLElement().
				Style().
				SetProperty("transform", rotate, false).
				SetProperty("border-color", "red", true)
		} else {
			r.GetCanvas().
				HTMLElement().
				Style().
				SetProperty("transform", "rotate(0deg)", false).
				SetProperty("border-color", "#666", true)
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
