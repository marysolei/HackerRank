package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/lonely-integer/problem

func main() {
	a := []int32{1, 2, 3, 3, 5, 2, 1}
	fmt.Println(lonelyinteger(a))
}

//Any number xor'd with itself will give zero.
//Any number xor'd with zero will give the number.
func lonelyinteger(a []int32) int32 {

	var value int32
	for _, element := range a {
		value ^= element
	}
	return value

}

func lonelyinteger2(a []int32) int32 {

	var hmap map[int32]int32
	hmap = make(map[int32]int32)

	for _, e := range a {

		hmap[e]++
	}
	for j, v := range hmap {
		if v == 1 {
			return j
		}
	}
	return 0
}
