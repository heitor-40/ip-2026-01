package main

import "fmt"

func main() {
    var nc int
    var t string
    var c float64
    var v float64
fmt.Println( "Número da conta")
fmt.Scan(&nc)
    fmt.Scan(&t, &c)

    if t == "R" {
        v = 5 + c*0.05
    } else if t == "C" {
        if c <= 80 {
            v = 500
        } else {
            v = 500 + (c-80)*0.25
        }
    } else if t == "I" {
        if c <= 100 {
            v = 800
        } else {
            v = 800 + (c-100)*0.04
        }
    }

    fmt.Println("Número da conta:",nc)
    fmt.Println(v)
}