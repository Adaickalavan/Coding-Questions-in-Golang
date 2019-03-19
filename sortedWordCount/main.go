package main

import (
	"fmt"
	"log"
	"regexp"
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

// Less sorts words by count, in descending order followed by alphabetical order
func (words byCount) Less(i int, j int) bool {
	switch {
	case words[i].count > words[j].count:
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

	// Print the word count
	for _, elem := range list {
		fmt.Println(elem.word, elem.count)
	}

}

func clean(str string) []string {

	// Convert letters to lowercase and split string into words
	words := strings.Fields(strings.ToLower(str))

	// Remove all non letter characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	for ii, elem := range words {
		words[ii] = reg.ReplaceAllString(elem, "")
	}

	return words
}

func main() {
	// Example string
	str := `The quick brown fox jumped over the lazy dog's bowl.
	The dog was angry with the fox for considering him lazy.`

	// Clean up the string
	words := clean(str)

	// Sort the words
	sorter(words)
}
