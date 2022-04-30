package 栈

/**
假设有一个同时存储文件和目录的文件系统。下图展示了文件系统的一个示例：求某一个文件的最长路径: x/y/z.ext
dir 【一级】
⟶ subdir1 【二级】
⟶ ⟶ file1.ext 【三级目录】
⟶ ⟶ subsubdir1
⟶ subdir2
⟶ ⟶ subsubdir2
⟶ ⟶ ⟶ file2.ext
如果是代码表示，上面的文件系统可以写为 "dir\n\t subdir1 \n\t\t file1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext" 。
'\n' 和 '\t' 分别是换行符和制表符。

思路：栈 T= O(n) S= O(n)
利用栈保存当前已遍历路径的长度，栈中元素的个数即为当前路径的深度，栈顶元素即为当前路径的长度。
*/
func lengthLongestPath(input string) (ans int) {
	st := []int{}
	for i, n := 0, len(input); i < n; {
		//检测当前文件的深度
		depth := 1
		for ; i < n && input[i] == '\t'; i++ {
			depth++ //每一个 \t 表示更下一层
		}

		//统计当前文件的长度
		length, isFile := 0, false
		for ; i < n && input[i] != '\n'; i++ {
			if input[i] == '.' {
				isFile = true
			}
			length++
		}
		i++ //跳过换行符 \n

		for len(st) >= depth { //说明当前的文件的层级 比 栈顶的要高，出栈
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			length += st[len(st)-1] + 1
		}

		if isFile {
			ans = max(ans, length)
		} else {
			st = append(st, length)
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
