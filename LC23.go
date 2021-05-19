package main

import "container/heap"

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	h := new(minHeap)
	for i := 0; i < n; i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummy := new(ListNode)
	cur := dummy
	for h.Len() > 0 {
		temp := heap.Pop(h).(*ListNode)
		if temp.Next != nil {
			heap.Push(h, temp.Next)
		}
		cur.Next = temp
		cur = cur.Next
	}
	return dummy.Next
}

type minHeap []*ListNode

func (h minHeap) Len() int {
	return len(h)
}
func (h minHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h minHeap) Swap(i ,j int){
	h[i], h[j] = h[j], h[i]
}
func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
