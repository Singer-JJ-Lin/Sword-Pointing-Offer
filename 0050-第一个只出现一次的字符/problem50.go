package _50_第一个只出现一次的字符

func firstUniqChar(s string) byte {
	m := [26]int{}

	for _, i := range s {
		m[i-'a']++
	}

	for _, i := range s {
		if m[i-'a'] == 1 {
			return byte(i)

		}
	}

	return ' '
}
