package main

import "fmt"

func main(){
	var x float64
	fmt.Scan(&x)

	cos := 1.0
	sinal := -1.0

	for i:=1;i<=20;i++{
		exp:=2*i
		fat:=1.0
		for j:=1;j<=exp;j++{
			fat*=float64(j)
		}
		xpo := 1.0
		for j:=0;j<exp;j++{
			xpo*=x
		}
		cos += sinal *(xpo/fat)
		sinal *=-1
	}

	fmt.Printf("Cosseno calculado: %.6f\n", cos)
}