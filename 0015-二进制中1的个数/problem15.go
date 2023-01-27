package _15_二进制中1的个数

func hammingWeight1(num uint32) int {
	res := 0
	for num > 0 {
		res = res + int(num&1)
		num >>= 1
	}

	return res
}

func hammingWeight2(num uint32) int {
	res := 0
	for num > 0 {
		res = res + 1
		num = num & (num - 1)
	}

	return res
}
