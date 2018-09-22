package main

import (
	"fmt"
)

//Node contains value and pointer to next node
type node struct {
	val      int
	nextNode *node
}

//Deletes node given by pointer
func delete(ptr *node) {
	ptrNext := ptr.nextNode
	ptr.val = ptrNext.val
	ptr.nextNode = ptrNext.nextNode
}

func main() {

	sixthNode := &node{val: 6, nextNode: nil}
	fifthNode := &node{val: 5, nextNode: sixthNode}
	fourthNode := &node{val: 4, nextNode: fifthNode}
	thirdNode := &node{val: 3, nextNode: fourthNode}
	secondNode := &node{val: 2, nextNode: thirdNode}
	firstNode := &node{val: 1, nextNode: secondNode}

	delete(fourthNode)

	for ii := firstNode; ii != nil; ii = ii.nextNode {
		fmt.Println(ii)
	}
}
