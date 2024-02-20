package wolfenstein

import (
    "math"
)

type GameState struct {
    level     []int
    mapSize   int
    blockSize int

    player Player
}

type Player struct {
    Position Point
    Delta    Point
}

type Point struct {
    X     float64
    Y     float64
    Angle float64
}

func NewGameState(width, height int) (*GameState, error) {
    var gs GameState

    // silly level
    gs.level = []int{
        1, 1, 1, 1, 1, 1, 1, 1,
        1, 0, 1, 0, 0, 0, 0, 1,
        1, 0, 1, 0, 0, 0, 0, 1,
        1, 0, 1, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 0, 0, 1,
        1, 0, 0, 0, 0, 1, 0, 1,
        1, 0, 0, 0, 0, 0, 0, 1,
        1, 1, 1, 1, 1, 1, 1, 1,
    }

    gs.mapSize = 8
    gs.blockSize = 64

    gs.player = Player{
        Position: Point{
            float64(gs.mapSize * gs.blockSize / 2),
            float64(gs.mapSize * gs.blockSize / 2),
            0.0,
        },
        Delta: Point{0, 0, 0.0},
    }

    gs.updateDelta()

    return &gs, nil
}

func (gs *GameState) GetMapSize() int {
    return gs.mapSize
}
func (gs *GameState) GetMapValue(x, y int) int {
    index := x + y*gs.mapSize

    if index < 0 || index >= len(gs.level) {
        return 0
    }

    return gs.level[x+y*gs.mapSize]
}

func (gs *GameState) GetLevel() []int {
    return gs.level
}

func (gs *GameState) GetPlayer() Player {
    return gs.player
}

func (gs *GameState) GetPlayerPosition() (x, y, deltaX, deltaY float64) {
    return gs.player.Position.X, gs.player.Position.Y, gs.player.Delta.X, gs.player.Delta.Y
}

func (gs *GameState) GetBlockSize() int {
    return gs.blockSize
}

func (gs *GameState) GetPlayerAngle() float64 {
    return gs.player.Position.Angle
}

func (gs *GameState) MoveUp() {
    gs.player.Position.X += gs.player.Delta.X
    gs.player.Position.Y += gs.player.Delta.Y
}

func (gs *GameState) MoveDown() {
    gs.player.Position.X -= gs.player.Delta.X
    gs.player.Position.Y -= gs.player.Delta.Y
}

func (gs *GameState) MoveLeft() {
    gs.player.Position.Angle -= 0.1

    if gs.player.Position.Angle < 0 {
        gs.player.Position.Angle += 2 * math.Pi
    }

    gs.updateDelta()
}

func (gs *GameState) MoveRight() {
    gs.player.Position.Angle += 0.1

    if gs.player.Position.Angle > 2*math.Pi {
        gs.player.Position.Angle -= 2 * math.Pi
    }

    gs.updateDelta()
}

func (gs *GameState) updateDelta() {
    gs.player.Delta.X = math.Cos(gs.player.Position.Angle) * 5
    gs.player.Delta.Y = math.Sin(gs.player.Position.Angle) * 5
}
