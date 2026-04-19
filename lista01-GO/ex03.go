package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 int32
	fmt.Scan(&num1, &num2, &num3)

	if num1 <= 0 || num1 > 9 || num2 <= 0 || num2 > 9 || num3 <= 0 || num3 > 9 {
		fmt.Println("DIGITO INVÁLIDA.")

		return
	}
	numtot := num1*100 + num2*10 + num3

	quadtot := numtot * numtot

	fmt.Printf(" É %d,  %d\n", numtot, quadtot)
}
