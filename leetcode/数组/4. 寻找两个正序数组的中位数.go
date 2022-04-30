package 数组

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组nums1 和nums2。请你找出并返回这两个正序数组的 中位数 。
算法的时间复杂度应该为 O(log (m+n)) 。

思路 ： 二路归并 || 二分查找
参考：https://leetcode-cn.com/problems/median-of-two-sorted-arrays/solution/xun-zhao-liang-ge-you-xu-shu-zu-de-zhong-wei-s-114/
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) { //表示第一个数组已经遍历完了
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) { //同理：第二个数组遍历完了
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		//需要找到2个数组中的第k个数，那么每个数组 各拿 k/2 出来比较，比较出较小的部分直接丢掉
		half := k / 2
		newIndex1, newIndex2 := min(index1+half, len(nums1))-1, min(index2+half, len(nums2))-1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 { //丢掉pivot1以及前面的所有数
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else { //丢掉pivot2以及之前的数
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
