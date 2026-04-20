package main

import "fmt"

func main() {
    var r float64

    fmt.Scan(&r)

    if r >= 9.0 {
        fmt.Printf("NOTA = %.1f CONCEITO = A\n", r)
    } else if r >= 7.5 {
        fmt.Printf("NOTA = %.1f CONCEITO = B\n", r)
    } else if r >= 6.0 {
        fmt.Printf("NOTA = %.1f CONCEITO = C\n", r)
    } else {
        fmt.Printf("NOTA = %.1f CONCEITO = D\n", r)
    }
}