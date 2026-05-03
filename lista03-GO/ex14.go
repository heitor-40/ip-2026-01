package main

import "fmt"

func main() {
    var n1, n2 int

    fmt.Scan(&n1)
    fmt.Scan(&n2)

    for i := n1; i <= n2; i++ {
        primo := true
        if i < 2 {
            primo = false
        }
        for j := 2; j*j <= i; j++ {
            if i%j == 0 {
                primo = false
                break
            }
        }
        if primo {
            fmt.Println(i)
        }
    }
}