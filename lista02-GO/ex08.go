package main

import "fmt"

func main() {
    var n int

    fmt.Scan(&n)
    if n > 20 && n < 90 {
        fmt.Println("sim")
    } else {
        fmt.Println("não")
    }
}