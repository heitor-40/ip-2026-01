package main

import "fmt"

func main() {
    var i int
    
    fmt.Print("Idade:")
    fmt.Scan(&i)
    
    if i < 16 {
        fmt.Println("Não-eleitor")
    } else if i >= 16 && i < 18 {
        fmt.Println("Eleitor facultativo")
    } else if i >= 18 && i <= 65 {
        fmt.Println("Eleitor obrigatório")
    }else{fmt.Println("indique outra idade")
    }
}