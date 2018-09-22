package main

import (
	"fmt"
)

//Recursively merge two ordered list
func arrange(arr1 []int, arr2 []int, output *[]int) {
	switch {
	case len(arr1) == 0 && len(arr2) == 0:
		return
	case len(arr1) == 0 && len(arr2) > 0:
		*output = append(*output, arr2...)
		return
	case len(arr2) == 0 && len(arr1) > 0:
		*output = append(*output, arr1...)
		return
	default:
		if arr1[0] <= arr2[0] {
			*output = append(*output, arr1[0])
			arrange(arr1[1:], arr2, output)
		} else {
			*output = append(*output, arr2[0])
			arrange(arr1, arr2[1:], output)
		}
	}
}

func main() {
	arr1 := []int{1, 4, 7}
	arr2 := []int{3, 4, 5}
	output := []int{}
	arrange(arr1, arr2, &output)
	fmt.Println(output)
}
