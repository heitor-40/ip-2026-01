package main

import "fmt"

func main() {

	var n1, n2, n3 float64

	fmt.Println("Insira as notas")
	fmt.Scan(&n1, &n2, &n3)
	media := (n1 + n2 + n3) / 3
	if media >= 6 {
		fmt.Printf("media=%.2f Aprovado\n", media)
	} else {
		fmt.Printf("media=%.2f Reprovado\n", media)
	}
}
