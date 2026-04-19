package main

import "fmt"

func main() {
	var sminimo, qkwr float64

	fmt.Scan(&sminimo, &qkwr)

	valorkw := (0.7 * sminimo) / 100
	valorcons := qkwr * valorkw
	cdesconto := valorcons * 0.9

	fmt.Printf("Custo por Kw:R$%.2f\n", valorkw)

	fmt.Printf("Custo do consumo:R$%.2f\n", valorcons)

	fmt.Printf("Custo com desconto:R$%.2f\n", cdesconto)
}
