package main

import (
	"fmt"
	"sort"
	"strings"
)

type byCount []wordCount

func (words byCount) Len() int {
	return len(words)
}

func (words byCount) Swap(i int, j int) {
	words[i], words[j] = words[j], words[i]
}

func (words byCount) Less(i int, j int) bool {
	switch {
	case words[i].count < words[j].count:
		return true
	case words[i].count == words[j].count && words[i].word < words[j].word:
		return true
	default:
		return false
	}
}

type wordCount struct {
	word  string
	count int
}

func sorter(words []string) {
	m := make(map[string]int)
	var list []wordCount

	for _, word := range words {
		m[word] = m[word] + 1
	}

	for key, val := range m {
		list = append(list, wordCount{word: key, count: val})
	}

	sort.Sort(byCount(list))

	fmt.Println(list)

}

func main() {

	str := "zxgh zxgh abcdefgh! hi jUOn  ABC abc DEF def DEf ki hi"

	// Convert letters to lowercase and split string into words
	words := strings.Fields(strings.ToLower(str))
	// Sort string
	sorter(words)

}
