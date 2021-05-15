package main

//Question Link: https://www.hackerrank.com/challenges/chocolate-feast/problem

import (
	"fmt"
)

func main() {

	fmt.Println(chocolateFeast(15, 3, 2))
}

func chocolateFeast(n int32, c int32, m int32) int32 {
	var ate int32
	chocolate := n / c
	ate += chocolate
	for chocolate >= m {
		ate += chocolate / m
		chocolate = chocolate/m + chocolate%m
	}
	return ate

}
