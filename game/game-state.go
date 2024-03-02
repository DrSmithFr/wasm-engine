package game

type GameState struct {
    level  *Map
    player *Player

    width  int
    height int
}

func NewGameState(width, height int) *GameState {
    return &GameState{
        width:  width,
        height: height,
    }
}

func (gs *GameState) LoadLevel(level Map) {
    // silly level
    gs.level = &level

    gs.player = &Player{
        Position: Position{
            level.Spawn.X,
            level.Spawn.Y,
            level.Spawn.Z,
            level.SpawnAngle,
            0,
        },
    }
}

func (gs *GameState) GetSize() (int, int) {
    return gs.width, gs.height
}

func (gs *GameState) SetSize(width, height int) {
    gs.width = width
    gs.height = height
}

func (gs *GameState) GetLevel() *Map {
    return gs.level
}

func (gs *GameState) GetPlayer() *Player {
    return gs.player
}

func (gs *GameState) GetCollisionAt(x, y float64) Collision {
    return None
}
