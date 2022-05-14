package 前缀和

/**
参考：https://leetcode.cn/problems/plates-between-candles/

思路：预处理 + 前缀和 T(n+q) S(n)
*/
func platesBetweenCandles(s string, queries [][]int) []int {
	n := len(s)
	preSum := make([]int, n)
	left := make([]int, n)  //left[i]表示下标为i的元素，它左手边第一个蜡烛的位置
	right := make([]int, n) //right[i]表示下标为i的元素，它右手边第一个蜡烛的位置

	sum, l, r := 0, -1, -1
	for i, ch := range s {
		if ch == '*' {
			sum++
		} else {
			l = i
		}
		preSum[i] = sum
		left[i] = l //i左边第一个蜡烛是l位置
	}

	for i := n - 1; i >= 0; i-- {
		if s[i] == '|' {
			r = i
		}
		right[i] = r
	}

	ans := make([]int, len(queries))
	for i, q := range queries {
		//【这里有点绕】q[0]右手边的第一个蜡烛就是左端点，同理。。。
		x, y := right[q[0]], left[q[1]]
		if x >= 0 && y >= 0 && x < y {
			//按照道理是preSum[y] - preSum[x-1]但是，下标x处是蜡烛，所以防止x = 0时，越界，写成这样
			ans[i] = preSum[y] - preSum[x]
		}
	}

	return ans
}
