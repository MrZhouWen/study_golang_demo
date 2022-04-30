package DFS____BFS

/**
题目：
给你一个整数 n ，按字典序返回范围 [1, n] 内所有整数。
要求 T = O(n) 且 S = O(1)

思路：DFS
字典序最小的整数为 number=1，我们从它开始，然后依次获取下一个字典序整数，加入结果中，结束条件为已经获取到 n 个整数。
那么对于一个整数 number，它的下一个字典序整数对应下面的规则：
	1、尝试在 number 后面附加一个零，即 number×10，如果 number×10≤n，那么说明 number×10 是下一个字典序整数；
	2、边界：如果 number%10=9 或 number+1>n，那么说明number末尾的数位已经搜索完成，退回上一位，即 number=⌊number/10⌋，
	然后继续判断直到 number%10=9 且 number+1≤n 为止，那么 number+1 是下一个字典序整数。
*/
func lexicalOrder(n int) []int {
	ans := make([]int, n)
	num := 1
	for i := range ans {
		ans[i] = num
		if num*10 <= n {
			num = num * 10
		} else {
			/*
				末尾位数字 到达边界条件，得回退到上一位
				注意这里得是for循环：比如 num = 199 进行回退
			*/
			for num%10 == 9 || num+1 > n {
				num /= 10
			}
			num++ //num/10后得到上一位，字典序+1
		}
	}
	return ans
}
