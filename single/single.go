package single

func CountMultiples(multipleOf int64, array []int64) int64 {
	var sum int64 = 0
	for _, v := range array {
		if v%multipleOf == 0 {
			sum++
		}
	}
	return sum
}
