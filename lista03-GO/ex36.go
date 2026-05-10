
package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println("Hexadecimal: 0")
        return
    }

    c := 0
    t := n
    for t > 0 {
        c++
        t /= 16
    }

    fmt.Print("Hexadecimal: ")
    for i := c - 1; i >= 0; i-- {

        pot:=1
        for j:=0;j<i;j++ {
            pot*=16
        }

        digito:=(n/pot)%16

        if digito<10{
            fmt.Print(digito)
        }else{
            fmt.Printf("%c", 'A'+digito-10)
        }
    }
}