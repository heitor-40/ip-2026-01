package main

import "fmt"

func main() {
    var n int
    fmt.Scan(&n)

    if n == 0 {
        fmt.Println("Binário: 0")
        return
    }

    resultado:=0
    base:=1

    termo:=n
    for termo>0{
        bit:=termo % 2
        resultado += bit * base
        base*=10
        termo/=2
    }

    fmt.Printf("Binário: %d\n", resultado)
}