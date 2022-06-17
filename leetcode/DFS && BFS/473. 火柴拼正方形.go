package DFS____BFS

import "sort"

/**
你将得到一个整数数组 matchsticks ，其中 matchsticks[i] 是第 i个火柴棒的长度。你要用 所有的火柴棍拼成一个正方形。
你 不能折断 任何一根火柴棒，但你可以把它们连在一起，而且每根火柴棒必须 使用一次 。
如果你能使这个正方形，则返回 true ，否则返回 false 。

思路：dfs + 回溯 T= O(4^n) S = O(n)
*/
func makesquare(matchsticks []int) bool {
	var totalLen int
	for _, v := range matchsticks {
		totalLen += v
	}
	if totalLen%4 != 0 {
		return false
	}
	sort.Slice(matchsticks, func(i, j int) bool { //递减为了减少递归的次数
		return matchsticks[i] > matchsticks[j]
	})
	edges := [4]int{} //正方形的四条边的长度
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		for i := range edges {
			edges[i] += matchsticks[idx] //选择matchsticks[i]
			if edges[i] <= totalLen/4 && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx] //回溯处理，不选择matchsticks[i]
		}
		return false
	}
	return dfs(0) //初始传入第一根火柴matchsticks[0]
}
