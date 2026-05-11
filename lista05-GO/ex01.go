
package main

import "fmt"

func main() {
    var arr [10]int
    maior50:= false

    fmt.Println("Digite 10 números inteiros:")
    for i := 0;i< 10;i++{
        fmt.Scan(&arr[i])
    }

    fmt.Println("Números maiores que 50 e suas posições:")
    for i:= 0;i <10;i++ {
        if arr[i] > 50 {
            fmt.Printf("Posição %d: %d\n", i+1, arr[i])
            maior50 = true
        }
    }

    if !maior50 {
        fmt.Println("Nenhum número é maior que 50.")
    }
}
