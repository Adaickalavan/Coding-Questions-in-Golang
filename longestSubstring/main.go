package main

import (
	"fmt"
)

//Returns the longest valid parentheses substring
func longestSubstr(arr []rune) (int, int) {
	start := 0
	end := 0
	longestStart := 0
	longestEnd := 0
	stack := []int{-1}

	if len(arr) <= 1 {
		return longestStart, longestEnd
	}

	for ii, val := range arr {
		switch val {
		case rune('('):
			stack = append(stack, ii)
		default: //case ')':
			stack = stack[:len(stack)-1] //pop opening bracket
			if len(stack) != 0 {
				start = stack[len(stack)-1]
				end = ii
				if end-start > longestEnd-longestStart {
					longestStart = start
					longestEnd = end
				}
			} else { //empty stack
				stack = append(stack, ii)
			}
		}
	}
	return longestStart + 1, longestEnd
}

func main() {
	strList := []string{"", "(()", ")()())", "())((())", "())(()"}
	for _, str := range strList {
		input := []rune(str)
		start, end := longestSubstr(input)
		fmt.Println("Longest valid substring of", str, "is", string(input[start:end+1]), ", from ", start, "to ", end)
	}
}
