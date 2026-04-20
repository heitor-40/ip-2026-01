package main

import "fmt"

func main() {
    var a, b, c, d float64

    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    fmt.Scan(&d)

    M := a*d - b*c

    fmt.Printf("O VALOR DO DETERMINANTE É = %.2f\n", M)
}