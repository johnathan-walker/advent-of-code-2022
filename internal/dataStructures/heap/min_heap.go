package heap

type MinHeap[T IntValue] []T

func (h MinHeap[T]) Init() {
	for i := (len(h) - 1) / 2; i >= 0; i-- {

	}
}

func (h MinHeap[T]) down(u int) {
	v := u
	if 2*u+1 < len(h) && h[2*u+1].Value() < h[v].Value() {
		v = 2*u + 1
	}
	if 2*u+2 < len(h) && h[2*u+2].Value() < h[v].Value() {
		v = 2*u + 2
	}
	if v != u {
		h[v], h[u] = h[u], h[v]
		h.down(v)
	}
}

func (h MinHeap[T]) up(u int) {
	for u != 0 && h[(u-1)/2].Value() > h[u].Value() {
		h[(u-1)/2], h[u] = h[u], h[(u-1)/2]
		u = (u - 1) / 2
	}
}

func (h *MinHeap[T]) Push(e T) {
	*h = append(*h, e)
	h.up(len(*h) - 1)
}

func (h *MinHeap[T]) Pop() T {
	x := (*h)[0]
	n := len(*h)
	(*h)[0], (*h)[n-1] = (*h)[n-1], (*h)[0]
	*h = (*h)[:n-1]
	h.down(0)
	return x
}
