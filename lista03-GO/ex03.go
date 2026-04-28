package main

import "fmt"

func main() {
    var sCar float64
    var sJoa float64
    var meses int
    fmt.Println("salário de Carlos:")
    fmt.Scan(&sCar)

    sJoa = sCar / 3

    for sJoa <= sCar {
        sCar += (sCar * 0.02)
        sJoa += (sJoa * 0.05)
        meses++
    }

    fmt.Printf("Após %d meses João igualou ou ultrapassou Carlos.\n", meses)
}