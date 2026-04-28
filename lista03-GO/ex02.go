package main

import "fmt"

func main() {
    var quant int
    var soma int
    var media float64

    for i := 50; i <= 70; i++ {
        if i%2 == 0 {
            soma += i
            quant++
        }
    }

    media = float64(soma) / float64(quant)

    fmt.Printf("Soma=%d\n Media=%.2f\n", soma, media)
}