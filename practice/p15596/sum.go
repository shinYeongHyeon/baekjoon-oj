package p15596

func Sum(a []int) int {
	var sum int
	for _, val := range a {
		sum += val
	}
	return sum
}
