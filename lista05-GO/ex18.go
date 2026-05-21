package main

import "fmt"

func main() {
    v := make([]int, 10)

    for i := 0; i < 10; i++ {
        var n int
        fmt.Scan(&n)
        
        pos := i
        for j := 0; j < i; j++ {
            if n < v[j] {
                pos = j
                break
            }
        }
        
        for k := i; k > pos; k-- {
            v[k] = v[k-1]
        }
        v[pos] = n
    }

    fmt.Println(v)
}