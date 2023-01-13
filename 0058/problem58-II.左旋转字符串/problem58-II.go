package problem58_II_左旋转字符串

func reverseLeftWords(s string, n int) string {
	charArray := []byte(s)
	reverse(charArray, 0, n-1)
	reverse(charArray, n, len(s)-1)
	reverse(charArray, 0, len(s)-1)
	return string(charArray)
}

func reverse(s []byte, l, r int) {
	for ; l < r; l, r = l+1, r-1 {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp
	}
}
