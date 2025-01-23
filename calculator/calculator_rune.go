package calculator

import (
	"unicode"
)

func calculate(s string) int {
	return calculatorFunc(s, new(int))
}

func calculatorFunc(s string, currentIndex *int) int {

	runes := []rune(s)
	currentNumber := 0
	result := 0
	previousSign := '+'
	for index := *currentIndex; index < len(runes); index++ {
		char := runes[index]
		if unicode.IsDigit(char) {
			currentNumber = currentNumber*10 + int(char-'0')
		}

		if char == '+' || char == '-' || char == '*' || char == '/' || index == len(runes)-1 || char == ')' {
			if previousSign == '+' {
				result = result + currentNumber
			}
			if previousSign == '-' {
				result = result - currentNumber
			}
			if previousSign == '*' {
				result = result * currentNumber
			}
			if previousSign == '/' {
				result = result / currentNumber
			}
			if char == ')' {
				*currentIndex = index
				return result
			}
			previousSign = char
			currentNumber = 0
		} else if char == '(' {
			skipIndex := index + 1
			currentNumber = calculatorFunc(s, &skipIndex)
			index = skipIndex
		} else {
			continue
		}
	}

	if currentNumber != 0 {
		if previousSign == '+' {
			result = result + currentNumber
		}
		if previousSign == '-' {
			result = result - currentNumber
		}
		if previousSign == '*' {
			result = result * currentNumber
		}
		if previousSign == '/' {
			result = result / currentNumber
		}
	}
	return result
}
