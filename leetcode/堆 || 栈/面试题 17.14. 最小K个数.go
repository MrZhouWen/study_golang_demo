package 堆____栈

import (
	"container/heap"
	"sort"
)

/**
请返回大量数组中 最小的K个数

思路：采用大根堆的思想，一开始维护前k个数组元素,从第k+1个元素开始，与堆顶的进行比较，如果小于堆顶的，则说明是属于前k个最小的数，
替换堆顶元素，重新建堆即可, 重建堆时间复杂度O(logk)
时间复杂度:O(nlogk)
空间复杂度:O(k)
*/
type hp struct {
	sort.IntSlice
}

//大根堆，所以自上而下 是越来越小的
func (h hp) Less(i, j int) bool {
	return h.IntSlice[i] > h.IntSlice[j]
}
func (hp) Push(interface{})     {}
func (hp) Pop() (_ interface{}) { return }

func smallestK(arr []int, k int) []int {
	if k == 0 {
		return nil
	}
	h := &hp{arr[:k]}
	heap.Init(h)
	for _, v := range arr[k:] {
		if v < h.IntSlice[0] { //说明v是前K小的数，与堆顶元素交换后，重新建堆
			h.IntSlice[0] = v
			heap.Fix(h, 0)
		}
	}
	return h.IntSlice
}
