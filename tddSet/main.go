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

//NewSet returns a new set
func NewSet() *Set {
	s := Set{}
	return &s
}

//Print prints the keys within the set
func (s *Set) Print() {
	for key := range *s {
		fmt.Println("--", key)
	}
}

//Add adds key to the set
func (s *Set) Add(key interface{}) {
	(*s)[key] = struct{}{}
}

//Remove removes key from set
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

//Contains checks whether a key exists in set
func (s *Set) Contains(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}
