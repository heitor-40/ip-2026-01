package main

import "fmt"

func main() {
    var nh int
    var VT float32

    fmt.Scan(&nh)

    pc:=nh/3
    hr:=nh%3

    VT=float32(pc*10+hr*5)

    fmt.Printf("O VALOR A PAGAR E = %.2f\n", VT)
}