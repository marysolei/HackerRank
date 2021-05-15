package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(sumXor(5))
}

//Question Link: https://www.hackerrank.com/challenges/sum-vs-xor/problem
//there's the carry bit, which addition considers while XOR discards
//so n^x == n+x won't work
func sumXor(n int64) int64 {

	var cnt int64

	for n != 0 {
		if n%2 == 0 {
			cnt++
			n /= 2
		} else if n%2 != 0 {
			n /= 2
		}
	}
	return int64(math.Pow(2, float64(cnt)))
}
