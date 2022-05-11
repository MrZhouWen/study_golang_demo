package 数组

/**
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。

思路：滑动窗口
T = O(n) S = O(1)

我们固定子数组 [i,j] 的右端点 j 时，显然左端点 i 越大，子数组元素乘积越小。
对于符合条件的子数组 [i,j]，那么元素乘积小于 k 的子数组数目为 j−i+1。返回所有数目之和。
######## 这里说明 为什么每次滑动确定一个窗口 后，结果累加 j−i+1， 因为从 [i, j-1]中的所有子数组，已经算在之前的数组中，
######## 那么每一轮 j 后移， 拼接上前面的每一个元素，比如 [j,j], [j-1, j], [j-2, j], [j-3, j]....[i, j], 会产生j-i+1个新的子数组
*/
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	product, i := 1, 0         //乘积 和 左端点
	for j, num := range nums { //右端点j
		product *= num
		for ; i <= j && product >= k; i++ { //对于每一个右端点j， 寻找对应的左端点i
			product /= nums[i]
		}
		ans += j - i + 1
	}
	return
}
