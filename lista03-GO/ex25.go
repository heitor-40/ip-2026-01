package main

import "fmt"

func main() {
    soma := 0.0
    numerador := 1.0
    denominador := 225.0
    sinal:=1.0

    for ;denominador >= 1;denominador-=25{
        soma += sinal * numerador / denominador
        numerador *= 2
        sinal *= -1
    }
    fmt.Printf("S = %.2f", soma)
}