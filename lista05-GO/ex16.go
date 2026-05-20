package main

import "fmt"

func main() {
    idades:= make([]int, 50)
    freq:= make([]int, 121)

    for i:= 0;i < 50;i++ {
        fmt.Scan(&idades[i])
        freq[idades[i]]++
    }

    moda:=0
    max := 0
    for idade, cont := range freq {
        if cont > max {
            max = cont
            moda = idade
        }
    }

    fmt.Printf("Moda: %d (frequência: %d)\n", moda, max)
}