package main

import "fmt"

func main() {
    fib:= make([]int, 50)
    fib[0]=1
    fib[1]=1
    
    for i:= 2;i < 50;i++ {
        fib[i]= fib[i-1]+fib[i-2]
    }
    
    for _, v:= range fib {
        fmt.Println(v)
    }
}