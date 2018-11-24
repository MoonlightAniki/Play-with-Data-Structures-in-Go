package queue

import (
	"bytes"
	"fmt"
)

// tail == front ==> 队列为空
// (tail + 1) / len(data) == front ==> 队列满
type LoopQueue struct {
	data  []interface{}
	front int
	tail  int //队列尾，新入队的元素添加到tail位置
	size  int
}

func CreateLoopQueue() *LoopQueue {
	q := &LoopQueue{}
	q.data = make([]interface{}, 10)
	q.front = 0
	q.tail = 0
	q.size = 0
	return q
}

// 返回循环队列的容量
func (q *LoopQueue) GetCapacity() int {
	return len(q.data) - 1
}

func (q *LoopQueue) GetSize() int {
	return q.size
}

func (q *LoopQueue) IsEmpty() bool {
	return q.size == 0
}

func (q *LoopQueue) Offer(e interface{}) {
	// 如果队列已经满了进行扩容
	if (q.tail+1)%len(q.data) == q.front {
		q.resize(q.GetCapacity() * 2)
	}
	q.data[q.tail] = e
	q.tail = (q.tail + 1) % len(q.data)
	q.size++
}

func (q *LoopQueue) Poll() interface{} {
	if q.tail == q.front {
		panic("Poll Failed, LoopQueue is empty")
	}
	res := q.data[q.front]
	q.front = (q.front + 1) % len(q.data)
	q.size--
	if q.size == q.GetCapacity()/4 && q.GetCapacity()/2 > 0 {
		q.resize(q.GetCapacity() / 2)
	}
	return res
}

func (q *LoopQueue) Front() interface{} {
	if q.tail == q.front {
		panic("Front Failed, LoopQueue is empty")
	}
	return q.data[q.front]
}

func (q *LoopQueue) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity+1)
	index := 0
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		newData[index] = q.data[i]
		index++
	}
	q.data = newData
	q.front = 0
	q.tail = q.size
}

func (q *LoopQueue) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("LoopQueue, front [")
	for i := q.front; i != q.tail; i = (i + 1) % len(q.data) {
		buffer.WriteString(fmt.Sprintf("%v", q.data[i]))
		if (i+1)%len(q.data) != q.tail {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("] tail")
	return buffer.String()
}
