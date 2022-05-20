package main

import "fmt"

func main() {
	n := subtractProductAndSum(234)
	fmt.Println(n)
}

func subtractProductAndSum(n int) int {
	aList := []int{}
	dSum := 0
	dProd := 1

	if n != 0 {
		for n != 0 {
			aList = append(aList, n%10)
			fmt.Println(aList)
			n /= 10
		}

		for i := range aList {
			dSum += aList[i]
			fmt.Println("dSum -> ", dSum)
			dProd *= aList[i]
			fmt.Println("dProd -> ", dProd)
		}
	}

	return dProd - dSum
}
