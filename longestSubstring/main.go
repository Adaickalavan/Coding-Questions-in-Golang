package main

import (
	"fmt"
)

//Returns the longest valid parentheses substring
func longestSubstr(arr []rune) (int, int) {
	start := 0
	end := 0
	stack := []int{}

	if arr == nil {
		return start, end
	}

	for ii, val := range arr {
		switch val {
		case '(':
			stack = append(stack, ii)
		default: //case ')':
			if len(stack) != 0 {
				end = ii
				stack = stack[:len(stack)-1] //pop opening bracket
			} else {
				start = stack[len(stack)-1]
			}

		}

	}

}

func main() {
	strList := []string{"", "(()", ")()())", "())((())", "())(()"}
	for _, str := range strList {
		input := []rune(str)
		start, end := longestSubstr(input)
		fmt.Println("The string", str, "is", string(input[start:end]))
	}
}
