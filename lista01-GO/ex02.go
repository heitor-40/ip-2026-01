package main

import (
	"fmt"
)

func main() {
	var ncasos int
	fmt.Scan(&ncasos)
	for laco := 1; laco <= ncasos; laco++ {
		var preço int
		var Pp, Ge, Ar, Ca float64
		fmt.Scan(&preço, &Pp, &Ge, &Ar, &Ca)

		Ppf := (Pp / 100) * float64(preço)
		Gef := (Ge / 100) * float64(preço)
		Arf := (Ar / 100) * float64(preço)
		Caf := (Ca / 100) * float64(preço)

		renda := Ppf*1 + Gef*5 + Arf*10 + Caf*20
		fmt.Printf("A RENDA DO JOGO N.%d.E=%.2f\n", laco, renda)
	}
}
