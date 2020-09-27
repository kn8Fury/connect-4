package errors

import (
	"fmt"
)

// ColumnFull error for when the colomn is full
type ColumnFull struct {
	arg int
}

// Error string format
func (cf *ColumnFull) Error() string {
	return fmt.Sprintf("column is full - %d", cf.arg)
}

// NewColumnFull creates an ColumnFull error with given column
func NewColumnFull(column int) *ColumnFull {
	return &ColumnFull{arg: column}
}
