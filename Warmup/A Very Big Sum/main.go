package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/a-very-big-sum/problem

func main() {
	a := []int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}
	fmt.Println(aVeryBigSum(a))
}

func aVeryBigSum(ar []int64) int64 {

	var res int64
	for i := 0; i < len(ar); i++ {
		res += ar[i]
	}
	return res
}
