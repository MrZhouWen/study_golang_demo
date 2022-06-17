package 二分查找

import "sort"

/**
给一堆香蕉，每一堆中香蕉个数为pile[i], 求每小时吃K根的话，h小时之内把所有堆的香蕉吃完，求最小的K？

思路：二分查找 T= O(logn) S= O(1)
对于每一堆香蕉有pile个，假设吃的K = speed, 那么吃掉这一堆要 「pile/speed「向上取整，==>等价于 [pile+(speed-1)]/speed
所以，最小速度是1， 最大速度是 max(数量最多的一个堆中的香蕉数量)
speed 从 1 ~ max 进行二分查找，吃掉每一堆的时间 = [pile+(speed-1)]/speed， 累加起来 <= h 即为最小的 K
*/
func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	//这里加1 是因为sort.Search返回的是 0 ~ (max-1)中的下标，但是吃的速度 K >= 1，所以+1
	return 1 + sort.Search(max-1, func(speed int) bool {
		speed++ //这里speed是 0 ~ (max-1) 中的下标【理解起来很关键】
		time := 0
		for _, pile := range piles {
			time += (pile + speed - 1) / speed
		}
		return time <= h
	})
}
