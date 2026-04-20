package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    if n == 0 {
        fmt.Println("nulo")
    } else if n < 0 {
        fmt.Println("negativo")
    } else {
        fmt.Println("positivo")
    }
}