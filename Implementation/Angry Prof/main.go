package main

import (
	"fmt"
)

func main() {

	fmt.Println(angryProfessor(3, []int32{0, -1, 2, 1}))
}

func angryProfessor(k int32, a []int32) string {
	var earlybird int32
	for _, e := range a {
		if e <= 0 {
			earlybird++
		}
	}
	if earlybird >= k {
		return "NO"
	}
	return "YES"

}
