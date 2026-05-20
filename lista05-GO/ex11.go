package main

import (
	"fmt"
	"math"
)

func main() {
	b:= make([]float64, 100)
	var s float64

	for i:=0;i< 100; i++ {
		fmt.Scan(&b[i])
	}

	for i:= 0;i < 50;i++ {
	    sub:=b[i]- b[99-i]
		s+= math.Pow(sub, 3)
	}
	fmt.Println(s)
}