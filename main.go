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

func main() {

	var row, col, key int
	turn := 0

	fmt.Println("Please input the number of rows for your game.")
	fmt.Scan(&row)

	fmt.Println("Please enter the number of columns for your game.")
	fmt.Scan(&col)

	board := genBoard(row, col)

	board.printBoard()
	fmt.Println("input")
	for {
		if turn%2 == 0 {
			fmt.Scan(&key)
			err := board.makeMove(key, player1)
			if err != nil {
				board.printBoard()
				fmt.Println(err)
				continue
			}
			board.printBoard()
			turn++
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

func (board *Board) endGame(piece string) bool {
	

	//horizontal check
	for j:=0;j<board.rows-3;j++{
			for i:=0;i<board.rows;i++{
					if board.state[i][j] == piece && board.state[i][j+1] == piece && board.state[i][j+2] == piece && board.state[i][j+3] == piece{
							fmt.Println("Victory")
						} 
				}
		}
	//vertical check
	for i:=0;i<board.cols-3;i++{
		for j:=0;j<board.rows;j++{
			if board.state[i][j]== piece && board.state[i][j+1] == piece && board.state[i][j+2] == piece && board.state[i][j+3] == piece{
				fmt.Println("Victory")
			}
		}
	}
	//diagonal check left
	for i:=3; i<board.cols;i++{
		for j:=0;j<board.rows-3;j++{
			if board.state[i][j] == piece && board.state[i-1][j+1] == piece && board.state[i-2][j+2] == piece && board.state[i-3][j+3]==piece{
				fmt.Println("Victory")
			}
		}
	}
	//diagonal check right
	for i:=0; i<board.cols;i++{
		for j:=3; j<board.rows-3; j++{
			if board.state[i][j] == piece && board.state[i-1][j-1] == piece && board.state[i-2][j-2] == piece && board.state [i-3][j-3]==piece{
				fmt.Println("Victory")
			}
		}
	}
  return false
} 
