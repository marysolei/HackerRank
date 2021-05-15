package main

import (
	"fmt"
	"math"
)

//Question Link: https://www.hackerrank.com/challenges/sum-vs-xor/problem
func main() {
	//a := []int32{1, 2, 3, 3, 5, 2, 1}
	fmt.Println(maximizingXor(10, 15))
}

func maximizingXor(l int32, r int32) int32 {

	maxv := int32(math.MinInt32)
	for i := l; i <= r; i++ {
		for j := i; j <= r; j++ {
			res := i ^ j
			if res > maxv {
				maxv = res
			}
		}
	}
	return maxv
}
