package _010_II_青蛙跳台阶

const mod int = 1e9 + 7

type mat [2][2]int

func multiply(a mat, b mat) (c mat) {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			c[i][j] = (a[i][0]*b[0][j] + a[i][1]*b[1][j]) % mod
		}
	}

	return c
}

func pow(m mat, n int) mat {
	res := mat{{1, 0}, {0, 1}}
	for n > 0 {
		if n&1 == 1 {
			res = multiply(res, m)
		}

		n >>= 1
		m = multiply(m, m)
	}

	return res
}

func numWays(n int) int {
	if n < 2 {
		return 1
	}

	res := pow(mat{{0, 1}, {1, 1}}, n-1)
	return (res[1][0] + res[1][1]) % mod
}
