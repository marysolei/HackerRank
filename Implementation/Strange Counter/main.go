package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/strange-code/problem
func main() {

	fmt.Println(strangeCounter(17))
}

func strangeCounter(t int64) int64 {

	n := 3
	for t > int64(n) {
		t = t - int64(n)
		n *= 2
	}
	return int64(n) - t + 1
}
