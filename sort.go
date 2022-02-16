package heap_sort

import "container/heap"

type Iterable interface {
	Next() (Ranked, bool)
}

type Ranked interface {
	Rank() int
}

type Queue []Ranked

func (pq Queue) Len() int {
	return len(pq)
}

func (pq Queue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq Queue) Less(i, j int) bool {
	return pq[i].Rank() > pq[j].Rank()
}

func (pq *Queue) Push(x interface{}) {
	// тут можно не проверять, тк извлекается из итерируемого объекта
	element := x.(Ranked)
	*pq = append(*pq, element)
}

func (pq *Queue) Pop() interface{} {
	n := len(*pq)
	element := (*pq)[n-1]
	*pq = (*pq)[:n-1]

	return element
}

// Sort Return [k]interface top elements(𝑛 log𝑘)
func Sort(input Iterable, k int) []interface{} {
	if k < 1 {
		return nil
	}

	q := Queue{}

	n := 0
	element, ok := input.Next()

	for ok {
		n++

		heap.Push(&q, element)

		element, ok = input.Next()
	}

	var result = make([]interface{}, min(n, k))

	for i := 0; i < n && i < k; i++ {
		elem := heap.Pop(&q)
		result[i] = elem
	}

	return result
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
