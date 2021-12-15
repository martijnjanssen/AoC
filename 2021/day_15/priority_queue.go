package day_15

type point struct {
	y int
	x int
	r int
	small bool
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].small && !pq[j].small {
		return true
	} else if !pq[i].small && pq[j].small {
		return false
	}
	// We want Pop to give us the lowest, not highest, priority so we use greater than here.
	return pq[i].r < pq[j].r
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*point))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
// func (pq *PriorityQueue) update(item *point, y int, x int, r int) {
// 	item.y = y
// 	item.x = x
// 	item.r = r
// 	heap.Fix(pq, item.index)
// }
