package main

import (
    "fmt"
    "math"
)

func main() {
    var n float64
    fmt.Scan(&n)
    
    if n >= 0 {
        rai := math.Sqrt(n)
        fmt.Println(rai)
    } else {
        quad := n * n
        fmt.Println(quad)
    }
}