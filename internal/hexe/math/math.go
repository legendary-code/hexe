package math

func Maxi(values ...int) int {
	m := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > m {
			m = values[i]
		}
	}

	return m
}

func Mini(values ...int) int {
	m := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] < m {
			m = values[i]
		}
	}

	return m
}

func Absi(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
