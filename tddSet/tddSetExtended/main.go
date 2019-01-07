package main

import (
	"fmt"
)

func main() {

	var s Set
	s = NewSetInt()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	s.Remove(1)
	fmt.Println(s.Contains(1))
	fmt.Println(s.Contains(2))
	// s.Print()
}

//Set implements the set
type Set interface {
	Add(interface{})
	Remove(interface{})
	Contains(interface{}) bool
}

//SetBase returns map of type Int
type SetBase map[interface{}]struct{}

//NewSetBase returns a new set
func NewSetBase() *SetBase {
	s := SetBase{}
	return &s
}

//Print prints the keys within the set
func (s *SetBase) Print() {
	fmt.Print("Contents of Set: ")
	for key := range *s {
		fmt.Print(key, ", ")
	}
}

//Add adds key to the set
func (s *SetBase) Add(key interface{}) {
	(*s)[key] = struct{}{}
}

//Remove removes key from the set
func (s *SetBase) Remove(key interface{}) {
	delete((*s), key)
}

//Contains checks whether a key exists in set
func (s *SetBase) Contains(key interface{}) bool {
	_, ok := (*s)[key]
	return ok
}



//SetInt returns map of type Int
type SetInt struct {
	SetBase
}

//NewSetInt returns a new set
func NewSetInt() *SetInt {
	s := make(map[int]struct{})
	return &s
}
