package main

import (
	"fmt"
)

func main() {
	// Input data
	sequence := "ATGCATTGC"
	subSequence := "CTG"
	sequenceLen := len(sequence)
	subSequenceLen := len(subSequence)

	var truncatedSequence string
	m := make(map[int]float64)
	var c = make(chan res)

	// Obtain substring and perform comparison
	for ii := 0; ii < sequenceLen; ii++ {
		if ii+subSequenceLen > sequenceLen {
			truncatedSequence = sequence[ii:sequenceLen]
		} else {
			truncatedSequence = sequence[ii : ii+subSequenceLen]
		}
		// Start goroutine for comparison
		go comp(truncatedSequence, subSequence, ii, c)
	}

	// Retrieve results from each comparison
	for ii := 0; ii < sequenceLen; ii++ {
		r := <-c
		m[r.key] = r.val
	}

	// Print results
	fmt.Println(m)
}

type res struct {
	key int
	val float64
}

// comp compares sequence and subSequence, letter-by-letter, returning the percentage match
func comp(sequence string, subSequence string, key int, c chan res) {
	sequenceLen := len(sequence)
	subSequenceLen := len(subSequence)
	var count int
	for jj := 0; jj < sequenceLen; jj++ {
		if sequence[jj] == subSequence[jj] {
			count++
		}
	}
	// Return result
	c <- res{key: key, val: float64(count) / float64(subSequenceLen)}
}