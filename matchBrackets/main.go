package main

import (
	"fmt"
)

func main() {
	strList := []string{"", "()", "()[]{}", "{[]}", "[", "}{", "{()", "(]", "([)]"}
	for _, str := range strList {
		input := []rune(str)
		valid := matchBrackets(&[]rune{}, input)
		fmt.Println("The string", str, "is", valid)
	}
}

var bracketMap = map[rune]rune{
	rune(')'): rune('('),
	rune('}'): rune('{'),
	rune(']'): rune('['),
}

func matchBrackets(stack *[]rune, input []rune) bool {
	switch {
	case len(*stack) == 0 && len(input) == 0:
		return true
	case len(*stack) != 0 && len(input) == 0:
		return false
	case len(*stack) == 0 && len(input) != 0:
		switch _, ok := bracketMap[input[0]]; ok {
		case true: //closing bracket
			return false
		default: //opening bracket
			*stack = append(*stack, input[0])
			return matchBrackets(stack, input[1:])
		}
	default: //len(*stack) != 0 && len(input) != 0
		switch elem, ok := bracketMap[input[0]]; ok {
		case true: //closing bracket
			if elem == (*stack)[len(*stack)-1] {
				*stack = (*stack)[:len(*stack)-1]
				return matchBrackets(stack, input[1:])
			}
			return false
		default: //opening bracket
			*stack = append(*stack, input[0])
			return matchBrackets(stack, input[1:])
		}
	}
}
