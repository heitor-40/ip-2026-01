package main

import "fmt"

func main() {
    ori := make([]int, 30)
    nov :=make([]int, 30)
    
    for i:= 0;i < 30;i++{
        fmt.Scan(&ori[i])
    }

    for i:=0;i < 30;i++{
        if i%2== 0{
            nov[i] = ori[i]*2
        } else{
            nov[i]= ori[i]*3
        }
    }

    fmt.Println(nov)
}