package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/sock-merchant/problem

func main() {

	arr := []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	fmt.Println(sockMerchant(9, arr))

}

func sockMerchant(n int32, ar []int32) int32 {
	var cnt int32
	var hmap map[int32]int32
	hmap = make(map[int32]int32)
	for _, e := range ar {
		hmap[e]++
	}
	for _, c := range hmap {
		cnt += (c / 2)
	}
	return cnt
}
