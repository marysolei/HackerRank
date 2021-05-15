package main

import "fmt"

func main() {
	a := []int32{4, 2, 6, 1, 10}
	fmt.Println(workbook(5, 3, a))
}

//Question link: https://www.hackerrank.com/challenges/lisa-workbook/problem

func workbook(n int32, k int32, arr []int32) int32 {
	pageNum := 1
	var cnt int32

	for _, e := range arr {
		for i := 1; i <= int(e); i++ {
			if i == pageNum {
				cnt++
			}
			if i%int(k) == 0 {
				pageNum++
			}
		}
		if e%k != 0 {
			pageNum++
		}
	}
	return cnt
}
