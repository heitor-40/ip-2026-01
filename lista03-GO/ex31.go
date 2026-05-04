package main

import "fmt"
func main() {
    var total uint64
    graos:=1
    for i:= 1;i<=64;i++ {
        total+=uint64(graos)
        fmt.Printf("Quadrado %2d  %d grãos\n",i,uint64(graos))
            graos*=2
    }
    fmt.Printf("Total: %d\n", total)
}