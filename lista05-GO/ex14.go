package main

import "fmt"

func main() {
    v1 :=make([]int, 10) 
    v2 :=make([]int, 10)
    res :=make([]int, 20)

    for i:= 0;i < 10;i++ {
        fmt.Scan(&v1[i])
    }
    for i:= 0;i < 10;i++ {
        fmt.Scan(&v2[i])
    }

    for i:= 0;i < 10;i++ {
        res[i*2]= v1[i]
        res[i*2+1]= v2[i]
    }

    fmt.Println(res)
}