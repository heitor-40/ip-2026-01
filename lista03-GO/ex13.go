package main

import "fmt"

func main() {
    h         := 0.0
    numerador := 1.0

    for i := 1; i <= 50; i++ {
        h+= numerador/float64(i)
        numerador+=2
    }

    fmt.Printf("H = %.2f\n",h)
}