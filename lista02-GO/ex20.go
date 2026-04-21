package main

import "fmt"

func main() {
    var pn float64
    var fp int
    var pf float64
    
    fmt.Println("Preço normal")
    fmt.Scan(&pn)
    
    fmt.Println("condição de pagamento: 1/2/3/4") 
    fmt.Scan(&fp)
    
    if fp == 1 {
        pf = pn * 0.90
        fmt.Printf("R$ %.2f\n", pf)
        
    } else if fp == 2 {
        pf = pn * 0.95
        fmt.Printf("R$ %.2f\n", pf)
        
    } else if fp == 3 {
        pf = pn
        par := pn / 2
        fmt.Printf("2x de R$ %.2f\n", par)
        
    } else if fp == 4 {
        pf = pn * 1.10
        par := pf / 3
        fmt.Printf("3x de R$ %.2f\n", par)
        
    }
    
    fmt.Printf("preço: %.2f\n", pf)
}