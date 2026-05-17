package main

import "fmt"

func main() {
    var vetor [10]int

    fmt.Println("Digite 10 números inteiros:")
    for i:=0; i < 10;i++{
        fmt.Scanf("%d",&vetor[i])
    }

    fmt.Print("a) Números pares:")
    for i := 0; i < 10;i++{
        if vetor[i]%2 ==0{
            fmt.Printf("%d ",vetor[i])
        }
    }
    fmt.Println()

    soma:=0
    for i:= 0; i < 10;i++{
        if vetor[i]%2 ==0{
            soma+=vetor[i]
        }
    }
    fmt.Printf("b Soma dos números pares: %d\n",soma)

    fmt.Print("c Números ímpares: ")
    for i := 0; i < 10; i++ {
        if vetor[i]%2 !=0{
            fmt.Printf("%d ",vetor[i])
        }
    }
    fmt.Println()

    qtdImpares:=0
    for i:=0; i < 10;i++{
        if vetor[i]%2 !=0{
            qtdImpares++
        }
    }
    fmt.Printf("d Quantidade de ímpares: %d\n",qtdImpares)
}