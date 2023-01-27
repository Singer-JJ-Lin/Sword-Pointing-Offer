package _16_数值的整数次方

func abs_int(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func abs_float64(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func myPow(x float64, n int) float64 {
	new_n, new_x := abs_int(n), abs_float64(x)
	res := 1.0
	for new_n > 0 {
		if new_n&1 == 1 {
			res = res * new_x
		}

		new_n >>= 1
		new_x = new_x * new_x
	}

	if x < 0 && n%2 == 1 {
		res = -res
	}

	if n < 0 {
		res = 1 / res
	}
	return res
}
