package main

import "fmt"

func main() {
    alturas := make([]float64, 10)
    var soma float64

    for i := 0; i < 10; i++ {
        fmt.Scan(&alturas[i])
        soma += alturas[i]
    }

    media := soma / 10
    fmt.Println("media:", media)

    for _, h := range alturas {
        if h > media {
            fmt.Println(h)
        }
    }
}