package 数组

import (
	"math/rand"
)

/**
给定一个可能含有重复元素的整数数组，要求随机输出给定的数字的索引。 您可以假设给定的数字一定存在于数组中。

注意：
数组大小可能非常大。 使用太多额外空间的解决方案将不会通过测试。

思路：水塘抽样  pick(target) = O(n) S = O(1)
	1、遍历 nums，当我们第 i 次遇到值为 target 的元素时，随机选择区间 [0,i) 内的一个整数，如果其等于 0，则将返回值置为该元素的下标，
	否则返回值不变。
	2、设 nums 中有 k 个值为 target 的元素，该算法会保证这 k 个元素的下标成为最终返回值的概率均为 1/k
	证明参考：https://leetcode-cn.com/problems/random-pick-index/solution/sui-ji-shu-suo-yin-by-leetcode-solution-ofsq/
*/
type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (s Solution) Pick(target int) (ans int) {
	cnt := 0 //target遇到的次数
	for i, num := range s {
		if target == num {
			cnt++ //遇到target的次数+1
			if rand.Intn(cnt) == 0 {
				ans = i
			}
		}
	}
	return
}
