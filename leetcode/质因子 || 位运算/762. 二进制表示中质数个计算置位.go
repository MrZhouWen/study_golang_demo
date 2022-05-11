package 质因子____位运算

/*
题目：
给你两个整数left和right ，在闭区间 [left, right]范围内，统计并返回 计算置位位数为质数 的整数个数。
计算置位位数 就是二进制表示中 1 的个数。
例如， 21的二进制表示10101有 3 个计算置位。

思路：
质因数 + 位运算
*/
func countPrimeSetBits(left int, right int) (ans int) {
	for i := left; i <= right; i++ {
		countOne := countOneBits(i)
		if isPrime(countOne) {
			ans++
		}
	}
	return
}

func countOneBits(num int) (ans int) {
	for ; num > 0; num >>= 1 {
		if num&1 == 1 {
			ans++
		}
	}
	return
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
