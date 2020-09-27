package errors

import (
	"fmt"

	"github.com/kn8Fury/connect4/game/colors"
)

// InvalidPlayer error for when it is not the player's turn
type InvalidPlayer struct {
	arg colors.Color
}

// Error string format
func (ip *InvalidPlayer) Error() string {
	return fmt.Sprintf("out of turn - %s", ip.arg)
}

// NewInvalidPlayer creates an InvalidPlayer error with given color
func NewInvalidPlayer(color colors.Color) *InvalidPlayer {
	return &InvalidPlayer{arg: color}
}
