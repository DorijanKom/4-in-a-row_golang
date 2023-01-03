package main

import (
	"fmt"
)

func main() {

	var row, col, key int

	fmt.Println("Please input the number of rows for your game.")
	fmt.Scan(&row)

	fmt.Println("Please enter the number of columns for your game.")
	fmt.Scan(&col)

	fmt.Println("Enter the space you want to enter your key.")
	fmt.Scan(&key)

	board := genBoard(row, col)
	printBoard(board, key)

}

func genBoard(row, col int) [][]string {
	board := make([][]string, row)

	for i := range board {
		board[i] = make([]string, col)
	}

	return board
}

func printBoard(board [][]string, key int) {

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[len(board)-1][key-1] == "" {
				board[len(board)-i-1][key-1] = "x"
			}
			fmt.Printf(" [ %s ] ", board[i][j])
		}
		fmt.Println()
	}

}
