package stack

type Stack interface {
	GetSize() int

	IsEmtpy() bool

	Push(e interface{})

	Pop() interface{}

	Peek() interface{}
}
