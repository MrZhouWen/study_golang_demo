package 双指针

/*
给定 数组[1,2,0,3,4,0,5,6] => [1,2,0,0,3,4,0,0] 长度不变，每个0复制一遍，其余元素后移，溢出的丢掉？

不需要原地的话，可以借助栈，每次遍历arr[i]写入栈中，是0就写入2个0,直到栈长度为 n时，依次出栈，倒着写入arr中即可

题目要求：原地复写，那么定义双指针，模拟栈操作，i指向arr中元素，top指向栈顶空位置，起始top=0,当top >= n时，栈满了
倒着写入arr即可
T= O(n) S= O(1)
*/
func duplicateZeros(arr []int) {
	i, top := -1, 0 //首先i指向arr的-1
	n := len(arr)
	for top < n { //新的 一轮
		i++ //如果i一开始定义为 0,那么这个就要写到for循环的最后一步，且最后要i--
		if arr[i] != 0 {
			top++
		} else {
			top += 2
		}
	}
	j := n - 1
	if top == n+1 { //最后遍历的某个元素arr[i] == 0，直接 top+2 = n+1 了，溢出了，要特殊处理一下
		arr[j] = 0
		j--
		i--
	}
	for j >= 0 {
		arr[j] = arr[i]
		j--
		if arr[i] == 0 {
			arr[j] = arr[i]
			j--
		}
		i--
	}
}
