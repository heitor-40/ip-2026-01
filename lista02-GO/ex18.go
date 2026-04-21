package main

import "fmt"

func main() {
    var V float64
    var d, t int

    fmt.Scan(&V, &d, &t)

    if t == 2 {
        V *= 1.15
    }

    if d == 2 || d == 3 || d == 5 {
        V *= 0.6
    }

    fmt.Println(V)
}