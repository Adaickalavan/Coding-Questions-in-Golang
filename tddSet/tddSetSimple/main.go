package main

import (
	"fmt"
)

func main() {

	fmt.Println("Hi")
	s := NewSet()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(1)
	fmt.Println(s.Contains(1))
	fmt.Println(s.Contains(2))
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
	fmt.Print("Contents of Set: ")
	for key := range *s {
		fmt.Print(key, ", ")
	}
}

//Add adds key to the set
func (s *Set) Add(key interface{}) {
	(*s)[key] = struct{}{}
}

//Remove removes key from the set
func (s *Set) Remove(key interface{}) {
	delete((*s), key)
}

//Contains checks whether a key exists in set
func (s *Set) Contains(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}

//Contains checks whether a key exists in set
func (s *Set) Union(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}
