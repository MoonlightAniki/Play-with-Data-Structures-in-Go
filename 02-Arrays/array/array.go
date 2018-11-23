package array

import (
	"bytes"
	"fmt"
	"strconv"
)

type Array struct {
	data []int // 自定义数组的数组
	size int   // 元素个数
}

// 创建指定容量的自定义数组
// Go语言中数组类型的长度是编译时指定，无法满足此需求，这里使用slice代替
func GetArray(capacity int) *Array {
	a := &Array{}
	a.data = make([]int, capacity)
	a.size = 0
	return a
}

// 获取数组的容量
func (a *Array) GetCapacity() int {
	return len(a.data)
}

// 获取数组中元素的数量
func (a *Array) GetSize() int {
	return a.size
}

// 检查数组是否为空
func (a *Array) IsEmpty() bool {
	return a.size == 0
}

// 添加元素
func (a *Array) Add(index int, element int) {
	if a.size == len(a.data) {
		panic("Add failed, Array is full.")
	}
	if index < 0 || index > len(a.data) {
		panic("Add failed, require index >= 0 && index <= a.size")
	}
	for i := a.size; i > index; i-- {
		a.data[i] = a.data[i-1]
	}
	a.data[index] = element
	a.size++
}

func (a *Array) AddFirst(element int) {
	a.Add(0, element)
}

func (a *Array) AddLast(element int) {
	a.Add(a.size, element)
}

func (a *Array) Get(index int) int {
	if index < 0 || index >= a.size {
		panic("Get failed, require index >= 0 && index < a.size")
	}
	return a.data[index]
}

func (a *Array) Set(index int, element int) {
	if index < 0 || index >= a.size {
		panic("Set failed, require index >= 0 && index < a.size")
	}
	a.data[index] = element
}

func (a *Array) Contains(element int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return true
		}
	}
	return false
}

// 返回第一个等于element的元素的索引，不存在则返回-1
func (a *Array) Find(element int) int {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			return i
		}
	}
	return -1
}

// 返回数组中所有等于element的元素的索引组成的切片,不存在则返回[]
func (a *Array) FindAll(element int) (indexes []int) {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			indexes = append(indexes, i)
		}
	}
	return
}

func (a *Array) Remove(index int) int {
	if index < 0 || index >= a.size {
		panic("Remove failed, require index >= 0 && index < a.size")
	}
	ret := a.data[index]
	for i := index; i+1 < a.size; i++ {
		a.data[i] = a.data[i+1]
	}
	a.data[a.size-1] = 0
	a.size--
	return ret
}

func (a *Array) RemoveFirst() int {
	return a.Remove(0)
}

func (a *Array) RemoveLast() int {
	return a.Remove(a.size - 1)
}

func (a *Array) RemoveElement(element int) bool {
	for i := 0; i < a.size; i++ {
		if a.data[i] == element {
			a.Remove(i)
			return true
		}
	}
	return false
}

func (a *Array) RemoveAll(element int) bool {
	res := false
	for i := a.size - 1; i >= 0; i-- {
		if a.data[i] == element {
			a.Remove(i)
			res = true
		}
	}
	return res
}

// 实现Stringer接口，重写String()函数
func (a *Array) String() string {
	var buffer bytes.Buffer

	buffer.WriteString(fmt.Sprintf("Array: size = %d, capacity = %d\n", a.size, len(a.data)))
	buffer.WriteString("[")
	for i := 0; i < a.size; i++ {
		buffer.WriteString(strconv.Itoa(a.data[i]))
		if i != a.size-1 {
			buffer.WriteString(", ")
		}
	}
	buffer.WriteString("]")
	return buffer.String()
}
