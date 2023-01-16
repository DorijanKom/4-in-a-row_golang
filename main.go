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
					reader.Scan() // takes user input
					consoleInput := reader.Text()
					if consoleInput == "S" || consoleInput == "s" {
						reader.Scan()
						filename := reader.Text()
						board.SaveGame(filename)
						continue
					} else if consoleInput == "L" || consoleInput == "l" {
						board.LoadGameList()
						continue
					} else if consoleInput == "E" || consoleInput == "e" {
						break
					}
					keyInput, _ := strconv.ParseInt(consoleInput, 10, 64)
					key = int(keyInput)
					err := board.MakeMove(key, player1) //cals make move function and saves response into the err variable
					if err != nil {                     // checks for errors and handles them accordingly
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerOne)
					board.PrintMoves()
					if board.EndGame(player1) == player1 {
						fmt.Printf("Player 1 %s Victory!!!\n", player1)
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}

					} else if board.EndGame(player1) == "Draw" {
						fmt.Println("The game is a draw")
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}
					}
				} else {
					reader.Scan()
					consoleInput := reader.Text()

					if consoleInput == "S" || consoleInput == "s" {
						reader.Scan()
						filename := reader.Text()
						board.SaveGame(filename)
						continue
					} else if consoleInput == "L" || consoleInput == "l" {
						board.LoadGameList()
						continue
					} else if consoleInput == "E" || consoleInput == "e" {
						break
					}
					keyInput, _ := strconv.ParseInt(consoleInput, 10, 64)
					key = int(keyInput)
					err := board.MakeMove(key, player2)
					if err != nil {
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerTwo)
					board.PrintMoves()
					if board.EndGame(player2) == player2 {
						fmt.Printf("Player 2 %s Victory!!!\n", player2)
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}

					} else if board.EndGame(player2) == "Draw" {
						fmt.Println("The game is a draw")
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}
					}
				}
			}
		}
		if userInput == 2 {
			board := boardpackage.GenBoard(6, 7)
			board.LoadGameList()
			for {
				if board.Turn%2 == 0 { // if turn counter is even then it's player1's turn, else it's player2's turn
					reader.Scan() // takes user input
					consoleInput := reader.Text()
					if consoleInput == "S" || consoleInput == "s" {
						reader.Scan()
						filename := reader.Text()
						board.SaveGame(filename)
						continue
					} else if consoleInput == "L" || consoleInput == "l" {
						board.LoadGameList()
						continue
					} else if consoleInput == "E" || consoleInput == "e" {
						break
					}
					keyInput, _ := strconv.ParseInt(consoleInput, 10, 64)
					key = int(keyInput)
					err := board.MakeMove(key, player1) //cals make move function and saves response into the err variable
					if err != nil {                     // checks for errors and handles them accordingly
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerOne)
					board.PrintMoves()
					if board.EndGame(player1) == player1 {
						fmt.Printf("Player 1 %s Victory!!!\n", player1)
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}

					} else if board.EndGame(player1) == "Draw" {
						fmt.Println("The game is a draw")
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}
					}

				} else {
					reader.Scan()
					consoleInput := reader.Text()

					if consoleInput == "S" || consoleInput == "s" {
						reader.Scan()
						filename := reader.Text()
						board.SaveGame(filename)
						continue
					} else if consoleInput == "L" || consoleInput == "l" {
						board.LoadGameList()
						continue
					} else if consoleInput == "E" || consoleInput == "e" {
						break
					}
					keyInput, _ := strconv.ParseInt(consoleInput, 10, 64)
					key = int(keyInput)
					err := board.MakeMove(key, player2)
					if err != nil {
						board.PrintBoard()
						fmt.Println(err)
						continue
					}
					board.PrintBoard()
					board.Turn++
					moveHistory(key, &board.MovesPlayerTwo)
					board.PrintMoves()
					if board.EndGame(player2) == player2 {
						fmt.Printf("Player 2 %s Victory!!!\n", player2)
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}

					} else if board.EndGame(player2) == "Draw" {
						fmt.Println("The game is a draw")
						fmt.Println("Play again? Y/N")
						reader.Scan()
						response := reader.Text()
						if response == "Y" || response == "y" {
							break
						} else if response == "N" || response == "n" {
							os.Exit(1)
						}
					}
				}
			}
		}
		if userInput == 3 {
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
