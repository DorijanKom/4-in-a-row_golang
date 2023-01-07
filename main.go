package main

import (
	"fmt"
)

type Board struct {
	state         [][]string
	rows          int
	cols          int
	currentPlayer string
}

const (
	emptyField string = " "
	player1    string = "X"
	player2    string = "O"
)

func main() {

	var row, col, key int
	turn := 0

	fmt.Println("Please input the number of rows for your game.")
	fmt.Scan(&row)

	fmt.Println("Please enter the number of columns for your game.")
	fmt.Scan(&col)

	board := genBoard(row, col)

	fmt.Println("input")
	for {
		if turn%2 == 0 {
			fmt.Scan(&key)
			board.makeMove(key, player1)
			board.printBoard()
			turn++
		} else {
			fmt.Scan(&key)
			board.makeMove(key, player2)
			board.printBoard()
			turn++
		}

	}
}

func genBoard(row, col int) Board {

	// Creates new board struct
	board := Board{
		state:         make([][]string, row),
		rows:          row,
		cols:          col,
		currentPlayer: "",
	}

	// Generates the 2d array (state)
	for i := range board.state {
		board.state[i] = make([]string, col)
		for j := range board.state[i] {
			board.state[i][j] = emptyField
		}
	}

	return board
}

func (board *Board) printBoard() {

	for i := board.rows - 1; i >= 0; i-- {
		for j := 0; j < board.cols; j++ {
			fmt.Printf(" [ %s ] ", board.state[i][j])
		}
		fmt.Println()
	}

}

func (board *Board) makeMove(key int, piece string) error {
	if key < 0 || key-1 >= board.cols {
		return fmt.Errorf("invalid column")
	}
	if board.state[board.rows-1][key-1] != emptyField {
		return fmt.Errorf("the column is full")
	}

	for i := 0; i < board.rows; i++ {
		if board.state[i][key-1] == emptyField {
			board.state[i][key-1] = piece
			break
		}
	}

	return nil

}
