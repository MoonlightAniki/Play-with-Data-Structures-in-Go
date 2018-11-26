package Solution0203_V2

import (
	"fmt"
	"dev/Play-with-Data-Structures-in-Go/05-Recursion/leetcode/Kit"
)

// 203. Remove Linked List Elements
// https://leetcode.com/problems/remove-linked-list-elements/
/*
Remove all elements from a linked list of integers that have value val.

Example:
Input:  1->2->6->3->4->5->6, val = 6
Output: 1->2->3->4->5
 */

type ListNode = Kit.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return removeElements(head.Next, val)
	} else {
		head.Next = removeElements(head.Next, val)
		return head
	}
}

func Test0203() {
	nums := []int{1, 2, 3, 4, 5, 5, 6, 5}
	head := Kit.CreateLinkedList(nums)
	fmt.Println(head)
	head = removeElements(head, 5)
	fmt.Println(head)
}

//Runtime: 24 ms, faster than 7.23% of Go online submissions for Remove Linked List Elements.
