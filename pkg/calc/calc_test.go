package calc_test

import (
	"testing"

	"github.com/arbuzick57/calc_go/pkg/calc"
)

func TestCalc(t *testing.T) {
	tests := []struct {
		name           string
		expression     string
		expectedResult float64
		expectedError  error
	}{
		{
			name:           "Valid Expression",
			expression:     "5+2",
			expectedResult: 7,
			expectedError:  nil,
		},
		{
			name:           "Valid Expression With Brackets",
			expression:     "2*(2+2)",
			expectedResult: 8,
			expectedError:  nil,
		},
		{
			name:           "Incorrect Brackets",
			expression:     "(2+2",
			expectedResult: 0,
			expectedError:  calc.ErrBrackets,
		},
		{
			name:           "Empty Brackets",
			expression:     "5+()",
			expectedResult: 0,
			expectedError:  calc.ErrEmptyBrackets,
		},
		{
			name:           "Division By Zero",
			expression:     "5/0",
			expectedResult: 0,
			expectedError:  calc.ErrDivisionByZero,
		},
		{
			name:           "Not Number",
			expression:     "5+a",
			expectedResult: 0,
			expectedError:  calc.ErrNotNumber,
		},
		{
			name:           "Incorrect Placement Of Operation",
			expression:     "2+",
			expectedResult: 0,
			expectedError:  calc.ErrOperation,
		},
		{
			name:           "Empty Expression",
			expression:     "",
			expectedResult: 0,
			expectedError:  calc.ErrEmptyExpression,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := calc.Calc(test.expression)
			if err != test.expectedError {
				if test.expectedError == nil {
					t.Errorf("Expected no error, but got %s", err.Error())
				} else if err == nil {
					t.Errorf("Expected %s error, but got no error", test.expectedError.Error())
				} else {
					t.Errorf("Expected %s error, but got %s", test.expectedError, err.Error())
				}
			}
			if result != test.expectedResult {
				t.Errorf("Expected %f, but got %f", test.expectedResult, result)
			}
		})
	}
}
