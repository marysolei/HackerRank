package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	a := "chillout"
	fmt.Println(encryption(a))
}

func encryption(s string) string {
	var str string
	//remove spaces from string
	t := strings.Replace(s, " ", "", -1)
	l := len(t)
	index := 0
	//floor and ceil of the value
	floor := int(math.Floor(math.Sqrt(float64(l))))
	ceiling := int(math.Ceil(math.Sqrt(float64(l))))

	for floor*ceiling < l {
		floor++
	}

	for i := 0; i < ceiling; i++ {
		for j := 0; j < floor; j++ {
			index = i + j*ceiling
			if index >= l {
				break
			}

			str = str + string(s[index])
		}

		str = str + " "
	}
	return str
}
