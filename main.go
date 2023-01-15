package main

import (
	"burch/4-in-a-row/boardpackage"
	"fmt"
)

type Board struct {
	State          [][]string
	Rows           int
	Cols           int
	MovesPlayerOne []int
	MovesPlayerTwo []int
}

const (
	player1 string = "◯"
	player2 string = "⬤"
)

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
			board := boardpackage.GenBoard(row, col)

			board.PrintBoard()
			fmt.Println("input")
			for {
				if turn%2 == 0 { // if turn counter is even then it's player1's turn, else it's player2's turn
					fmt.Scan(&key)                      // takes user input
					err := board.MakeMove(key, player1) //cals make move function and saves response into the err variable
					if err != nil {                     // checks for errors and handles them accordingly
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					turn++
					moveHistory(key, &board.MovesPlayerOne)
					fmt.Printf("%v ", board.MovesPlayerOne)
					fmt.Println()
					fmt.Printf("%v ", board.MovesPlayerTwo)

				} else {
					fmt.Scan(&key)
					err := board.MakeMove(key, player2)
					if err != nil {
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					turn++
					moveHistory(key, &board.MovesPlayerTwo)
					fmt.Printf("%v ", board.MovesPlayerOne)
					fmt.Println()
					fmt.Printf("%v ", board.MovesPlayerTwo)
				}

			}
		}

		fmt.Println("The size difference between rows and cols must be at most 2")

	}
}

func moveHistory(key int, moves *[]int) {

	*moves = append(*moves, key)
}

func checkRowCol(row, col int) bool {
	return col-row == 2 || col-row <= 2 && col > 6 && row > 5
}
