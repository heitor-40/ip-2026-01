package main

import "fmt"

func main() {
    var n1, n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    res := 0
    for i:=1;i<=n2;i++{
        res+=n1
    }
    fmt.Printf("%d x %d = %d\n", n1, n2, res)
}