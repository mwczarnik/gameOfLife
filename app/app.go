package app

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type Game struct {
	Width       int
	Height      int
	Population  int
	Quit        <-chan bool
	RefreshTime time.Duration
}

func (g Game) Run() {

	g.RefreshTime = time.Millisecond * 100

	// hide  blinking cursor
	fmt.Print("\033[?25l")
	// initialize board
	board := newBoard(g.Width, g.Height)
	// set random cells on board to be alive
	populateBoard(board, g.Population, g.Width, g.Height)

	// first generation
	// printGeneration(board)

	for {
		// handle exit
		select {
		case <-g.Quit:
			return
		default:
			applyRules(board, g.Width, g.Height)
			time.Sleep(g.RefreshTime)
			fmt.Print("\033[H\033[2J")
		}

	}

}

func newBoard(width, height int) [][]Cell {

	// rows
	board := make([][]Cell, height)

	for row := range board {
		// columns
		board[row] = make([]Cell, width)

		// initiate new cell for every slot in board
		for col := range board[row] {
			board[row][col] = newCell(false, Position{Col: col, Row: row})
		}
	}
	return board
}

func populateBoard(board [][]Cell, num, width, height int) {
	population := []Position{}

	// Generate random board positions
	for i := 0; i < num; i++ {
		population = append(population, Position{Col: rand.Intn(width - 1), Row: rand.Intn(height - 1)})
	}
	for _, cell := range population {
		board[cell.Row][cell.Col].changeState()
	}
}

func printGeneration(board [][]Cell) {
	output := ""
	for _, row := range board {
		for _, cell := range row {
			if cell.State {
				output += "■ "
				// fmt.Printf("■ ")
			} else {
				// output += "□ "
				output += "  "
				// fmt.Printf("  ")
			}
		}
		// fmt.Printf("\n")
		output += "\n"
	}
	fmt.Println(output)
	// return output
}

func applyRules(board [][]Cell, width, height int) {
	var (
		cellsToChange = []*Cell{}
		alive         = 0
		cell          = &Cell{}
	)

	for row := 0; row < height; row++ {
		for col := 0; col < width; col++ {

			// Check cell neighobours
			cell = &board[row][col]
			alive = checkNeighbours(*cell, board, width, height)

			// Conway rules
			switch {
			case cell.State && (alive <= 1 || alive > 3):
				cellsToChange = append(cellsToChange, cell)
			case !cell.State && alive == 3:
				cellsToChange = append(cellsToChange, cell)
			}
		}
	}

	applyChanges(cellsToChange)
	printGeneration(board)
}

func checkNeighbours(c Cell, board [][]Cell, width, height int) int {
	var (
		alive int
	)
	// Get fields from struct
	fields := reflect.ValueOf(c.Neighbours)
	var neighbour Position

	for i := 0; i < fields.NumField(); i++ {
		neighbour = fields.Field(i).Interface().(Position)

		// Check if neigh cell exists and if state is true
		if neighbour.Row >= 0 &&
			neighbour.Col >= 0 &&
			neighbour.Row < height-1 &&
			neighbour.Col < width-1 &&
			board[neighbour.Row][neighbour.Col].State {
			// increment num of alive neighbours
			alive++
		}
	}
	return alive
}

func applyChanges(cells []*Cell) {
	for _, cell := range cells {
		cell.changeState()
	}
}
