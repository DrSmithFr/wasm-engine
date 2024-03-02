package game

import (
    "math"
)

const MovementMomentum = 5
const AngularMomentum = 5
const frontalOffset = 15
const strafeOffset = 5

type Position struct {
    X float64
    Y float64
    Z float64

    Angle          int
    VerticalOffset int
}

func (p *Position) Rad() float64 {
    return DegToRad(p.Angle)
}

func (p *Position) GetDelta() (float64, float64) {
    x := math.Cos(p.Rad()) * MovementMomentum
    y := math.Sin(p.Rad()) * MovementMomentum
    return x, y
}

type Player struct {
    Game     *GameState
    Position Position
}

func (p *Player) MoveUp() {
    deltaX, deltaY := p.Position.GetDelta()

    posX := p.Position.X + deltaX
    posY := p.Position.Y + deltaY

    var xo, yo float64

    if deltaX > 0 {
        xo = frontalOffset
    } else {
        xo = -frontalOffset
    }

    if deltaY > 0 {
        yo = frontalOffset
    } else {
        yo = -frontalOffset
    }

    // allow front wall sliding
    if p.Game.GetCollisionAt(posX+xo, p.Position.Y) == None {
        p.Position.X = posX
    }

    if p.Game.GetCollisionAt(p.Position.X, posY+yo) == None {
        p.Position.Y = posY
    }
}

func (p *Player) MoveDown() {
    deltaX, deltaY := p.Position.GetDelta()

    posX := p.Position.X - deltaX
    posY := p.Position.Y - deltaY

    if p.Game.GetCollisionAt(posX, p.Position.Y) == None {
        p.Position.X = posX
    }

    if p.Game.GetCollisionAt(p.Position.X, posY) == None {
        p.Position.Y = posY
    }
}

func (p *Player) MoveLeft() {
    deltaX, deltaY := p.Position.GetDelta()

    posX := p.Position.X + deltaY
    posY := p.Position.Y - deltaX

    var xo, yo float64

    if deltaY > 0 {
        xo = -strafeOffset
    } else {
        xo = strafeOffset
    }

    if deltaX > 0 {
        yo = -strafeOffset
    } else {
        yo = strafeOffset
    }

    // allow front wall sliding
    if p.Game.GetCollisionAt(posX+xo, p.Position.Y) == None {
        p.Position.X = posX
    }

    if p.Game.GetCollisionAt(p.Position.X, posY+yo) == None {
        p.Position.Y = posY
    }
}

func (p *Player) MoveRight() {
    deltaX, deltaY := p.Position.GetDelta()

    posX := p.Position.X - deltaY
    posY := p.Position.Y + deltaX

    var xo, yo float64

    if deltaY > 0 {
        xo = strafeOffset
    } else {
        xo = -strafeOffset
    }

    if deltaX > 0 {
        yo = strafeOffset
    } else {
        yo = -strafeOffset
    }

    // allow front wall sliding
    if p.Game.GetCollisionAt(posX+xo, p.Position.Y) == None {
        p.Position.X = posX
    }

    if p.Game.GetCollisionAt(p.Position.X, posY+yo) == None {
        p.Position.Y = posY
    }
}

func (p *Player) TurnLeft() {
    p.Position.Angle = FixAngle(p.Position.Angle - AngularMomentum)
}

func (p *Player) TurnRight() {
    p.Position.Angle = FixAngle(p.Position.Angle + AngularMomentum)
}

func (p *Player) Action() {
    deltaX, deltaY := p.Position.GetDelta()

    posX := p.Position.X + deltaX
    posY := p.Position.Y + deltaY

    var xo, yo float64
    const offset = 20

    if deltaX > 0 {
        xo = offset
    } else {
        xo = -offset
    }

    if deltaY > 0 {
        yo = offset
    } else {
        yo = -offset
    }

    // allow side door opening
    if p.Game.GetCollisionAt(posX+xo, p.Position.Y) == Door {
        // open door
    }

    if p.Game.GetCollisionAt(p.Position.X, posY+yo) == Door {
        // open door
    }
}
