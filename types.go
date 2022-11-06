package main

type Position struct {
	X int
	Y int
}

type Cell struct {
	State      bool
	Position   Position
	Neighbours Neighbours
}

func newCell(state bool, p Position) Cell {
	c := Cell{State: state, Position: p}

	// set cell neighbours positions
	c.Neighbours.Top = Position{X: p.X, Y: p.Y - 1}
	c.Neighbours.TopLeft = Position{X: p.X - 1, Y: p.Y - 1}
	c.Neighbours.TopRight = Position{X: p.X + 1, Y: p.Y - 1}
	c.Neighbours.Bottom = Position{X: p.X, Y: p.Y + 1}
	c.Neighbours.BottomLeft = Position{X: p.X + -1, Y: p.Y + 1}
	c.Neighbours.BottomRight = Position{X: p.X + 1, Y: p.Y + 1}
	c.Neighbours.Left = Position{X: p.X - 1, Y: p.Y}
	c.Neighbours.Right = Position{X: p.X + 1, Y: p.Y}
	return c
}

func (c *Cell) changeState() {
	// From dead to alive or vice versa
	if c.State {
		c.State = false
	} else {
		c.State = true
	}

}

type Neighbours struct {
	Top         Position
	Bottom      Position
	Left        Position
	Right       Position
	TopLeft     Position
	TopRight    Position
	BottomLeft  Position
	BottomRight Position
}
