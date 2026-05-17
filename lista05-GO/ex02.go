package main

import "fmt"

func main() {
    var v1 [10]int
    var v2 [5]int

    fmt.Println("Digite 10 números para o primeiro vetor:")
    for i:=0;i < 10; i++{
        fmt.Scanf("%d",&v1[i])
    }

    fmt.Println("Digite 5 números para o segundo vetor:")
    for i :=0; i< 5;i++ {
        fmt.Scanf("%d",&v2[i])
    }

    somaV2:=0
    for i :=0; i <5;i++ {
        somaV2+=v2[i]
    }

    fmt.Println("Primeiro vetor resultante (pares + soma do v2):")
    for i:=0;i < 10;i++ {
        if v1[i]%2 ==0{
            fmt.Printf("%d ",v1[i]+somaV2)
        }
    }
    fmt.Println()

    fmt.Println("Segundo vetor resultante (ímpares + soma do v2):")
    for i:=0;i < 10;i++{
        if v1[i]%2 !=0{
            fmt.Printf("%d ",v1[i]+somaV2)
        }
    }
    fmt.Println()
}