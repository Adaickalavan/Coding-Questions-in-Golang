package main

import "fmt"

var m = make(map[string]bool)

type node struct {
	next []*node
	end  bool
	name string
}

type queue []*node

// depth first search
func dfs(curNode *node) bool {
	if _, ok := m[curNode.name]; ok {
		return false
	}

	fmt.Print(curNode.name, " ")

	if curNode.end {
		return true
	}

	for _, newNode := range curNode.next {
		if dfs(newNode) {
			return true
		}
	}

	return false
}

// breadth first search
func bfs(q *queue) bool {
	if len(*q) == 0 {
		return false
	}

	curNode := (*q)[0]

	if _, ok := m[curNode.name]; ok {
		return false
	}

	fmt.Print(curNode.name, " ")

	if curNode.end {
		return true
	}

	*q = append((*q)[1:], curNode.next...)

	if bfs(q) {
		return true
	}

	return false
}

func main() {

	node1 := &node{end: false, name: "1"}
	node2 := &node{end: false, name: "2"}
	node3 := &node{end: false, name: "3"}
	node4 := &node{end: false, name: "4"}
	node5 := &node{end: false, name: "5"}
	node6 := &node{end: false, name: "6"}
	node7 := &node{end: false, name: "7"}
	node8 := &node{end: true, name: "8"}

	nexta := []*node{node2, node3}
	nextb := []*node{node4, node5}
	nextc := []*node{node6}
	nextd := []*node{node7, node8}

	node1.next = nexta
	node2.next = nextb
	node3.next = nextc
	node5.next = nextd

	// Tree structure
	//               node1
	// 				   |
	//       ---------------------
	// 		 |                   |
	//     node2               node 3
	//       |                   |
	//   -----------             |
	//   |         |             |
	// node4     node5         node 6
	//             |
	//        -----------
	//        |         |
	//      node7     node8

	// Breadth first search
	fmt.Print("Breath first search = ")
	n := queue([]*node{node1})
	tree := bfs(&n)
	fmt.Println(tree)

	// Reset map
	m = make(map[string]bool)

	// Depth first search
	fmt.Print("Depth first search = ")
	tree = dfs(node1)
	fmt.Println(tree)
}
