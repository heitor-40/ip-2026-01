package main

import "fmt"

func main() {
    var S float64
    var rea float64
     var Srea float64
    fmt.Scan(&S)

    

    if S <= 300.0 {
        rea = 0.5
    } else {
        rea = 0.3
    }

    Srea= S+(S * rea)

    fmt.Printf("SALARIO COM REAJUSTE = %.2f\n", Srea)
}