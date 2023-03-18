package utils

func Sum(a int, b int) int {
	return a + b
}

func SumAll(n ...int) (c int) {
	for _, v := range n {
		c += v
	}

	return
}
