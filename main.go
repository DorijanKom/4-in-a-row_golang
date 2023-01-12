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
	player1    string = "◯"
	player2    string = "⬤"
)

var movesPlayerOne []int
var movesPlayerTwo []int

func main() {

	var row, col, key, userInput int
	turn := 0

	for {

		fmt.Println("Press 1 for default board size - Press 2 for custom size")
		fmt.Scan(&userInput)

		if userInput == 1 {
			row = 6
			col = 7
		}

		if userInput == 2 {
			fmt.Println("Please input the number of rows for your game.")
			fmt.Scan(&row)

			fmt.Println("Please enter the number of columns for your game.")
			fmt.Scan(&col)

		}

		if checkRowCol(row, col) {
			board := genBoard(row, col)

			board.printBoard()
			fmt.Println("input")
			for {
				if turn%2 == 0 {                    // if turn counter is even then it's player1's turn, else it's player2's turn
					fmt.Scan(&key)                      // takes user input
					err := board.makeMove(key, player1) //cals make move function and saves response into the err variable
					if err != nil {                     // checks for errors and handles them accordingly
						board.printBoard()
						fmt.Println(err)
						continue
					}
					board.printBoard()
							turn++
			        moveHistory(key, &movesPlayerOne)
			        fmt.Printf("%v ", movesPlayerOne)
			        fmt.Println()
			        fmt.Printf("%v ", movesPlayerTwo)
          
		  } else {
			  fmt.Scan(&key)
			  err := board.makeMove(key, player2)
			  if err != nil {
				board.printBoard()
				fmt.Println(err)
				continue
			}
			board.printBoard()
			  turn++
			  moveHistory(key, &movesPlayerTwo)
			  fmt.Printf("%v ", movesPlayerOne)
			  fmt.Println()
			  fmt.Printf("%v ", movesPlayerTwo)
				}

			}
		}

		fmt.Println("The size difference between rows and cols must be at most 2")
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
			fmt.Printf(" [ %s  ] ", board.state[i][j])
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


func moveHistory(key int, moves *[]int) {

	*moves = append(*moves, key)


func checkRowCol(row, col int) bool {
	return col-row == 2 || col-row <= 2 && col > 6 && row > 5
}
