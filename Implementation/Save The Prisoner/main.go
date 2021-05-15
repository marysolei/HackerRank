package main

import (
	"fmt"
)

// Question Link: https://www.hackerrank.com/challenges/save-the-prisoner/problem

func main() {

	fmt.Println(saveThePrisoner(7, 19, 2))
}

func saveThePrisoner(n int32, m int32, s int32) int32 {

	target := (s + m - 1) % n
	if target == 0 {
		return n
	}
	return target
}
