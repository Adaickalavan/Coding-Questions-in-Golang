package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi")
	s := NewSet()
	s.Add("Hi")
	s.Print()

}

//Set implements the set
type Set map[interface{}]struct{}

//Add adds value to the set
func (s *Set) Add(key interface{}) {
	(*s)[key] = struct{}{}
}

//NewSet returns a new set
func NewSet() *Set {
	s := Set{}
	return &s
}

//Print prints the values within the set
func (s *Set)Print(){
	for key,val := range *s{
		fmt.Println("--",key,"--",val)
	}
}
