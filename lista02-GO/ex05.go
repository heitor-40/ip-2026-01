package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    if n%5 == 0 {
        fmt.Println("É divisível")
    } else {
        fmt.Println("Não é divisível")
    }
}