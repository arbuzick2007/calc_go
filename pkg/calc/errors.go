package calc

import (
	"errors"
)

var (
	ErrBrackets        = errors.New("brackets don't match")
	ErrEmptyBrackets   = errors.New("empty brackets")
	ErrDivisionByZero  = errors.New("division by zero is forbidden")
	ErrNotNumber       = errors.New("not numbers")
	ErrOperation       = errors.New("operation is not between two expressions")
	ErrEmptyExpression = errors.New("empty expression")
)
