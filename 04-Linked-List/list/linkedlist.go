package list

import (
	"bytes"
	"fmt"
)

type node struct {
	val  interface{}
	next *node
}

/*type LinkedList struct {
	head *node
	size int
}*/

/*func CreateLinkedList() *LinkedList {
	list := &LinkedList{}
	list.head = nil
	list.size = 0
	return list
}*/

type LinkedList struct {
	dummyHead *node
	size      int
}

func CreateLinkedList() *LinkedList {
	list := &LinkedList{}
	list.dummyHead = &node{}
	list.size = 0
	return list
}

func (list *LinkedList) GetSize() int {
	return list.size
}

func (list *LinkedList) IsEmpty() bool {
	return list.size == 0
}

/*func (list *LinkedList) Add(index int, val interface{}) {
	if index < 0 || index > list.size {
		panic("Add failed, require index >= 0 && index <= list.size")
	}
	if index == 0 {
		list.head = &node{val, list.head.next}
		list.size++
	} else {
		prev := list.head
		for i := 0; i < index-1; i++ {
			prev = prev.next
		}
		prev.next = &node{val, prev.next}
		list.size++
	}
}*/

func (list *LinkedList) Add(index int, val interface{}) {
	if index < 0 || index > list.size {
		panic("Add failed, require index >= 0 && index <= list.size")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = &node{val, prev.next}
	list.size++
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
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	return cur.val
}

func (list *LinkedList) GetFirst() interface{} {
	return list.Get(0)
}

func (list *LinkedList) GetLast() interface{} {
	return list.Get(list.size - 1)
}

func (list *LinkedList) Set(index int, val interface{}) {
	if index < 0 || index >= list.size {
		panic("Set failed, require index >= 0 && index < list.size")
	}
	cur := list.dummyHead.next
	for i := 0; i < index; i++ {
		cur = cur.next
	}
	cur.val = val
}

func (list *LinkedList) Contains(val interface{}) bool {
	for cur := list.dummyHead.next; cur != nil; cur = cur.next {
		if cur.val == val {
			return true
		}
	}
	return false
}

func (list *LinkedList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		panic("Remove failed, require index >= 0 && index < list.size")
	}
	prev := list.dummyHead
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	delNode := prev.next
	prev.next = delNode.next
	delNode.next = nil
	list.size--
	return delNode.val
}

func (list *LinkedList) RemoveFirst() interface{} {
	return list.Remove(0)
}

func (list *LinkedList) RemoveLast() interface{} {
	return list.Remove(list.size - 1)
}

func (list *LinkedList) String() string {
	var buffer bytes.Buffer
	for cur := list.dummyHead.next; cur != nil; cur = cur.next {
		buffer.WriteString(fmt.Sprintf("%v->", cur.val))
	}
	buffer.WriteString("nil")
	return buffer.String()
}
