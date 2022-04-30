package 数组

/**
给你一个整数数组 nums，将 nums 中的的所有偶数元素移动到数组的前面，后跟所有奇数元素。
返回满足此条件的 任一数组 作为答案。

思路：双指针 + 原地交换
S= O(1) T= O(n)
*/
func sortArrayByParity(nums []int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]%2 == 0 {
			left++
		}
		for left < right && nums[right]%2 == 1 {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
		left++
		right--
	}
	return nums
}
