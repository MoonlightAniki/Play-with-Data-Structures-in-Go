package Kit

import (
	"bytes"
	"fmt"
)

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func CreateLinkedList(nums []int) *ListNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	dummyHead := &ListNode{-1, nil}
	prev := dummyHead
	for _, v := range nums {
		prev.Next = &ListNode{v, nil}
		prev = prev.Next
	}
	ret := dummyHead.Next
	dummyHead.Next = nil
	return ret
}

func (head *ListNode) String() string {
	var buffer bytes.Buffer
	for cur := head; cur != nil; cur = cur.Next {
		buffer.WriteString(fmt.Sprintf("%d->", cur.Val))
	}
	buffer.WriteString("nil")
	return buffer.String()
}
