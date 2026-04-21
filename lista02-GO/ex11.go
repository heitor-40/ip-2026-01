package main

import "fmt"

func main() {
    var x float64
    var r float64

    fmt.Scan(&x)

    if x == 2 {
        fmt.Println("Sem resultado")
    } else {
        r = 8 / (2 - x)
        fmt.Println(r)
    }
}