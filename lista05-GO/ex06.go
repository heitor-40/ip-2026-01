package main
import "fmt"

func main() {
    
    contagem:= make([]int, 100)

    
    for i:= 0;i < 100;i++{
        contagem[i]= 100 - i
    }

    
    fmt.Println("Lista organizada:")
    for i:= 0;i < 100;i++{
        fmt.Printf("Armário %d: Número %d\n", i, contagem[i])
    }
}