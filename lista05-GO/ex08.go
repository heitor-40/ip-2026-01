package main

import (
	"fmt"
	"math"
)

func main() {
	vetor := make([]float64, 15)

	for i:=0;i < 15;i++ {
		var n float64
		fmt.Scan(&n)

		if n<0 {
			vetor[i]= -1
		} else {
			vetor[i]= math.Sqrt(n)
		}
	}

	for _, v := range vetor {
		fmt.Println(v)
	}
}