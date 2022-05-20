package multiplepointerspattern

// Works for sorted array
// [1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 4]
func CountUnique(a1 []int) int {
	if len(a1) == 0 {
		return 0
	}
	count := 1
	for i, j := 0, 1; i < len(a1)-2; {
		if a1[i] != a1[j] {
			count++
			i = j
			j++
		}
		j++
	}
	return count
}
