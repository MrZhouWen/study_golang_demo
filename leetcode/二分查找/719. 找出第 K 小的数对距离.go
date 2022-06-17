package 二分查找

import "sort"

/**
给定数组，求任意2个元素的绝对差值，第K小的差值？

思路：二分 + 排序
先将数组 nums 从小到大进行排序。因为第 k 小的数对距离假设为res, 那么res必然在区间 [0,max(nums)−min(nums)] 内，
令 left=0，right=max(nums)−min(nums)，我们在区间 [left,right] 上进行二分查找res。

对于当前搜索的距离 mid，计算所有距离小于等于 mid 的数对数目 cnt，如果 cnt≥k，那么 right=mid−1，否则 left=mid+1。
当 left>right 时，终止搜索，那么第 k 小的数对距离为 left。

给定距离 mid，计算所有距离小于等于 mid 的数对数目 cnt 可以使用二分查找：
	枚举所有数对的右端点 j，二分查找大于等于 nums[j]−mid 的最小值的下标 i，那么右端点为 j 且距离小于等于 mid 的数对数目为 j−i，计算这些数目之和。

T= O(nlogn×logD)，其中 n 是数组 nums 的长度，D=max(nums)−min(nums)
*/
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool { //求出 数对的差值 <= mid 的总个数
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid) //在nums[:j]中找到第一个 >= num-mid 的下标
			cnt += j - i
		}
		return cnt >= k //要求数对的差值 <= mid 的个数cnt >= k，那么 mid才能是第K小的
	})
}
