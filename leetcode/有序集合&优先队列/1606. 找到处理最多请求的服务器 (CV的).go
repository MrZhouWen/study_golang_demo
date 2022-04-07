package 有序集合_优先队列

import "container/heap"

/*
题目：
你有 k个服务器，编号为 0到 k-1。每个服务器有无穷的计算能力但是 一次只能处理一个请求。请求分配到服务器的规则如下：
第i（序号从 0 开始）个请求到达。
	如果所有服务器都已被占据，那么该请求被舍弃（完全不处理）。
	如果第(i % k)个服务器空闲，那么对应服务器会处理该请求。
	否则，将请求安排给下一个空闲的服务器（服务器构成一个环，必要的话可能从第 0 个服务器开始继续找下一个空闲的服务器）。
	比方说，如果第 i个服务器在忙，那么会查看第 (i+1)个服务器，第 (i+2)个服务器等等。

给你一个 严格递增的正整数数组arrival，表示第i个任务的到达时间，和另一个数组load，其中load[i]表示第i个请求的工作量（也就是服务器完成它所需要的时间）。
你的任务是找到 最繁忙的服务器。最繁忙定义为一个服务器处理的请求数是所有服务器里最多的。
请你返回包含所有最繁忙服务器序号的列表，你可以以任意顺序返回这个列表。

思路：
有序集合(红黑树) + 优先队列
*/
type pair struct{ end, id int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func busiestServers(k int, arrival, load []int) (ans []int) {
	available := redblacktree.NewWithIntComparator()
	for i := 0; i < k; i++ {
		available.Put(i, nil)
	}
	busy := hp{}
	requests := make([]int, k)
	maxRequest := 0
	for i, t := range arrival {
		for len(busy) > 0 && busy[0].end <= t {
			available.Put(busy[0].id, nil)
			heap.Pop(&busy)
		}
		if available.Size() == 0 {
			continue
		}
		node, _ := available.Ceiling(i % k)
		if node == nil {
			node = available.Left()
		}
		id := node.Key.(int)
		requests[id]++
		if requests[id] > maxRequest {
			maxRequest = requests[id]
			ans = []int{id}
		} else if requests[id] == maxRequest {
			ans = append(ans, id)
		}
		heap.Push(&busy, pair{t + load[i], id})
		available.Remove(id)
	}
	return
}
