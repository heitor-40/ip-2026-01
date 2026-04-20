package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    
    if n <= 1 {
        fmt.Println("Numero invalido!")
        return
    }
    
    S:=0.
    for k := 1; k <= n; k++ {
       S+=1.0/float64(k)
    }
    
    fmt.Printf("%f\n", S)
}