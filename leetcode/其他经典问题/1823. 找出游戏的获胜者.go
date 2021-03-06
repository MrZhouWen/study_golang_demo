package 其他经典问题

/**
约瑟夫环问题
共有 n 名小伙伴一起做游戏。小伙伴们围成一圈，按 顺时针顺序 从 1 到 n 编号。确切地说，
游戏遵循如下规则：
	1、从第 1 名小伙伴所在位置 开始 。沿着顺时针方向数 k 名小伙伴，计数时需要 包含 起始时的那位小伙伴。逐个绕圈进行计数，一些小伙伴可能会被数过不止一次。
	你数到的最后一名小伙伴需要离开圈子，并视作输掉游戏。
	2、如果圈子中仍然有不止一名小伙伴，从刚刚输掉的小伙伴的 顺时针下一位 小伙伴 开始，回到步骤 1 继续执行。
	否则，圈子中最后一名小伙伴赢得游戏。
给你参与游戏的小伙伴总数 n ，和一个整数 k ，返回游戏的获胜者。

思路：
方法一：队列 + 模拟 T= O(nk) S= O(n)
方法二：数学 + 递归 || 迭代
*/

/*方法一：队列
func findTheWinner(n, k int) int {
    q := make([]int, n)
    for i := range q {
        q[i] = i + 1
    }
    for len(q) > 1 {
        for i := 1; i < k; i++ {
            q = append(q, q[0])[1:]
        }
        q = q[1:]
    }
    return q[0]
}
*/

/**
方法二：数学 + 递归 || 迭代
以下用 f(n, k) 表示当前轮人数为n时获胜者编号,意思是 从当前轮第一个人开始数，数到k'个人结束，淘汰。
	当 n=1 时，圈子中只有一名小伙伴，该小伙伴即为获胜者，因此 f(1, k) = 1。
	当 n > 1 时，将有一名小伙伴离开圈子，圈子中剩下 n - 1 名小伙伴。假设圈子中的第 k'名小伙伴离开圈子，

	且k'满足 1≤k'≤n 且 k - k' 是 n 的倍数 ==> 等价于 k' = k % n，

当圈子中剩下 n - 1 名小伙伴时，可以递归地计算 f(n−1,k)，得到剩下的 n−1 名小伙伴中的获胜者。令 x=f(n−1,k)。
	【意思是 从当前轮第一个人开始数即从 k'+1 开始数，数到第x个人结束，淘汰】
由于在第 k'名小伙伴离开圈子之后，圈子中剩下的 n - 1名小伙伴从第 k' + 1名小伙伴开始计数，

	☆☆☆☆☆这一句很重要：n-1轮的「获胜者编号是从第 (k' + 1) 名小伙伴开始数的第 x 名小伙伴」,即编号是 第 (k‘+x) 名 = [(k'+(x-1))%n + 1] 名小伙伴
	☆☆☆☆☆这一句很重要：n-1轮的「获胜者编号是从第 (k' + 1) 名小伙伴开始数的第 x 名小伙伴」
	☆☆☆☆☆这一句很重要：n-1轮的「获胜者编号是从第 (k' + 1) 名小伙伴开始数的第 x 名小伙伴」

	因此f(n,k)=(k′+x−1)%n+1 = (k % n+x−1)%n+1。

	将 x=f(n−1,k) 代入上述关系，可得：f(n,k)=(k + f(n−1,k) −1 ) % n+1。

	T = O(n) 每一轮淘汰的人的时间复杂度都是O(1) ，要进行n轮
	S= O(n) 递归的栈空间
*/
func findTheWinner(n, k int) int {
	if n == 1 {
		return 1
	}
	return (k+findTheWinner(n-1, k)-1)%n + 1
}

//以上还可以用迭代来节省递归的栈空间
func findTheWinner_upgrade(n, k int) int {
	winner := 1
	for i := 2; i <= n; i++ {
		winner = (k+winner-1)%i + 1 //这里一定要注意 是 %i 而不是%n ,因为每一次循环，其实 就是计算一轮 f(i, k)
	}
	return winner
}
