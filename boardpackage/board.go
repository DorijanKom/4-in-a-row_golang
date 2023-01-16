package boardpackage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Board struct {
	State          [][]string
	Rows           int
	Cols           int
	MovesPlayerOne []int
	MovesPlayerTwo []int
	Turn           int
}

const emptyField string = " "

func GenBoard(row, col int) *Board {

	// Creates new board struct
	board := &Board{
		State: make([][]string, row),
		Rows:  row,
		Cols:  col,
		Turn:  0,
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

func (board *Board) PrintBoard() {

	for i := 0; i < board.Rows; i++ {
		for j := 0; j < board.Cols; j++ {
			fmt.Printf(" [ %s  ] ", board.State[i][j])
		}
		fmt.Println()
	}

}

func (board *Board) MakeMove(key int, piece string) error {
	if key < 1 || key > board.Cols {
		return fmt.Errorf("invalid column")
	}

	for i := board.Rows - 1; i >= 0; i-- {
		if board.State[i][key-1] == emptyField {
			board.State[i][key-1] = piece
			return nil
		}
	}
	return fmt.Errorf("the column is full")
}

func (board *Board) EndGame(piece string) string {

	//horizontal check
	for j := 0; j < board.Cols-3; j++ {
		for i := 0; i < board.Rows; i++ {
			if board.State[i][j] == piece && board.State[i][j+1] == piece && board.State[i][j+2] == piece && board.State[i][j+3] == piece {
				return piece
			}
		}
	}
	//vertical check
	for i := 0; i < board.Rows-3; i++ {
		for j := 0; j < board.Cols; j++ {
			if board.State[i][j] == piece && board.State[i+1][j] == piece && board.State[i+2][j] == piece && board.State[i+3][j] == piece {
				return piece
			}
		}
	}
	//diagonal check left
	for i := 3; i < board.Cols; i++ {
		for j := 0; j < board.Rows-3; j++ {
			if board.State[i][j] == piece && board.State[i-1][j+1] == piece && board.State[i-2][j+2] == piece && board.State[i-3][j+3] == piece {
				return piece
			}
		}
	}
	//diagonal check right
	for i := 0; i < board.Cols; i++ {
		for j := 3; j < board.Rows-3; j++ {
			if board.State[i][j] == piece && board.State[i-1][j-1] == piece && board.State[i-2][j-2] == piece && board.State[i-3][j-3] == piece {
				return piece
			}
		}
	}

	// Check if draw
	for i := 0; i < board.Rows; i++ {
		for j := 0; j < board.Cols; j++ {
			if board.State[i][j] == emptyField {
				return "Draw"
			}
		}
	}
	return ""
}

func (board *Board) SaveGame(fileName string) error {
	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	save := strings.TrimSpace(fileName)

	file, err := os.Create(save)

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

	fmt.Println("Game saved!")
	board.PrintBoard()
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

func (board *Board) LoadGameList() {
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

	err = board.loadGame(files[input-1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Game loaded")
	board.PrintBoard()

}

func getSavedGames() ([]string, error) {
	files, err := filepath.Glob("*.json")
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (board *Board) PrintMoves() {
	fmt.Println()
	fmt.Printf("%v ", board.MovesPlayerOne)
	fmt.Println()
	fmt.Printf("%v ", board.MovesPlayerTwo)
	fmt.Println()
}
