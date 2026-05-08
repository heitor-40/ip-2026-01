package main

import"fmt"

func inverso(lista []int) []int{
    if len(lista)<=1{
        return lista
    }
    primeiro:=lista[0]
    return append(inverso(lista[1:]),primeiro)
}

func main(){
    var t int 
    fmt.Println("digite o tamanho:")
    fmt.Scan(&t)
    l:= make([]int,t)
    for i:=0; i<t; i++{
        fmt.Scan(&l[i])
    }
    a:=inverso(l)
    fmt.Println("lista invertida:",a )
}