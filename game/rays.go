package game

type Ray struct {
    Origin   Point
    Impact   Impact
    Distance float64
}

type ImpactType string

const (
    Horizontal ImpactType = "horizontal"
    Vertical              = "vertical"
)

type Impact struct {
    X float64
    Y float64

    CellX    int
    CellY    int
    CellType Cell

    Type ImpactType
}
