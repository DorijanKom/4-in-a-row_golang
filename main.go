package main

import "fmt"

func main() {
	var row, col int

	fmt.Println("Please input the number of rows for your game.")
	fmt.Scan(&row)

	fmt.Println("Please enter the number of columns for your game.")
	fmt.Scan(&col)

	board := generateBoard(row, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			board[i][j] = "d"
			fmt.Printf("| %s ", board[i][j])
		}
		fmt.Printf("|")
		fmt.Println()
	}

}

func generateBoard(row, col int) [][]string {
	board := make([][]string, row)

	for i := range board {
		board[i] = make([]string, col)
	}

	return board
}
