package main

import (
	"fmt"
)

func main() {
	// fmt.Println(sameFCP([]int{1, 2, 3, 2, 5}, []int{9, 1, 4, 4, 25}))
	result := anagrams("anagram", "nagaram")
	fmt.Println(result)
}

func anagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	fmt.Println(s1)
	// s1 = strings.ToLower(s1)
	// s2 = strings.ToLower(s2)

	fc1 := make(map[rune]int)
	fc2 := make(map[rune]int)

	for _, char := range s1 {
		fc1[char] += 1
	}
	for _, char := range s2 {
		fc2[char] += 1
	}

	for key := range fc1 {
		if _, ok := fc1[key]; !ok {
			return false
		}
		if fc1[key] != fc2[key] {
			return false
		}
	}
	return true
}

func sameFCP(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}

	fc1 := make(map[int]int)
	fc2 := make(map[int]int)

	for _, val := range a1 {
		fc1[val] += 1
	}
	for _, val := range a2 {
		fc2[val] += 1
	}

	for key := range fc1 {
		if _, ok := fc2[key*key]; !ok {
			return false
		}
		if fc1[key] != fc2[key*key] {
			return false
		}
	}
	return true
}

func same(a1, a2 []int) bool {
	if len(a1) != len(a2) {
		return false
	}
	// loop a1
	// check square value of a1 is present in a2 looping a2
	// if yes return index of a2 value
	// splice the index value from a2 and continue loop
	idx := -1
	for i := range a1 {
		for j := range a2 {
			if a2[j] == a1[i]*a1[i] {
				idx = j
				break
			}
		}
		if idx == -1 {
			return false
		}
		a2 = append(a2[:idx], a2[idx+1:]...)
	}
	return true
}
