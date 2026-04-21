package main

import "fmt"

func main() {
    var C, Vvenda float64
    fmt.Scan(&C)

    if C< 10 {
        Vvenda = C + (C * 0.70)

    } else if C >= 10 && C < 30 {
        Vvenda = C + (C * 0.50)

    } else if C >= 30 && C < 50 {
        Vvenda = C+ (C * 0.40)

    } else {
        Vvenda = C + (C * 0.30)
    }

    fmt.Printf("valor da venda=%f", Vvenda)
}