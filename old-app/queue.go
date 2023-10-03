package simian

import ()

// slice of pointers to Events
type EventHeap []*Event

func (h EventHeap) Len() int {
	return len(h)
}

// timestamp determines priority
func (h EventHeap) Less(i, j int) bool {
	return h[i].timestamp < h[j].timestamp
}

func (h EventHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push and Pop supported (peak?)
func (h *EventHeap) Push(x interface{}) {
	event := x.(*Event)
	*h = append(*h, event)
}

func (h *EventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	event := old[n-1]
	old[n-1] = nil //avoid mem leak
	*h = old[0 : n-1]
	return event
}
