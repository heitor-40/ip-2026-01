package main

import "fmt"

func main() {
    var a1, r, n int
    fmt.Scan(&a1, &r, &n)

    TA:= a1
    S:=0
    for i := 0; i < n; i++ {
        S+=TA
        TA += r
    }

    fmt.Println(S)
}