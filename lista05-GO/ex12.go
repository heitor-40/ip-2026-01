package main

import "fmt"

func main() {
    contagem:= make([]int, 11)

    for i:= 0;i < 15;i++ {
        var n int
        fmt.Scan(&n)
        if n>= 0 && n<= 10 {
            contagem[n]++
        }
    }

    fmt.Println("Nota | Abs | Rel")
    for nota, abs:= range contagem {
        rel:= float64(abs)/15.0
        fmt.Printf("%d    | %d   | %.2f\n", nota, abs, rel)
    }
}