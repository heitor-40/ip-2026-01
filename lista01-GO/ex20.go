package main

import "fmt"

func main() {
    var h, min, s int
    
    fmt.Scan(&h)
    fmt.Scan(&min)
    fmt.Scan(&s)
    
    TS := h*3600 + min*60 + s
    
    fmt.Printf("O TEMPO EM SEGUNDOS É = %d\n", TS)
}