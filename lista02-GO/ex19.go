package main

import (
    "fmt"
    "math"
)

func main() {
    var f int
    var r, a float64
    
    fmt.Println("1 - Cone")
    fmt.Println("2 - Cilindro")
    fmt.Println("3 - Esfera")
    fmt.Scan(&f)
    
    if f == 1 {
     
            fmt.Scan(&r)
            fmt.Scan(&a)
        
        V := (math.Pi * r * r * a) / 3
        Ar := math.Pi * r * math.Sqrt(r*r+a*a)
        
        fmt.Printf("Volume cone: %.2f\n", V)
        fmt.Printf("Área cone: %.2f\n", Ar)
        
    } else if f == 2 {
       
        fmt.Scan(&r)
        fmt.Scan(&a)
        
        V := math.Pi * r * r * a
        Ar := 2 * math.Pi * r * a
        
        fmt.Printf("Volume cilindro: %.2f\n", V)
        fmt.Printf("Área cilindro: %.2f\n", Ar)
        
    } else if f == 3 {
       
        fmt.Scan(&r)
        
        V := (4.0 / 3.0) * math.Pi * r * r * r
        Ar := 4 * math.Pi * r * r
        
        fmt.Printf("Volume esfera: %.2f\n", V)
        fmt.Printf("Área esfera: %.2f\n", Ar)
        
    } 
}