package heap

import (
	"container/heap"
	"sort"
)

var factors = []int{2, 3, 5}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) {
	h.IntSlice = append(h.IntSlice, v.(int))
}

func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func nthUglyNumber(n int) int {
	h := &hp{sort.IntSlice{1}}
	seen := map[int]struct{}{1: {}}
	for i := 1; ; i++ {
		x := heap.Pop(h).(int)
		if i == n {
			return x
		}
		for _, f := range factors {
			next := x * f
			if _, has := seen[next]; !has {
				heap.Push(h, next)
				seen[next] = struct{}{}
			}
		}
	}
}

//func swap(nums []int, i, j int) {
//	nums[i], nums[j] = nums[j], nums[i]
//}
//
//func heapify(nums []int, i, arrLen int) {
//	leftSon, rightSon := 2*i+1, 2*i+2
//	smallest := i
//	if leftSon < arrLen && nums[leftSon] < nums[smallest] {
//		smallest = leftSon
//	}
//	if rightSon < arrLen && nums[rightSon] < nums[smallest] {
//		smallest = rightSon
//	}
//	if i != smallest {
//		swap(nums, i, smallest)
//		heapify(nums, smallest, arrLen)
//	}
//}
//
//func buildMinHeap(nums []int, arrLen int) {
//	for i := arrLen / 2; i >= 0; i-- {
//		heapify(nums, i, arrLen)
//	}
//}
//
//func nthUglyNumber(n int) int {
//	count := 1
//	heap := []int{1}
//	seen := map[int]struct{}{1: {}}
//	for count != n {
//		buildMinHeap(heap, len(heap))
//		top := heap[0]
//		heap = heap[1:]
//		for _, i := range factors {
//			next := i * top
//			if _, has := seen[next]; !has {
//				heap = append(heap, next)
//				seen[next] = struct{}{}
//			}
//		}
//		count++
//	}
//	return heap[0]
//}
