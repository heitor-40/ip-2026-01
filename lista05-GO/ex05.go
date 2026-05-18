package main
import "fmt"

func main() {
    
    vetor:= make([]int, 10)
    
    for i:= 0;i < 10; i++{
        fmt.Printf("Digite o %dº número:",i+1)
        fmt.Scan(&vetor[i])
    }
    
    menor:= vetor[0]
    posicao:= 0
    
    for i:= 1;i < 10;i++{
        if vetor[i] < menor {
            menor= vetor[i]
            posicao= i
        }
    }
    
    fmt.Printf("O menor elemento do vetor é %d e sua posição é %d\n", menor, posicao+1)
}