package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/counting-valleys/problem

func main() {

	arr := "UDDDUDUU"
	fmt.Println(countingValleys(9, arr))

}

func countingValleys(n int32, s string) int32 {
	var cnt int32
	level := 0
	for _, e := range s {
		if e == 'U' {
			level++
		} else {
			if level == 0 {
				cnt++
			}
			level--
		}
	}
	return cnt

}
