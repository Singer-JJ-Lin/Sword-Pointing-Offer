package _65_不用加减乘除做加法

func add(a int, b int) int {
	for b != 0 {
		c := uint(a&b) << 1
		a ^= b
		b = int(c)
	}
	return a
}
