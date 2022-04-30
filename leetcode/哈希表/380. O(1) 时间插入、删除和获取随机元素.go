package 哈希表

import "math/rand"

/**
题目：
实现RandomizedSet 类

思路：
哈希表 T = O(1) S = O(n)
*/
type RandomizedSet struct {
	nums    []int
	indices map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, map[int]int{}}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if _, ok := rs.indices[val]; ok {
		return false
	}
	rs.indices[val] = len(rs.nums)
	rs.nums = append(rs.nums, val)
	return true
}

func (rs *RandomizedSet) Remove(val int) bool {
	id, ok := rs.indices[val]
	if !ok {
		return false
	}
	last := len(rs.nums) - 1

	rs.nums[id] = rs.nums[last] //把最后一个元素，替换到id下标所在的地方

	rs.indices[rs.nums[id]] = id //更新最后一个元素的map,这里map下标应该是rs.nums[id]，因为上面更新了rs.nums[id] = rs.nums[last]
	delete(rs.indices, val)
	rs.nums = rs.nums[:last]
	return true
}

func (rs *RandomizedSet) GetRandom() int {
	return rs.nums[rand.Intn(len(rs.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
