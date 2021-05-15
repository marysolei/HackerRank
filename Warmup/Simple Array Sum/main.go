package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/simple-array-sum/problem

func main() {
	a := []int32{1, 2, 3, 4, 10, 11}
	fmt.Println(simpleArraySum(a))
}

func simpleArraySum(ar []int32) int32 {
	sum := int32(0)
	for _, e := range ar {
		sum += e
	}
	return sum

}
