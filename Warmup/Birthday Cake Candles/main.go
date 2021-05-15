package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func main() {
	arr := []int32{4, 1, 4, 2, 4}
	fmt.Println(birthdayCake(arr))

}

func birthdayCake(arr []int32) int32 {
	var max int32
	var cnt int32
	for _, e := range arr {
		if e > max {
			max = e

		}
	}
	for _, num := range arr {
		if num == max {
			cnt++
		}
	}
	return cnt
}
