package problem58_I_翻转单词顺序

import "strings"

func reverseWords(s string) string {
	var res string
	var i = len(s) - 1
	var j = i
	for i >= 0 {
		// 找到单词的结尾
		for i >= 0 && s[i] == ' ' {
			i--
		}

		// 找单词的开头
		j = i
		for i >= 0 && s[i] != ' ' {
			i--
		}

		// 添加到result中
		res += s[i+1:j+1] + " "
	}

	// 去除末尾多余的空格
	return strings.TrimRight(res, " ")
}
