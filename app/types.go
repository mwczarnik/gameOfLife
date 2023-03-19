package app

type Position struct {
	Col int
	Row int
}

type Cell struct {
	State      bool
	Position   Position
	Neighbours Neighbours
}

func newCell(state bool, p Position) Cell {
	c := Cell{State: state, Position: p}

	// set cell neighbours positions
	c.Neighbours.Top = Position{Col: p.Col, Row: p.Row - 1}
	c.Neighbours.TopLeft = Position{Col: p.Col - 1, Row: p.Row - 1}
	c.Neighbours.TopRight = Position{Col: p.Col + 1, Row: p.Row - 1}
	c.Neighbours.Bottom = Position{Col: p.Col, Row: p.Row + 1}
	c.Neighbours.BottomLeft = Position{Col: p.Col + -1, Row: p.Row + 1}
	c.Neighbours.BottomRight = Position{Col: p.Col + 1, Row: p.Row + 1}
	c.Neighbours.Left = Position{Col: p.Col - 1, Row: p.Row}
	c.Neighbours.Right = Position{Col: p.Col + 1, Row: p.Row}
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
