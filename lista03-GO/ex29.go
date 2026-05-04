package main

import"fmt"

func main(){
    var n, soma  int
    fmt.Scan(&n)

    for i:=1;i<=n;i++{
        soma+=i
    }

    fmt.Printf("soma=%d\n", soma)
}