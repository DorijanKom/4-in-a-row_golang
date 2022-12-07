package main

import (
	"fmt"
)

func main() {

	var row, col int

	fmt.Println("Please input the number of rows for your game.")
	fmt.Scan(&row)

	fmt.Println("Please enter the number of columns for your game.")
	fmt.Scan(&col)

	board := genBoard(row, col)
	printBoard(board)

}

func genBoard(row, col int) [][]string {
	board := make([][]string, row)

	for i := range board {
		board[i] = make([]string, col)
	}

	return board
}

func printBoard(board [][]string) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Printf("[ %s ]", board[i][j])
		}
		fmt.Println()
	}

}
