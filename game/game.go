package game

import (
	"strings"

	"github.com/kn8Fury/connect4/game/colors"
	"github.com/kn8Fury/connect4/game/constants"
	"github.com/kn8Fury/connect4/game/errors"
	"github.com/kn8Fury/connect4/game/results"
)

// Board is the structure of a game board
type Board struct {
	matrix        [constants.MaxColumns][constants.MaxRows](colors.Color)
	currentPlayer colors.Color
}

// NewBoard creates an instance of a new game board
func NewBoard() Board {
	return Board{currentPlayer: colors.Yellow}
}

// CurrentPlayer gets the color of the current player
func (gb *Board) CurrentPlayer() colors.Color {
	return gb.currentPlayer
}

// Draw returns a string representing the boards current state
func (gb *Board) Draw() string {
	var sb strings.Builder
	for r := constants.MaxRows - 1; r >= 0; r-- {
		sb.WriteByte('|')
		for c := 0; c < constants.MaxColumns; c++ {
			switch gb.matrix[c][r] {
			case colors.None:
				sb.WriteByte('O')
			case colors.Yellow:
				sb.WriteByte('Y')
			case colors.Red:
				sb.WriteByte('R')
			}
			sb.WriteByte('|')
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// Drop drops a coin of color 'color' into 'column'.
// returns true on success and false otherwise
func (gb *Board) Drop(color colors.Color, column int) (results.Result, error) {
	if color != gb.currentPlayer {
		return results.Invalid, errors.NewInvalidPlayer(color)
	}
	if column < 0 || constants.MaxColumns <= column {
		return results.Invalid, errors.NewInvalidColumn(column)
	}
	if gb.matrix[column][constants.MaxRows-1] != 0 {
		return results.Invalid, errors.NewColumnFull(column)
	}
	var row int
	for ; row < constants.MaxRows; row++ {
		if gb.matrix[column][row] == 0 {
			gb.matrix[column][row] = color
			break
		}
	}
	// fmt.Printf("c: %d, r: %d\n", column, row)
	return gb.check(column, row, color), nil
}

func (gb *Board) check(column, row int, color colors.Color) results.Result {
	// check south
	// fmt.Println("checking south")
	var c, r int
	count := 1
	c = column - 1
	for c >= 0 && c > column-4 {
		if gb.matrix[c][row] != color {
			break
		}
		count++
		c--
	}
	if count == 4 {
		return results.Win
	}
	// check south-east
	// fmt.Println("checking south-east")
	count, c, r = 1, column-1, row+1
	for c >= 0 && c > column-4 && r < constants.MaxRows && r < row+4 {
		if gb.matrix[c][r] != color {
			break
		}
		count++
		c--
		r++
	}
	if count == 4 {
		return results.Win
	}
	// check east
	// fmt.Println("checking east")
	count, r = 1, row+1
	for r < constants.MaxRows && r < row+4 {
		if gb.matrix[column][r] != color {
			break
		}
		count++
		r++
	}
	if count == 4 {
		return results.Win
	}
	// check north-east
	// fmt.Println("checking north-east")
	count, c, r = 1, column+1, row+1
	for c < constants.MaxColumns && c < column+4 && r < constants.MaxRows && r < row+4 {
		if gb.matrix[c][r] != color {
			break
		}
		count++
		c++
		r++
	}
	if count == 4 {
		return results.Win
	}
	// check north
	// fmt.Println("checking north")
	count, c = 1, column+1
	for c < constants.MaxColumns && c < column+4 {
		if gb.matrix[c][row] != color {
			break
		}
		count++
		c++
	}
	if count == 4 {
		return results.Win
	}
	// check north-west
	// fmt.Println("checking north-west")
	count, c, r = 1, column+1, row-1
	for c < constants.MaxColumns && c < column+4 && r >= 0 && r > row-4 {
		if gb.matrix[c][r] != color {
			break
		}
		count++
		c++
		r--
	}
	if count == 4 {
		return results.Win
	}
	// check west
	// fmt.Println("checking west")
	count, r = 1, row-1
	for r >= 0 && r > row-4 {
		if gb.matrix[column][r] != color {
			break
		}
		count++
		r--
	}
	if count == 4 {
		return results.Win
	}
	// check south-west
	// fmt.Println("checking south-west")
	count, c, r = 1, column-1, row-1
	for c >= 0 && c > column-4 && r >= 0 && r > row-4 {
		if gb.matrix[c][r] != color {
			break
		}
		count++
		c--
		r--
	}
	if count == 4 {
		return results.Win
	}
	// check if board is full
	// fmt.Println("checking if board is full")
	for c = 0; c < constants.MaxColumns; c++ {
		if gb.matrix[c][constants.MaxRows-1] == colors.None {
			break
		}
	}
	if c == constants.MaxColumns {
		return results.Draw
	}
	if gb.currentPlayer == colors.Yellow {
		gb.currentPlayer = colors.Red
	} else {
		gb.currentPlayer = colors.Yellow
	}
	return results.Continue
}
