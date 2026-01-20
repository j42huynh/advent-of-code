package main

type Distance struct {
	distance    float64
	coord1      []int
	coord1Index int
	coord2      []int
	coord2Index int
}

// Define a custom type for the min-distance-heap, which is a slice of distances.
type MinDistanceHeap []Distance

// Implement the sort.Interface methods: Len, Less, and Swap.
func (h MinDistanceHeap) Len() int {
	return len(h)
}

// Less implements the custom comparator.
func (h MinDistanceHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinDistanceHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Implement the heap.Interface methods: Push and Pop.
// These methods are used internally by the heap package functions.
func (h *MinDistanceHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length.
	*h = append(*h, x.(Distance))
}

func (h *MinDistanceHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
