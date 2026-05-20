package main

import "fmt"

func main() {
    v:= make([]int, 10)
    for i:= 0;i < 10;i++ {
        fmt.Scan(&v[i])
    }

    for i, num:= range v {
        primo:= num > 1
        for j:= 2; j*j<= num; j++ {
            if num%j== 0 {
                primo= false
                break
            }
        }
        if primo {
            fmt.Printf("Primo %d na posição %d\n", num, i)
        }
    }
}