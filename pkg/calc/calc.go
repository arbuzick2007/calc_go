package calc

import (
	"errors"
	"unicode"
)

var (
	ErrBrackets        = errors.New("brackets don't match")
	ErrEmptyBrackets   = errors.New("empty brackets")
	ErrDivisionByZero  = errors.New("division by zero is forbidden")
	ErrNotNumber       = errors.New("not numbers")
	ErrOperation       = errors.New("operation is not between two expressions")
	ErrEmptyExpression = errors.New("empty expression")
)

func isOperation(symb rune) bool {
	return symb == '+' || symb == '-' || symb == '*' || symb == '/'
}

func getPriority(symb rune) int {
	if symb == '*' || symb == '/' {
		return 2
	}
	return 1
}

func CheckExpression(expressionStr string) error {
	expression := []rune(expressionStr)
	if len(expression) == 0 {
		return ErrEmptyExpression
	}
	bal := 0
	for _, symb := range expression {
		if symb == '(' {
			bal++
		} else if symb == ')' {
			if bal == 0 {
				return ErrBrackets
			}
			bal--
		}
	}
	if bal != 0 {
		return ErrBrackets
	}
	for ind, symb := range expression {
		if isOperation(symb) {
			if ind == 0 || ind+1 == len(expression) || (isOperation(expression[ind-1]) || expression[ind-1] == '(') || (isOperation(expression[ind+1]) || expression[ind+1] == ')') {
				return ErrOperation
			}
		}
		if symb == '(' && expression[ind+1] == ')' {
			return ErrEmptyBrackets
		}
		if symb != '(' && symb != ')' && !isOperation(symb) && !unicode.IsDigit(symb) {
			return ErrNotNumber
		}
	}
	return nil
}

func Calc(expressionStr string) (float64, error) {
	expression := []rune(expressionStr)
	ord := make([]rune, 0)
	stack := make([]rune, 0)
	for _, symb := range expression {
		if symb == '(' {
			stack = append(stack, symb)
		} else if symb == ')' {
			for stack[len(stack)-1] != '(' {
				ord = append(ord, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
		} else if unicode.IsDigit(symb) {
			ord = append(ord, symb)
		} else if isOperation(symb) {
			for len(stack) > 0 && isOperation(stack[len(stack)-1]) && getPriority(stack[len(stack)-1]) >= getPriority(symb) {
				ord = append(ord, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, symb)
		} else {
			return 0, ErrNotNumber
		}
	}
	for len(stack) > 0 {
		ord = append(ord, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	stackNumbers := make([]float64, 0)
	for _, symb := range ord {
		if unicode.IsDigit(symb) {
			stackNumbers = append(stackNumbers, float64(symb-'0'))
		} else {
			secondNumber := stackNumbers[len(stackNumbers)-1]
			stackNumbers = stackNumbers[:len(stackNumbers)-1]
			firstNumber := stackNumbers[len(stackNumbers)-1]
			stackNumbers = stackNumbers[:len(stackNumbers)-1]
			var resNumber float64
			if symb == '+' {
				resNumber = firstNumber + secondNumber
			} else if symb == '-' {
				resNumber = firstNumber - secondNumber
			} else if symb == '*' {
				resNumber = firstNumber * secondNumber
			} else if symb == '/' {
				if secondNumber == 0 {
					return 0, ErrDivisionByZero
				}
				resNumber = firstNumber / secondNumber
			}
			stackNumbers = append(stackNumbers, resNumber)
		}
	}
	return stackNumbers[0], nil
}
