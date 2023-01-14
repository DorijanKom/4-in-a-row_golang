package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Board struct {
	State          [][]string
	Rows           int
	Cols           int
	MovesPlayerOne []int
	MovesPlayerTwo []int
}

const (
	emptyField string = " "
	player1    string = "◯"
	player2    string = "⬤"
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
			board := genBoard(row, col)

			board.printBoard()
			fmt.Println("input")
			for {
				if turn%2 == 0 { // if turn counter is even then it's player1's turn, else it's player2's turn
					fmt.Scan(&key)                      // takes user input
					err := board.makeMove(key, player1) //cals make move function and saves response into the err variable
					if err != nil {                     // checks for errors and handles them accordingly
						board.printBoard()
						fmt.Println(err)
						continue
					}
					board.printBoard()
					turn++
					moveHistory(key, &board.MovesPlayerOne)
					fmt.Printf("%v ", board.MovesPlayerOne)
					fmt.Println()
					fmt.Printf("%v ", board.MovesPlayerTwo)

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

func genBoard(row, col int) Board {

	// Creates new board struct
	board := Board{
		State: make([][]string, row),
		Rows:  row,
		Cols:  col,
	}

	// Generates the 2d array (state)
	for i := range board.State {
		board.State[i] = make([]string, col)
		for j := range board.State[i] {
			board.State[i][j] = emptyField
		}
	}

	return board
}

func (board *Board) printBoard() {

	for i := board.Rows - 1; i >= 0; i-- {
		for j := 0; j < board.Cols; j++ {
			fmt.Printf(" [ %s  ] ", board.State[i][j])
		}
		fmt.Println()
	}

}

func (board *Board) makeMove(key int, piece string) error {
	if key < 0 || key-1 >= board.Cols {
		return fmt.Errorf("invalid column")
	}
	if board.State[board.Rows-1][key-1] != emptyField {
		return fmt.Errorf("the column is full")
	}

	for i := 0; i < board.Rows; i++ {
		if board.State[i][key-1] == emptyField {
			board.State[i][key-1] = piece
			break
		}
	}

	return nil

}

func (board *Board) endGame(piece string) bool {

	//horizontal check
	for j := 0; j < board.Rows-3; j++ {
		for i := 0; i < board.Rows; i++ {
			if board.State[i][j] == piece && board.State[i][j+1] == piece && board.State[i][j+2] == piece && board.State[i][j+3] == piece {
				fmt.Println("Victory")
			}
		}
	}
	//vertical check
	for i := 0; i < board.Cols-3; i++ {
		for j := 0; j < board.Rows; j++ {
			if board.State[i][j] == piece && board.State[i][j+1] == piece && board.State[i][j+2] == piece && board.State[i][j+3] == piece {
				fmt.Println("Victory")
			}
		}
	}
	//diagonal check left
	for i := 3; i < board.Cols; i++ {
		for j := 0; j < board.Rows-3; j++ {
			if board.State[i][j] == piece && board.State[i-1][j+1] == piece && board.State[i-2][j+2] == piece && board.State[i-3][j+3] == piece {
				fmt.Println("Victory")
			}
		}
	}
	//diagonal check right
	for i := 0; i < board.Cols; i++ {
		for j := 3; j < board.Rows-3; j++ {
			if board.State[i][j] == piece && board.State[i-1][j-1] == piece && board.State[i-2][j-2] == piece && board.State[i-3][j-3] == piece {
				fmt.Println("Victory")
			}
		}
	}
	return false
}

func moveHistory(key int, moves *[]int) {

	*moves = append(*moves, key)
}

func checkRowCol(row, col int) bool {
	return col-row == 2 || col-row <= 2 && col > 6 && row > 5
}

func (board *Board) saveGame(fileName string) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	data, err := json.MarshalIndent(board, "", "   ")
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (board *Board) loadGame(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)

	err = decoder.Decode(board)
	if err != nil {
		return err
	}

	return nil
}

func loadGameList() {
	files, err := getSavedGames()
	if err != nil {
		fmt.Println(err)
	}

	if len(files) == 0 {
		fmt.Println("There aren't any saved games.")
		return
	}

	fmt.Println("Select a game to load: ")
	for i, file := range files {
		fmt.Printf("%d. %s\n", i+1, file)
	}

	var input int
	fmt.Scan(&input)

	if input < 1 || input > len(files) {
		fmt.Println("Invalid selection!")
		return
	}

	board := &Board{}
	err = board.loadGame(files[input-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Game loaded")
	board.printBoard()

}

func getSavedGames() ([]string, error) {
	files, err := filepath.Glob("*.json")
	if err != nil {
		return nil, err
	}

	return files, nil
}
