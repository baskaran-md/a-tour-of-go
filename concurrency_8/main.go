package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	ch := make(chan int)
	go Walk(tree.New(2), ch)
	fmt.Println("Sample Tree:")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Tree(2) vs Tree(2): ", Same(tree.New(2), tree.New(2)))
	fmt.Println("Tree(1) vs Tree(2): ", Same(tree.New(1), tree.New(2)))
}
