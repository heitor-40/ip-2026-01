package main

import "fmt"

func main() {
    var A, B, C float64
    fmt.Scan(&A)
    fmt.Scan(&B)
    fmt.Scan(&C)

    VDelta := B*B - 4*A*C
    fmt.Printf("O VALOR DE DELTA É = %.2f\n", VDelta)
}