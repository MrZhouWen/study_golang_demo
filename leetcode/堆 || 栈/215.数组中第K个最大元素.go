package 堆____栈

import "container/heap"

/**
小根堆 维护最大的K个数，最后返回堆顶元素即可
T = O(nlogk) S = O(k)
*/
type minHeap []int

func (hp minHeap) Less(i, j int) bool {
	return hp[i] < hp[j]
}
func (hp minHeap) Len() int {
	return len(hp)
}
func (hp minHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp *minHeap) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}
func (hp *minHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1] //堆顶元素
	*hp = old[:n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	hp := &minHeap{}
	*hp = append(*hp, nums[:k]...)
	heap.Init(hp) //不能漏
	for _, v := range nums[k:] {
		if v > (*hp)[0] { //大于堆顶元素，说明v是前K大的数
			_ = heap.Pop(hp).(int)
			heap.Push(hp, v)
		}
	}
	return (*hp)[0]
}
