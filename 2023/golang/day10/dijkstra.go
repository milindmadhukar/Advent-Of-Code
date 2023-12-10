package day10

import (
	"container/heap"
	"math"
)

func dijkstra(start *vertex, vertices map[point]*vertex) {
	// Initialize distances to infinity
	for _, v := range vertices {
		v.distance = math.MaxInt
	}
	start.distance = 0

	// Priority queue (min heap) for vertices
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Item).value

		for _, neighbor := range current.adjacents {
			// Calculate tentative distance to neighbor
			distance := current.distance + 1 // Assuming unweighted graph

			if distance < neighbor.distance {
				// Update distance if new distance is smaller
				neighbor.distance = distance
				heap.Push(&pq, &Item{value: neighbor, priority: distance})
			}
		}
	}
}

// Priority queue implementation
type Item struct {
	value    *vertex
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
