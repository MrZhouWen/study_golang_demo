package DFS____BFS

/**
电话上得数字 分别对应一下按键，比如给定 "233" 请问有哪些字母组合？

递归 回溯
*/
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

var combinations []string

func letterCombinations(digits string) []string {
	combinations = []string{}
	if len(digits) == 0 {
		return combinations
	}
	var dfs func(int, string)
	//index表示处理到digits的下标，combination表示现有的字符串
	dfs = func(index int, combination string) {
		if index == len(digits) {
			combinations = append(combinations, combination)
			return
		}
		letters, _ := phoneMap[string(digits[index])]
		for i := 0; i < len(letters); i++ {
			dfs(index+1, combination+string(letters[i]))
		}
	}
	dfs(0, "")
	return combinations
}
