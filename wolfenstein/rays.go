package wolfenstein

type Ray struct {
    Origin   Point
    Impact   Impact
    Distance float64
}

type Impact struct {
    X float64
    Y float64

    CellX int
    CellY int
}
