package 字符串

import (
	"strings"
	"unicode"
)

/**
判断给定的一段字符串是否是 符合要求的 html 代码，标签是否闭合

思路 ： 栈 + 字符串的遍历
T= O(n) S= O(n)
*/
func isVaild(code string) bool {
	tags := []string{} //定义一个栈存放 待闭合的标签
	for code != "" {
		//不是标签元素，直接跳过该 byte
		if code[0] != '<' {
			if len(tags) == 0 {
				return false
			}
			code = code[1:]
			continue
		}
		//边界
		if len(code) == 1 {
			return false
		}
		//三种情况 </xxx>、<![CDATA[****]]>、<xxx>
		if code[1] == '/' { //1、闭合标签 </xxx>
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			if len(tags) == 0 || tags[len(tags)-1] != code[2:j] {
				return false
			}
			tags = tags[0 : len(tags)-1]
			code = code[j+1:]
			if len(tags) == 0 && code != "" {
				return false
			}
		} else if code[1] == '!' { //2、特殊的cdata标签 <![CDATA[****]]>
			if len(tags) == 0 || len(code) < 9 || code[2:9] != "[CDATA[" {
				return false
			}
			j := strings.Index(code, "]]>") //还能把一个字符串整体当成一个'字符'，来确定下标
			if j == -1 {
				return false
			}
			code = code[j+1:]
		} else { //3、一个标签的开始 <xxx>
			j := strings.IndexByte(code, '>')
			if j == -1 {
				return false
			}
			tagName := code[1:j]
			if tagName == "" || len(tagName) > 9 {
				return false
			}
			for _, ch := range tagName {
				if !unicode.IsUpper(ch) {
					return false
				}
			}
			tags = append(tags, tagName)
			code = code[j+1:]
		}
	}
	return len(tags) == 0
}
