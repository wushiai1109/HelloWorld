package hmath

/**
 * @Author: Bruce
 * @Date: 2021/1/16 10:59 下午
 * @Description:
 */

func combination(m, n int) int {
	if n > m-n {
		n = m - n
	}

	c := 1
	for i := 0; i < n; i++ {
		c *= m - i
		c /= i + 1
	}

	return c
}