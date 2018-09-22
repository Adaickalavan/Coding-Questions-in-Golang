package main

import (
	"fmt"
)

func evenFib() func() float64 {
	x := float64(0)
	y := float64(1)
	return func() float64 {
		x, y = y, x+y
		return y
	}
}

func main() {
	f := evenFib()
	for ii := 0; ii < 5; ii++ {
		result := f()
		fmt.Println(result)
	}
}
