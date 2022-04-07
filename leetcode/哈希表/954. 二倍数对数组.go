package 哈希表

import "sort"

/**
题目：
给定一个长度为偶数的整数数组 arr，只有对 arr 进行重组后可以满足 “对于每个 0 <=i < len(arr) / 2，
都有 arr[2 * i + 1] = 2 * arr[2 * i]”时，返回 true；否则，返回 false。

思路：
哈希表 + 自定义切片排序
T(nlogn) S(n)
*/
func canReorderDoubled(arr []int) bool {
	cnt := make(map[int]int, len(arr))
	for _, x := range arr {
		cnt[x]++
	}
	if cnt[0]%2 == 1 {
		return false
	}
	//对哈希表按照键值的绝对值从小到大排序，这样遍历到x的时候，只要判断cnt[2x]的个数，不用再去判断cnt[x/2]的个数
	vals := make([]int, 0, len(cnt))
	for x := range cnt {
		vals = append(vals, x)
	}
	sort.Slice(vals, func(i, j int) bool {
		return abs(vals[i]) < abs(vals[j])
	})
	for _, i := range vals {
		if cnt[2*i] < cnt[i] {
			return false
		}
		cnt[2*i] -= cnt[i]
	}
	return true
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
