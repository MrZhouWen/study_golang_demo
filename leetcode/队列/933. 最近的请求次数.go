package 队列

/**
写一个RecentCounter类来计算特定时间范围内最近的请求。

请你实现 RecentCounter 类：

RecentCounter() 初始化计数器，请求数为 0 。
int ping(int t) 在时间 t 添加一个新请求，其中 t 表示以毫秒为单位的某个时间，并返回过去 3000 毫秒内发生的所有请求数（包括新请求）。
确切地说，返回在 [t-3000, t] 内发生的请求数。
保证 每次对 ping 的调用都使用比之前更大的 t 值。s

思路：队列
我们可以用一个队列维护发生请求的时间，当在时间 t 收到请求时，将时间 t 入队。
由于每次收到的请求的时间都比之前的大，因此从队首到队尾的时间值是单调递增的。当在时间 t 收到请求时
为了求出 [t-3000,t] 内发生的请求数，我们可以不断从队首弹出 < t-3000 的时间。循环结束后队列的长度就是 [t−3000,t] 内发生的请求数。

T = O(1) 平均每次操作的时间复杂度是 O(1) 因为每个元素最多出队入队一次
S = O(n)
*/
type RecentCounter []int

func Constructor() (_ RecentCounter) {
	return
}

func (r *RecentCounter) Ping(t int) int {
	*r = append(*r, t)
	for (*r)[0] < t-3000 {
		*r = (*r)[1:] //队首元素出队
	}
	return len(*r)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
