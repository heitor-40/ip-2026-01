package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)

    if n >= 100 && n <= 999 {
        d := (n / 10) % 10
        fmt.Println(d)
    } else {
        fmt.Println("Número inválido")
    }
}