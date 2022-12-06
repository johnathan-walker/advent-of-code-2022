package heap

type Value[T comparable] interface {
	Value() T
}

type HeapInt int

func (hi HeapInt) Value() int {
	return int(hi)
}
