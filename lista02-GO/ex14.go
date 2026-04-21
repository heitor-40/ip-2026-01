package main

import "fmt"

func main() {
    var V float64
    var imp int

    fmt.Scan(&V)

    for {
        fmt.Scan(&imp)

        if imp == 0 {
            break
        }

        if imp == 1 {
            V += 1750
        } else if imp == 2 {
            V += 800
        } else if imp == 3 {
            V += 1200
        } else if imp == 4 {
            V += 2000
        }
    }

    fmt.Println(V)
}