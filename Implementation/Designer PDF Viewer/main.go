//Question Link: https://www.hackerrank.com/challenges/designer-pdf-viewer/problem

package main

import (
	"fmt"
)

func main() {

	a := []int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}
	fmt.Println(designerPdfViewer(a, "zaba"))
}

func designerPdfViewer(h []int32, word string) int32 {
	var minv int32
	for _, e := range word {
		tempascii := int(e)
		value := h[tempascii-97]
		if value > minv {
			minv = value
		}
	}
	return minv * int32(len(word))

}
