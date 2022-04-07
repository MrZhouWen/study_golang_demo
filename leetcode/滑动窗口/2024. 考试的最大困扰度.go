package 滑动窗口

/**
题目：
给你一个字符串answerKey，其中answerKey[i]是第 i个问题的正确结果T 或 F。
除此以外，还给你一个整数 k，表示你能进行以下操作的最多次数：
	每次操作中，将 answerKey[i] 改为'T'或者'F'。请你返回在不超过 k次操作的情况下，最大连续 'T'或者 'F'的数目。

思路：
滑动窗口 T(n) S(1)
*/
func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(maxConsecutiveChar(answerKey, k, 'T'), maxConsecutiveChar(answerKey, k, 'F'))
}

func maxConsecutiveChar(answerKey string, k int, ch byte) (ans int) {
	left, sum := 0, 0
	for right := range answerKey {
		if answerKey[right] != ch {
			sum++
		}
		for sum > k { //这里的for循环，不能写成if判断
			if answerKey[left] != ch {
				sum--
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
