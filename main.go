package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"reflect"
	"time"
)

var (
	width       = 50
	height      = 50
	population  = 2000
	refreshTime = time.Millisecond * 100
)

func handleInterrupt() {
	// CTRL + C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			// enable blinking cursor
			fmt.Print("\033[?25h")
			os.Exit(1)
		}
	}()
}

func startGame() {

	board := newBoard(width, height)
	// set random cells on board to be alive
	populateBoard(board, population)

	printGeneration(board)

	for {
		applyRules(board)
		time.Sleep(refreshTime)
		fmt.Print("\033[H\033[2J")

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
			board[row][col] = newCell(false, Position{X: col, Y: row})
		}
	}
	return board
}

func populateBoard(board [][]Cell, num int) {
	population := []Position{}

	// Generate random board positions
	for i := 0; i < num; i++ {
		population = append(population, Position{X: rand.Intn(width - 1), Y: rand.Intn(height - 1)})
	}
	for _, cell := range population {
		board[cell.Y][cell.X].changeState()
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
}

func applyRules(board [][]Cell) {
	var (
		cellsToChange = []*Cell{}
		alive         = 0
		cell          = &Cell{}
	)

	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[0]); col++ {

			// Check cell neighobours
			cell = &board[row][col]
			alive = checkNeighbours(*cell, board)

			// Conway rules
			switch {
			case cell.State && (alive <= 1 || alive > 3):
				cellsToChange = append(cellsToChange, cell)
			case !cell.State && alive == 3:
				cellsToChange = append(cellsToChange, cell)
			}
		}
	}
	//
	applyChanges(cellsToChange)
	printGeneration(board)
}

func checkNeighbours(c Cell, board [][]Cell) int {
	var (
		alive int
	)
	// Get fields from struct
	fields := reflect.ValueOf(c.Neighbours)
	var neighbour Position

	for i := 0; i < fields.NumField(); i++ {
		neighbour = fields.Field(i).Interface().(Position)

		// Check if neigh cell exists and if state is true
		if neighbour.Y >= 0 &&
			neighbour.X >= 0 &&
			neighbour.Y < height-1 &&
			neighbour.X < width-1 &&
			board[neighbour.Y][neighbour.X].State {
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

func main() {

	//TODO bubbletea TUI
	// p := tea.NewProgram(initialModel())
	// if err := p.Start(); err != nil {
	// 	fmt.Printf("There's been an error: %v", err)
	// 	os.Exit(1)
	// }

	// hide  blinking cursor
	fmt.Print("\033[?25l")

	handleInterrupt()

	startGame()

}
