package main

import (
    "fmt"
    "math"
)

func main() {
    var S float64
    sinal:=1.0

    for i:=1;i<= 51;i+=2{
        S+=sinal*(1.0/float64(i))
        sinal*=-1
    }
    valor:=math.Cbrt(S * 32)
    fmt.Printf("É aproximado = %.4f\n",valor)
}