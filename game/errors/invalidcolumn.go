package errors

import (
	"fmt"

	"github.com/kn8Fury/connect4/game/constants"
)

// InvalidColumn error for when column is out of range
type InvalidColumn struct {
	arg int
}

// Error string format
func (ic *InvalidColumn) Error() string {
	return fmt.Sprintf("%d out of range [0,%d)", ic.arg, constants.MaxColumns)
}

// NewInvalidColumn creates an InvalidColumn error with given column
func NewInvalidColumn(column int) *InvalidColumn {
	return &InvalidColumn{arg: column}
}
