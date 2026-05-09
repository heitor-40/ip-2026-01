package main

import "fmt"

func Binario(n int) []int {
    if n==1 {
        return []int{1} } 
        if n==0 {
        return []int{0}}
    
    listaAnterior :=Binario(n / 2)
    resto := n % 2
    return append(listaAnterior, resto)
}

func main() {
    t:=0
    fmt.Println("digite o tamanho da lista")
    fmt.Scan(&t) 
    numeros:= []int{}
    for i:=0; i<t; i++{
        fmt.Println("numeros:")
        v:=0
        fmt.Scan(&v)
        numeros=append(numeros,v)
    }   
    for _, num := range numeros {
        bin :=Binario(num)
        fmt.Printf("%d em binário: %v\n", num, bin)
    }
}