//Question Link: https://www.hackerrank.com/challenges/bon-appetit/problem
package main

import "fmt"

func main() {
	bill := []int32{3, 10, 2, 9}
	bonAppetit2(bill, 1, 7)

}

func bonAppetit2(bill []int32, k int32, b int32) {
	var sum int32
	var item int32
	//var newsum int32
	for i, e := range bill {
		sum += e
		if i == int(k) {
			item = e
		}
	}
	if b == sum/2 {
		fmt.Println(item / 2)
	} else {
		fmt.Println("Bon Appetit")
	}
}

/// version 1
func bonAppetit(bill []int32, k int32, b int32) {
	var sum int32
	var newsum int32
	for _, e := range bill {
		sum += e
	}
	newbill := append(bill[:k], bill[k+1:]...)
	for _, v := range newbill {
		newsum += v
	}
	if b != (newsum / 2) {
		fmt.Println((sum / 2) - (newsum / 2))
	} else if newsum/2 == b {
		fmt.Println("Bon Appetit")
	}
}
