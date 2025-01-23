package main

import (
	"strconv"
)

// func main() {
// 	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
// 	fmt.Println(calculate(" 2-1 + 2 "))
// }

func recurse(input string, index int) int {
	close := ')'
	open := '('
	runeArray := []rune(input)
	var number int = 0
	var sign int = 1
	var result = 0
	for index < len(input) {
		if isNumber(string(runeArray[index])) {
			num, _ := strconv.Atoi(string(runeArray[index]))
			number = number*10 + num
		} else if runeArray[index] == open {
			index++
			number = calculate(input)
		} else if runeArray[index] == close {
			result += number * sign
			return result
		} else if runeArray[index] == '+' {
			result += number * sign
			number = 0
			sign = 1
		} else if runeArray[index] == '-' {
			result += number * sign
			number = 0
			sign = -1
		}
		index++
	}
	if number != 0 {
		result += number * sign
	}
	return result
}

func calculate(input string) int {
	// index = 0
	return recurse(input, 0)
}

func isNumber(characer string) bool {
	if _, err := strconv.Atoi(characer); err != nil {
		return false
	}
	return true
}
