package list

import (
	"bytes"
	"fmt"
)

type node struct {
	val  interface{}
	next *node
}

type LinkedList struct {
	head *node
	size int
}

func CreateLinkedList() *LinkedList {
	list := &LinkedList{}
	list.head = nil
	list.size = 0
	return list
}

func (list *LinkedList) GetSize() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

func (list *LinkedList) Add(index int, val interface{}) {
	if index < 0 || index > list.size {
		panic("Add failed, require index >= 0 && index <= list.size")
	}
	list.head = list.add(list.head, index, val)
}

func (list *LinkedList) add(head *node, index int, val interface{}) *node {
	if index == 0 {
		newHead := &node{val, head}
		head = newHead
		list.size++
		return head
	}
	head.next = list.add(head.next, index-1, val)
	return head
}

func (list *LinkedList) AddFirst(val interface{}) {
	list.Add(0, val)
}

func (list *LinkedList) AddLast(val interface{}) {
	list.Add(list.size, val)
}

func (list *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Get failed, require index >= 0 && index < list.size")
	}
	return list.get(list.head, index)
}

func (list *LinkedList) get(head *node, index int) interface{} {
	if index == 0 {
		return head.val
	}
	return list.get(head.next, index-1)
}

func (list *LinkedList) GetFirst() interface{} {
	return list.Get(0)
}

func (list *LinkedList) GetLast() interface{} {
	return list.Get(list.size - 1)
}

func (list *LinkedList) Contains(val interface{}) bool {
	return list.contains(list.head, val)
}

func (list *LinkedList) contains(head *node, val interface{}) bool {
	if head == nil {
		return false
	}
	if head.val == val {
		return true
	}
	return list.contains(head.next, val)
}

func (list *LinkedList) Set(index int, val interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed, require index >= 0 && index < list.size")
	}
	list.set(list.head, index, val)
}

func (list *LinkedList) set(head *node, index int, val interface{}) {
	if index == 0 {
		head.val = val
		return
	}
	list.set(head.next, index-1, val)
}

func (list *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Remove failed, require index >= 0 && index < list.size")
	}
	newHead, delNode := list.remove(list.head, index)
	list.head = newHead
	return delNode.val
}

func (list *LinkedList) remove(head *node, index int) (newHead *node, delNode *node) {
	if index == 0 {
		delNode = head
		head = delNode.next
		delNode.next = nil
		list.size--
		return head, delNode
	}
	head.next, delNode = list.remove(head.next, index-1)
	return head, delNode
}

func (list *LinkedList) RemoveFirst() interface{} {
	return list.Remove(0)
}

func (list *LinkedList) RemoveLast() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedList) RemoveElement(val interface{}) bool {
	newHead, ok := list.removeElement(list.head, val)
	list.head = newHead
	return ok
}

func (list *LinkedList) removeElement(head *node, val interface{}) (newHead *node, ok bool) {
	if head == nil {
		return head, false
	}
	if head.val == val {
		newHead = head.next
		head.next = nil
		list.size--
		return newHead, true
	}
	newHead, ok = list.removeElement(head.next, val)
	head.next = newHead
	return head, ok
}

func (list *LinkedList) RemoveAll(val interface{}) bool {
	newHead, ok := list.removeAll(list.head, val, false)
	list.head = newHead
	return ok
}

func (list *LinkedList) removeAll(head *node, val interface{}, ok bool) (*node, bool) {
	if head == nil {
		return head, ok
	}
	if head.val == val {
		list.size--
		ok = true
		head, ok = list.removeAll(head.next, val, ok)
		return head, ok
	} else {
		head.next, ok = list.removeAll(head.next, val, ok)
		return head, ok
	}
}

func (list *LinkedList) String() string {
	var buffer bytes.Buffer
	for cur := list.head; cur != nil; cur = cur.next {
		buffer.WriteString(fmt.Sprintf("%v->", cur.val))
	}
	buffer.WriteString("nil")
	return buffer.String()
}
