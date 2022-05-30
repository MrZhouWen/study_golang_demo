package 数组

/**
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

思路: 快速选择排序 T= O(n) S = O(n)
快排可以解决问题，但是它需要确定数组中所有元素的正确位置，对于本题而言，我们只需要确定第k大元素的位置pos,
我们只需要确保pos左边的元素都比它小，pos右边的元素都比它大即可，不需要关心其左边和右边的集合是否有序，
所以，我们需要对快排进行改进(以下称分区)，将目标值的位置pos与分区函数Partition求得的位置index进行比对，
	如果两值相等，说明index对应的元素即为所求值，
	如果index<pos，则递归的在[index+1, right]范围求解；
	否则则在[left, index-1]范围求解，如此便可大幅缩小求解范围。
*/
func findKthLargest(nums []int, k int) int {
	TopSplit(nums, len(nums)-k, 0, len(nums)-1) //这里第二个参数，表示我们需要 第K大的元素
	return nums[len(nums)-k]
}

func TopSplit(nums []int, k, start, stop int) {
	if start < stop {
		index := Parition(nums, start, stop) //找基准元素的位置
		if index == k {
			return
		} else if index < k {
			TopSplit(nums, k, index+1, stop)
		} else {
			TopSplit(nums, k, start, index-1)
		}
	}
}

//获取分治后的基准元素下标
func Parition(nums []int, start, stop int) int {
	if start >= stop {
		return -1
	}
	pivot := nums[start]
	l, r := start, stop
	for l < r {
		for l < r && nums[r] > pivot { //从右边开始，找第一个小于pivot的元素
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot { //从左边开始，找第一个大于 pivot的元素
			l++
		}
		nums[r] = nums[l]
	}
	//循环结束后，l == r
	nums[l] = pivot //基准元素，放到他的最终位置
	return l
}

/** 以下为TopK类问题的通用解
1 获得前k小的数
func TopkSmallest(nums []int, k int) []int {
	TopkSplit(nums, k-1, 0, len(nums)-1)
	return nums[:k]
}

2 获得前k大的数
func TopkLargest(nums []int, k int) []int {
	TopkSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k:]
}

3 获取第k小的数
func TopkSmallestElement(nums []int, k int) int {
	TopkSplit(nums, k-1, 0, len(nums)-1)
	return nums[k-1]
}

4 获取第k大的数
func TopkLargestElement(nums []int, k int) int {
	TopkSplit(nums, len(nums)-k, 0, len(nums)-1)
	return nums[len(nums)-k]
}
*/
