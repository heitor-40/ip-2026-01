package main

import "fmt"

func main() {
    numeros := make([]int, 10)
    divisores := make([]int, 5)

    for i := 0; i < 10; i++ {
        fmt.Scan(&numeros[i])
    }
    for i := 0; i < 5; i++ {
        fmt.Scan(&divisores[i])
    }

    for i := 0; i < 10; i++ {
        fmt.Printf("Número %d:\n", numeros[i])
        for j := 0; j < 5; j++ {
            if numeros[i]%divisores[j] == 0 {
                fmt.Printf("Divisível por %d na posição %d\n", divisores[j], j)
            }
        }
    }
}