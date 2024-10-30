package queue

import (
	"pulseCommunity/ranking/infrastructure/pkg/pri_que"
	"pulseCommunity/ranking/infrastructure/pkg/pri_que/internal/queue"
)

type PriorityQueue[T any] struct {
	priorityQueue *queue.PriorityQueue[T]
}

func NewPriorityQueue[T any](capacity int, compare pri_que.Comparator[T]) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	pq.priorityQueue = queue.NewPriorityQueue[T](capacity, compare)
	return pq
}

func (pq *PriorityQueue[T]) Len() int {
	return pq.priorityQueue.Len()
}

func (pq *PriorityQueue[T]) Peek() (T, error) {
	return pq.priorityQueue.Peek()
}

func (pq *PriorityQueue[T]) Enqueue(t T) error {
	return pq.priorityQueue.Enqueue(t)
}

func (pq *PriorityQueue[T]) Dequeue() (T, error) {
	return pq.priorityQueue.Dequeue()
}
