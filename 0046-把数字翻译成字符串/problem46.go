package _46_把数字翻译成字符串

import "strconv"

func translateNum(num int) int {
	str := strconv.Itoa(num)
	a, b := 1, 1
	for i := 1; i < len(str); i++ {
		temp := str[i-1 : i+1]
		if temp <= "25" && temp >= "10" {
			a, b = a+b, a
		} else {
			b = a
		}
	}

	return a
}
