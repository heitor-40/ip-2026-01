package main

import "fmt"

func main() {
    vetor := make([]int, 100)
    impar := 1

    for i := 0; i < 100; i++ {
        vetor[i] = impar
        impar += 2
    }

    for i := 0; i < 100; i++ {
        fmt.Println(vetor[i])
    }
}