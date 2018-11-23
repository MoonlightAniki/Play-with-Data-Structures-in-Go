package array

import (
	"bytes"
	"fmt"
)

type Array struct {
	data []interface{}
	size int
}

func CreateArray() *Array {
	a := &Array{}
	a.data = make([]interface{}, 10)
	a.size = 0
	return a
}

func (a *Array) GetSize() int {
	return a.size
}

func (a *Array) GetCapacity() int {
	return len(a.data)
}

func (a *Array) IsEmpty() bool {
	return a.size == 0
}

func (a *Array) resize(newCapacity int) {
	newData := make([]interface{}, newCapacity)
	for i := 0; i < a.size; i++ {
		newData[i] = a.data[i]
	}
	a.data = newData
}

func (a *Array) Add(index int, element interface{}) {
	if index < 0 || index > a.size {
		panic("Add failed, require index >= 0 && index <= a.size")
	}
	if a.size == len(a.data) {
		a.resize(2 * len(a.data))
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = element
	a.size++
}

func (a *Array) AddFirst(element interface{}) {
	a.Add(0, element)
}

func (a *Array) AddLast(element interface{}) {
	a.Add(a.size, element)
}

func (a *Array) Get(index int) interface{} {
	if index < 0 || index >= a.size {
		panic("Get failed, require index >= 0 && index < a.size")
	}
	return a.data[index]
}

func (a *Array) GetFirst() interface{} {
	return a.Get(0)
}

func (a *Array) GetLast() interface{} {
	return a.Get(a.size - 1)
}

func (a *Array) Set(index int, element interface{}) {
	if index < 0 || index >= a.size {
		panic("Set failed, require index >= 0 && index < a.size")
	}
	a.data[index] = element
}

func (a *Array) Contains(element interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}
	return false
}

func (a *Array) Find(element interface{}) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i
		}
	}
	return -1
}

func (a *Array) FindAll(element interface{}) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			indexes = append(indexes, i)
		}
	}
	return
}

func (a *Array) Remove(index int) (element interface{}) {
	if index < 0 || index >= a.size {
		panic("Remove failed, require index >= 0 && index < a.size")
	}
	element = a.data[index]
	for i := index; i+1 < a.size; i++ {
		a.data[i] = a.data[i+1 ]
	}
	a.size--
	if a.size == len(a.data)/4 && len(a.data)/2 > 0 {
		a.resize(len(a.data) / 2)
	}
	return
}

func (a *Array) RemoveFirst() interface{} {
	return a.Remove(0)
}

func (a *Array) RemoveLast() interface{} {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(element interface{}) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			a.Remove(i)
			return true
		}
	}
	return false
}

func (a *Array) RemoveAll(element interface{}) bool {
	res := false
	for i := a.size - 1; i >= 0; i-- {
		if a.data[i] == element {
			a.Remove(i)
			res = true
		}
	}
	return res
}

func (a *Array) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Array, size = %d, capacity = %d", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(fmt.Sprintf("%v", a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
