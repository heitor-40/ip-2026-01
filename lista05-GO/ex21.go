package main

import "fmt"

func main() {
    vetor:= make([]float64, 10)

    for i:= 0;i < 10;i++ {
        fmt.Scan(&vetor[i])
    }

    var codigo int
    fmt.Scan(&codigo)

    if codigo== 0 {
        return
    } else if codigo== 1 {
        for i:= 0;i < 10;i++ {
            fmt.Printf("%.2f ", vetor[i])
        }
    } else if codigo== 2{
        for i:= 9;i >= 0;i--{
            fmt.Printf("%.2f ", vetor[i])
        }
    }
    fmt.Println()
}