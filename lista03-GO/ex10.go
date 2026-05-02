package main

import "fmt"

func main() {

    var n int


    fmt.Scan(&n)

    pri := 0

    seg := 1

    fmt.Print(pri, " ", seg)

    for i := 2; i < n; i++ {

        pro := pri + seg

        fmt.Printf(" %d", pro)

        pri = seg

        seg = pro

    }

}