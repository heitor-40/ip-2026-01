package main

import "fmt"

func main() {
    var VL, rai, alt float64
    const pi = 3.14159

    fmt.Scan(&rai)
    fmt.Scan(&alt)

    Al := 2 * pi * rai * alt
    Ac := pi * rai * rai

    VL = (Al + 2*Ac) * 100

    fmt.Printf("O VALOR DO CUSTO E = %.2f\n", VL)
}