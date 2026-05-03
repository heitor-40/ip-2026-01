package main

import "fmt"

func main() {
    var x float64

    fmt.Print("Digite X: ")
    fmt.Scan(&x)

    s     := 0.0
    fat   := 1.0
    sinal := -1.0

    for i := 1; i <= 20; i++ {
        fat  *= float64(i)
        s   += sinal * (x / fat)
        sinal *= -1
    }

    s = x + s
    fmt.Printf("Soma: %.2f\n", s)
}