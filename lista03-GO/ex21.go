package main

import "fmt"

func main() {
    var b,n int
    fmt.Scan(&b)
    fmt.Scan(&n)
    r:=1
    for i:=1;i<=n;i++ {
        r*=b
    }
    fmt.Printf("%d^%d=%d",b,n,r)
}