package main

import "fmt"

func main() {
    var soma       int
    var qtdpares   int
    var quant      int
    var qtdimpares int
    var maiorn     int
    var menorn     int
    var somap      int
    var n          int

    fmt.Println("Digite n (30000 para sair):")
    fmt.Scan(&n)

    maiorn = n
    menorn = n
    somap  = 0

    for {
        if n == 30000 {
            break
        }

        soma  += n
        quant++

        if n > maiorn {
            maiorn = n
        }
        if n < menorn {
            menorn = n
        }

        if n%2 == 0 {
            somap += n
            qtdpares++
        } else {
            qtdimpares++
        }

        fmt.Println("Digite n (30000 para sair):")
        fmt.Scan(&n)
    }

    media  := float64(soma) / float64(quant)

    mediap := 0.0
    if qtdpares > 0 {
        mediap = float64(somap) / float64(qtdpares)
    }

    pcem := float64(qtdimpares) / float64(quant) * 100

    fmt.Printf("Soma: %d\n",                    soma)
    fmt.Printf("Quantidade: %d\n",              quant)
    fmt.Printf("Média: %.2f\n",                 media)
    fmt.Printf("Maior numero: %d\n",            maiorn)
    fmt.Printf("Menor numero: %d\n",            menorn)
    fmt.Printf("Media pares: %.2f\n",           mediap)
    fmt.Printf("Porcentagem impares: %.2f%%\n", pcem)
}