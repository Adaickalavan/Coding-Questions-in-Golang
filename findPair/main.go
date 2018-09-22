package main

import (
	"fmt"
)

func findPair(arr1 []int, arr2 []int, n int) [][]int {
	hm := make(map[int]int)
	output := [][]int{}

	for _, val := range arr1 {
		hm[val] = hm[val] + 1
	}
	fmt.Print(hm)

	for _, val := range arr2 {
		temp := n - val
		if count, ok := hm[temp]; ok {
			for jj := count; jj > 0; jj-- {
				output = append(output, []int{temp, val})
			}
		}
	}

	return output
}

func main() {

	arr1 := []int{-1, -2, 4, 4, -2, -6, 5, -7}
	arr2 := []int{0, 6, 3, 4, 0}
	output := findPair(arr1, arr2, 4)
	fmt.Print(output)

}
