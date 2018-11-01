// https://www.dropbox.com/s/incy1xuj1j3f1wt/Q1.pdf?dl=0

package main

import "fmt"

//Adder is a class
type Adder struct {
	arr    []int
	cumArr []int
}

//Run time complexity of range sum is O(1)
//rangeSum sums integers between two indices
func (adder *Adder) rangeSum(i int, j int) int {
	if i == 0 {
		return adder.cumArr[j]
	}
	return adder.cumArr[j] - adder.cumArr[i-1]
}

//newAdder is a Adder class constructor
func newAdder(array []int) *Adder {
	return &Adder{
		arr:    array,
		cumArr: cumSum(array),
	}
}

//cumSum returns a cumulative sum array
func cumSum(array []int) []int {
	var output = []int{}
	var sum = 0
	for _, val := range array {
		sum = sum + val
		output = append(output, sum)
	}
	return output
}

func main() {
	A := []int{-2, 0, 3, -5, 2, -1}

	adder := newAdder(A)

	sum := adder.rangeSum(0, 2) // answer = 1
	fmt.Println(sum)

	sum = adder.rangeSum(2, 5) // answer = -1
	fmt.Println(sum)

	sum = adder.rangeSum(0, 5) // answer =  -3
	fmt.Println(sum)
}
