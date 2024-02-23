package wolfenstein

import (
    "math"
)

type GameState struct {
    level     Map
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
    Angle int
}

func (p Point) Rad() float64 {
    return float64(p.Angle) * (math.Pi / 180.)
}

func NewGameState(width, height int) (*GameState, error) {
    var gs GameState

    // silly level
    gs.level = Level0

    gs.mapSize = 8
    gs.blockSize = 32

    gs.player = Player{
        Position: Point{
            float64(gs.mapSize*gs.blockSize/2) + 10,
            float64(gs.mapSize*gs.blockSize/2) + 10,
            0,
        },
        Delta: Point{0, 0, 0.0},
    }

    gs.updateDelta()

    return &gs, nil
}

func (gs *GameState) GetMapSize() int {
    return gs.mapSize
}
func (gs *GameState) GetMapValue(x, y int) Cell {
    index := x + y*gs.mapSize

    if index < 0 || index >= len(gs.level.Walls) {
        return 0
    }

    return gs.level.Walls[x+y*gs.mapSize]
}

func (gs *GameState) GetMapValueAt(x, y float64) Cell {
    cellX := int(math.Trunc(x / float64(gs.blockSize)))
    cellY := int(math.Trunc(y / float64(gs.blockSize)))
    return gs.GetMapValue(cellX, cellY)
}

func (gs *GameState) SetMapValue(x, y int, cell Cell) {
    gs.level.Walls[x+y*gs.mapSize] = cell
}

func (gs *GameState) SetMapValueAt(x, y float64, cell Cell) {
    cellX := int(x / float64(gs.blockSize))
    cellY := int(y / float64(gs.blockSize))
    gs.SetMapValue(cellX, cellY, cell)
}

func (gs *GameState) GetLevel() Map {
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

func (gs *GameState) MoveUp() {
    posX := gs.player.Position.X + gs.player.Delta.X
    posY := gs.player.Position.Y + gs.player.Delta.Y

    var xo, yo float64
    const offset = 15

    if gs.player.Delta.X > 0 {
        xo = offset
    } else {
        xo = -offset
    }

    if gs.player.Delta.Y > 0 {
        yo = offset
    } else {
        yo = -offset
    }

    // allow front wall sliding
    if gs.GetMapValueAt(posX+xo, gs.player.Position.Y) == EmptyCell {
        gs.player.Position.X = posX
    }

    if gs.GetMapValueAt(gs.player.Position.X, posY+yo) == EmptyCell {
        gs.player.Position.Y = posY
    }
}

func (gs *GameState) MoveDown() {
    posX := gs.player.Position.X - gs.player.Delta.X
    posY := gs.player.Position.Y - gs.player.Delta.Y

    if gs.GetMapValueAt(posX, gs.player.Position.Y) == EmptyCell {
        gs.player.Position.X = posX
    }

    if gs.GetMapValueAt(gs.player.Position.X, posY) == EmptyCell {
        gs.player.Position.Y = posY
    }
}

const AngularMomentum = 5

func (gs *GameState) TurnLeft() {
    gs.player.Position.Angle = FixAngle(gs.player.Position.Angle - AngularMomentum)
    gs.updateDelta()
}

func (gs *GameState) TurnRight() {
    gs.player.Position.Angle = FixAngle(gs.player.Position.Angle + AngularMomentum)
    gs.updateDelta()
}

func (gs *GameState) updateDelta() {
    gs.player.Delta.X = math.Cos(gs.player.Position.Rad()) * 5
    gs.player.Delta.Y = math.Sin(gs.player.Position.Rad()) * 5
}

func (gs *GameState) Action() {
    posX := gs.player.Position.X + gs.player.Delta.X
    posY := gs.player.Position.Y + gs.player.Delta.Y

    var xo, yo float64
    const offset = 20

    if gs.player.Delta.X > 0 {
        xo = offset
    } else {
        xo = -offset
    }

    if gs.player.Delta.Y > 0 {
        yo = offset
    } else {
        yo = -offset
    }

    // allow side door opening
    if gs.GetMapValueAt(posX+xo, gs.player.Position.Y) == Door {
        gs.SetMapValueAt(posX+xo, gs.player.Position.Y, EmptyCell)
    }

    if gs.GetMapValueAt(gs.player.Position.X, posY+yo) == Door {
        gs.SetMapValueAt(gs.player.Position.X, posY+yo, EmptyCell)
    }
}
