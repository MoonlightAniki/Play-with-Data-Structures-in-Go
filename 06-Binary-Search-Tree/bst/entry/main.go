package main

import (
	"dev/Play-with-Data-Structures-in-Go/06-Binary-Search-Tree/bst"
	"math/rand"
	"fmt"
)

func main() {
	bst := bst.CreateBST()
	for i := 0; i < 10; i++ {
		e := rand.Int() % 100
		fmt.Println(e)
		bst.Add(e)
	}

	fmt.Println("PreOrder")
	bst.PreOreder()
	fmt.Println("InOrder")
	bst.InOrder()
	fmt.Println("PostOrder")
	bst.PostOrder()
	fmt.Println("LevelOrder")
	bst.LevelOrder()
	fmt.Println("Maximum:", bst.Maximum())
	fmt.Println("Minimum:", bst.Minimum())
	bst.RemoveMin()
	bst.RemoveMax()
	fmt.Println(bst.GetSize())
	bst.InOrder()
	fmt.Println("Remove")
	bst.Remove(51)
	bst.Remove(20)
	bst.InOrder()
}
