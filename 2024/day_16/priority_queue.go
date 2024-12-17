package day_16

import "container/heap"

type bot struct {
	y         int
	x         int
	score     int
	direction rune

	visited string

	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*bot

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use greater than here.
	return pq[i].score < pq[j].score
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*bot)
	item.index = n
	// item.inQueue = true
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	// item.inQueue = false
	return item
}

// Update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) Update(item *bot, score int) {
	item.score = score
	heap.Fix(pq, item.index)
}
