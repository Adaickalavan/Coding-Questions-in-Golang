package main

import "fmt"

//Reverse a string in parallel
func reverser(arr1 []rune, length int, c0 chan int) {
	c1 := make(chan int)
	switch {
	case length == 0:
		c0 <- 1
		return
	case length == 1:
		c0 <- 1
		return
	case length >= 2:
		go reverser(arr1[1:length-1], length-2, c1)
		arr1[0], arr1[length-1] = arr1[length-1], arr1[0]
		<-c1
		c0 <- 1
	}
}

func main() {

	str := "abcdefgh"
	str1 := []rune(str)

	ch := make(chan int)

	go reverser(str1, len(str1), ch)
	<-ch

	fmt.Print(string(str1))
}
