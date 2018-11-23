package queue

type Queue interface {
	GetSize() int
	IsEmpty() bool
	Offer(e interface{})
	Poll() interface{}
	Front() interface{}
}
