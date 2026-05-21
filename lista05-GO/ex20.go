package main

import "fmt"

func main() {
    jogadas:= make([]int,20)
    frequencia:= make([]int,7)

    for i:= 0; i < 20; i++ {
        fmt.Scan(&jogadas[i])
        n:= jogadas[i]
        frequencia[n]++
    }

    fmt.Println("Resultados sorteados:",jogadas)

    for i:= 1;i <= 6;i++ {
        fmt.Printf("O número %d saiu %d vezes\n", i, frequencia[i])
    }
}