package 迭代

/**
给定一个长度为 n 的整数数组nums。
	假设arrk是数组nums顺时针旋转 k 个位置后的数组，我们定义nums的 旋转函数F为：
		F(k) = 0 * arrk[0] + 1 * arrk[1] + ... + (n - 1) * arrk[n - 1]
	返回F(0), F(1), ..., F(n-1)中的最大值。生成的测试用例让答案符合32 位 整数。

输入: nums = [4,3,2,6]
输出: 26
解释:
F(0) = (0 * 4) + (1 * 3) + (2 * 2) + (3 * 6) = 0 + 3 + 4 + 18 = 25
F(1) = (0 * 6) + (1 * 4) + (2 * 3) + (3 * 2) = 0 + 4 + 6 + 6 = 16
F(2) = (0 * 2) + (1 * 6) + (2 * 4) + (3 * 3) = 0 + 6 + 8 + 9 = 23
F(3) = (0 * 3) + (1 * 2) + (2 * 6) + (3 * 4) = 0 + 2 + 12 + 12 = 26
所以 F(0), F(1), F(2), F(3) 中的最大值是 F(3) = 26 。

思路：迭代
记数组 nums 的元素之和为 numSum。根据公式，可以得到：
	F(0)=0×nums[0]+1×nums[1]+…+(n-2)×nums[n−2]+(n−1)×nums[n−1]
	F(1)=1×nums[0]+2×nums[1]+…+(n-1)×nums[n−2]+ 0   ×nums[n−1] = F(0)+numSum−n×nums[n−1]
更一般地，当 1≤k<n 时，F(k)=F(k−1)+numSum−n×nums[n−k]。我们可以不停迭代计算出不同的 F(k)，并求出最大值。
*/
func maxRotateFunction(nums []int) int {
	n := len(nums)
	numSum := 0
	for _, num := range nums {
		numSum += num
	}
	F := 0 //初始F(0)
	for i, num := range nums {
		F += i * num
	}
	ans := F //F(0)
	for i := n - 1; i > 0; i-- {
		F += numSum - n*nums[i]
		ans = max(ans, F)
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
