package 堆____栈

import (
	"container/heap"
)

/**
请返回大量数组中 最小的K个数

思路：采用大根堆的思想，一开始维护前k个数组元素,从第k+1个元素开始，与堆顶的进行比较，如果小于堆顶的，则说明是属于前k个最小的数，
替换堆顶元素，重新建堆即可, 重建堆时间复杂度O(logk)
时间复杂度:O(nlogk)
空间复杂度:O(k)
*/
//type hp struct {
//	sort.IntSlice
//}
//
////大根堆，所以自上而下 是越来越小的
//func (h hp) Less(i, j int) bool {
//	return h.IntSlice[i] > h.IntSlice[j]
//}
//func (hp) Push(interface{})     {}
//func (hp) Pop() (_ interface{}) { return }
//
//func smallestK(arr []int, k int) []int {
//	if k == 0 {
//		return nil
//	}
//	h := &hp{arr[:k]}
//	heap.Init(h)
//	for _, v := range arr[k:] {
//		if v < h.IntSlice[0] { //说明v是前K小的数，与堆顶元素交换后，重新建堆
//			h.IntSlice[0] = v
//			heap.Fix(h, 0)
//		}
//	}
//	return h.IntSlice
//}

//-----------------------------------------------------------------
type maxHeap []int

func (hp maxHeap) Len() int {
	return len(hp)
}
func (hp maxHeap) Less(i, j int) bool {
	return hp[i] > hp[j] //大根堆
}
func (hp maxHeap) Swap(i, j int) {
	hp[i], hp[j] = hp[j], hp[i]
}
func (hp *maxHeap) Push(x interface{}) {
	*hp = append(*hp, x.(int))
}
func (hp *maxHeap) Pop() interface{} {
	old := *hp
	n := len(old)
	x := old[n-1] //堆顶元素
	*hp = old[:n-1]
	return x
}
func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	hp := &maxHeap{}
	*hp = append(*hp, arr[:k]...)
	heap.Init(hp)
	for _, v := range arr[k:] {
		if v < (*hp)[0] { //v小于堆顶元素，说明是前K小的，要加入堆,这里是堆顶元素，总是取(*hp)[0]
			_ = heap.Pop(hp).(int)
			heap.Push(hp, v)
		}
	}
	return *hp
}
