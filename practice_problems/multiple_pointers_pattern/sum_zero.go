package multiplepointerspattern

// Works for sorted array
// [-5, -2, -1, 0, 2, 3, 4, 7, 9]
func SumZeroPair(a1 []int) []int {
	j := len(a1) - 1
	for i := 0; i < j; {
		if a1[i]+a1[j] == 0 {
			return []int{a1[i], a1[j]}
		}
		if a1[i]+a1[j] > 0 {
			j = j - 1
		}
		if a1[i]+a1[j] < 0 {
			i = i + 1
		}
		if i >= j {
			break
		}
	}
	return []int{}
}
