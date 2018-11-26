package bst

import (
	"fmt"
	"dev/Play-with-Data-Structures-in-Go/04-Linked-List/queue"
)

type node struct {
	e     int
	left  *node
	right *node
}

type BST struct {
	root *node
	size int
}

func CreateBST() *BST {
	bst := &BST{}
	bst.root = nil
	bst.size = 0
	return bst
}

func (bst *BST) GetSize() int {
	return bst.size
}

func (bst *BST) IsEmpty() bool {
	return bst.size == 0
}

func (bst *BST) Add(e int) {
	bst.root = bst.add(bst.root, e)
}

func (bst *BST) add(root *node, e int) *node {
	if root == nil {
		bst.size++
		return &node{e: e}
	}
	if e < root.e {
		root.left = bst.add(root.left, e)
		return root
	} else if e > root.e {
		root.right = bst.add(root.right, e)
		return root
	} else {
		return root
	}
}

func (bst *BST) Contains(e int) bool {
	return bst.contains(bst.root, e)
}

func (bst *BST) contains(root *node, e int) bool {
	if root == nil {
		return false
	}
	if e == root.e {
		return true
	} else if e < root.e {
		return bst.contains(root.left, e)
	} else {
		return bst.contains(root.right, e)
	}
}

// 前序遍历
func (bst *BST) PreOreder() {
	bst.preorder(bst.root)
}

func (bst *BST) preorder(root *node) {
	if root == nil {
		return
	}
	fmt.Println(root.e)
	bst.preorder(root.left)
	bst.preorder(root.right)
}

func (bst *BST) InOrder() {
	bst.inorder(bst.root)
}

func (bst *BST) inorder(root *node) {
	if root == nil {
		return
	}
	bst.inorder(root.left)
	fmt.Println(root.e)
	bst.inorder(root.right)
}

func (bst *BST) PostOrder() {
	bst.postorder(bst.root)
}

func (bst *BST) postorder(root *node) {
	if root == nil {
		return
	}
	bst.postorder(root.left)
	bst.postorder(root.right)
	fmt.Println(root.e)
}

func (bst *BST) LevelOrder() {
	if bst.root == nil {
		return
	}
	q := queue.CreateLoopQueue()
	q.Offer(bst.root)
	for !q.IsEmpty() {
		front := q.Poll().(*node)
		fmt.Println(front.e)
		if front.left != nil {
			q.Offer(front.left)
		}
		if front.right != nil {
			q.Offer(front.right)
		}
	}
}

func (bst *BST) Minimum() int {
	if bst.size == 0 {
		panic("BST is empty")
	}
	return bst.minimum(bst.root).e
}

func (bst *BST) minimum(root *node) *node {
	if root.left == nil {
		return root
	}
	return bst.minimum(root.left)
}

func (bst *BST) Maximum() int {
	if bst.size == 0 {
		panic("BST is empty")
	}
	return bst.maximum(bst.root).e
}

func (bst *BST) maximum(root *node) *node {
	if root.right == nil {
		return root
	}
	return bst.maximum(root.right)
}

func (bst *BST) RemoveMin() int {
	if bst.size == 0 {
		panic("BST is empty")
	}
	ret := bst.Minimum()
	bst.root = bst.removeMin(bst.root)
	return ret
}

func (bst *BST) removeMin(root *node) *node {
	if root.left == nil {
		bst.size--
		return root.right
	}
	root.left = bst.removeMin(root.left)
	return root
}

func (bst *BST) RemoveMax() int {
	if bst.size == 0 {
		panic("BST is empty")
	}
	ret := bst.Maximum()
	bst.root = bst.removeMax(bst.root)
	return ret
}

func (bst *BST) removeMax(root *node) *node {
	if root.right == nil {
		bst.size--
		return root.left
	}
	root.right = bst.removeMax(root.right)
	return root
}

func (bst *BST) Remove(e int) {
	bst.root = bst.remove(bst.root, e)
}

func (bst *BST) remove(root *node, e int) *node {
	if root == nil {
		panic(fmt.Sprintf("Remove failed, element:%d doesn't exist", e))
	}
	var ret *node
	if e < root.e {
		root.left = bst.remove(root.left, e)
		ret = root
	} else if e > root.e {
		root.right = bst.remove(root.right, e)
		ret = root
	} else { // e == root.e
		if root.left == nil {
			bst.size--
			ret = root.right
		} else if root.right == nil {
			bst.size--
			ret = root.left
		} else { // root.left != nil && root.right != nil
			successor := &node{}
			successor.e = bst.Minimum()
			successor.left = root.left
			successor.right = bst.remove(root.right, successor.e)
			root.left = nil
			root.right = nil
			ret = successor
		}
	}
	return ret
}
