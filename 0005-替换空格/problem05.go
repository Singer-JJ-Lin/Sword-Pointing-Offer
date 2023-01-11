package _05_替换空格

func replaceSpace(s string) string {

	// 统计空格的数量
	count := 0
	for _, c := range s {
		if string(c) == " " {
			count++;
		}
	}

	// 创建一个byte切片，长度为s原本的长度 + 2倍count
	result := make([]byte, len(s)+count*2)
	p1 := len(s) - 1
	p2 := len(result)

	// 从后往前遍历s1
	for p2 > p1 && p1 >= 0 {
		// 遇到空格，则从后往前向resul中填充%20
		if string(s[p1]) == " " {
			p2 -= 3
			result[p2] = byte('%')
			result[p2+1] = byte('2')
			result[p2+2] = byte('0')
			// 不是空格，则直接copy到result
		} else {
			p2--
			result[p2] = s[p1]
		}
		p1--
	}

	// 返回结果
	return string(result)
}