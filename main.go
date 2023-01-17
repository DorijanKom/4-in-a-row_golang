package main

import (
	"bufio"
	"burch/4-in-a-row/boardpackage"
	"fmt"
	"os"
	"strconv"
)

const (
	player1 string = "◯"
	player2 string = "⬤"
)

func main() {

	var row, col, key int
	fmt.Println("\nWelcome to 4-in-a-row Go edition!\n")
	reader := bufio.NewScanner(os.Stdin)

	for {

		fmt.Println("Press 1 new game - Press 2 to load an existing game - Press 3 to exit")
		reader.Scan()
		userInput, _ := strconv.ParseInt(reader.Text(), 10, 64)

		if userInput == 1 {
			fmt.Println("Press 1 for default size (6x7) - Press 2 for custom size")
			reader.Scan()
			gameInput, err := strconv.ParseInt(reader.Text(), 10, 64)
			if gameInput == 1 {
				row = 6
				col = 7
			} else if gameInput == 2 {
				fmt.Println("Please input the number of rows for your game.")
				reader.Scan()
				rowInput, _ := strconv.ParseInt(reader.Text(), 10, 64)
				row = int(rowInput)

				fmt.Println("Please enter the number of columns for your game.")
				reader.Scan()
				colInput, _ := strconv.ParseInt(reader.Text(), 10, 64)
				col = int(colInput)
			}

			if err != nil {
				fmt.Println("Invalid input!")
			}

			if !checkRowCol(row, col) {
				fmt.Println("\nThe size difference between rows and cols must be at most 2 and it cannot be less than 6x7!\n")
				continue
			}
			board := boardpackage.GenBoard(row, col)

			board.PrintBoard()
			fmt.Println("input")
			for {
				if board.Turn%2 == 0 { // if turn counter is even then it's player1's turn, else it's player2's turn
					bottomUi(reader, board, &key)
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerOne)
					board.PrintMoves()
					gameOver := checkForEnd(reader, player1, board)
					if !gameOver {
						break
					}
				} else {
					bottomUi(reader, board, &key)
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerTwo)
					board.PrintMoves()
					gameOver := checkForEnd(reader, player2, board)
					if !gameOver {
						break
					}
				}
			}
		}
		if userInput == 2 {
			board := boardpackage.GenBoard(6, 7)
			board.LoadGameList()
			for {
				if board.Turn%2 == 0 { // if turn counter is even then it's player1's turn, else it's player2's turn
					bottomUi(reader, board, &key)
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerOne)
					board.PrintMoves()
					//checkForEnd(reader, player1, board)
				} else {
					bottomUi(reader, board, &key)
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerTwo)
					board.PrintMoves()
					//checkForEnd(reader, player2, board)
				}
			}
		}
		if userInput == 3 {
			fmt.Println("Goodbye...")
			os.Exit(1)
		}

	}

}

func moveHistory(key int, moves *[]int) {

	*moves = append(*moves, key)
}

func checkRowCol(row, col int) bool {
	return row >= 6 && col >= 7 && (col-row) <= 2
}

func checkForEnd(reader *bufio.Scanner, piece string, board *boardpackage.Board) bool {
	gameOver, winner := board.EndGame()
	if gameOver {
		if winner == "Draw" {
			fmt.Println("The game is a draw.")
			fmt.Println("Play again? Y/N")
			reader.Scan()
			response := reader.Text()
			if response == "Y" || response == "y" {
				board.ResetBoard()
				return true
			} else if response == "N" || response == "n" {
				fmt.Println("Goodbye...")
				os.Exit(1)
			}
		} else {
			fmt.Printf("Player %s  is the winner\n", piece)
			fmt.Println("Play again? Y/N")
			reader.Scan()
			response := reader.Text()
			if response == "Y" || response == "y" {
				board.ResetBoard()
				return true
			} else if response == "N" || response == "n" {
				fmt.Println("Goodbye...")
				os.Exit(1)
			}
		}

	}
	return false
}

func bottomUi(reader *bufio.Scanner, board *boardpackage.Board, key *int) {
	reader.Scan()
	consoleInput := reader.Text()

	if consoleInput == "S" || consoleInput == "s" {
		reader.Scan()
		filename := reader.Text()
		board.SaveGame(filename)
		return
	} else if consoleInput == "L" || consoleInput == "l" {
		board.LoadGameList()
		return
	} else if consoleInput == "E" || consoleInput == "e" {
		fmt.Println("Ending game...")
		return
	}
	keyInput, _ := strconv.ParseInt(consoleInput, 10, 64)
	*key = int(keyInput)
	err := board.MakeMove(*key, player2)
	if err != nil {
		board.PrintBoard()
		fmt.Println(err)
	}
}
