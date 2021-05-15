package main

import "fmt"

//Question Link: https://www.hackerrank.com/challenges/permutation-equation/problem

func main() {

	p := []int32{4, 3, 5, 1, 2}
	fmt.Println(permutationEquation(p))
}
func permutationEquation(p []int32) []int32 {
	var value map[int32]int32
	value = make(map[int32]int32)
	var position map[int32]int32
	position = make(map[int32]int32)
	var res []int32
	for i, e := range p {
		value[e] = int32(i + 1)
		position[int32(i+1)] = e
	}
	for j := 1; j <= len(p); j++ {
		t := value[int32(j)]
		y := value[t]
		res = append(res, y)
	}
	return res

}
