package main

import (
"fmt"
"math"
)

func main() {
var A, B, C float64

fmt.Scan(&A, &B, &C)  

delta := B*B - 4*A*C  

if delta < 0 {  
    fmt.Println("RAÍZES IMAGINÁRIAS")  
} else if delta == 0 {  
    r := -B / (2 * A)  
    fmt.Println("RAIZ ÚNICA:", r)  
} else {  
    r := (-B + math.Sqrt(delta)) / (2 * A)  
    r1 := (-B - math.Sqrt(delta)) / (2 * A)  
    fmt.Println("RAÍZES DISTINTAS:", r, xr1)  
}

}