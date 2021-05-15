package main

import "fmt"

func main() {
	apples := []int32{2, 3, -4}
	oranges := []int32{3, -2, -4}
	countApplesAndOranges(7, 11, 5, 15, apples, oranges)

}

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	resapple := 0
	resorange := 0
	for _, e := range apples {
		e += a
		if s <= e && e <= t {
			resapple++
		}
		//resapple = append(resapple, e)
	}
	for _, f := range oranges {
		f += b
		if s <= f && f <= t {
			resorange++
		}
	}
	fmt.Println(resapple)
	fmt.Println(resorange)

}
