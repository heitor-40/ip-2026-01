package main

import "fmt"

func main(){
    var s float64
    fat:=1.0
    for i:=0;i<20;i++{
        if i>0{
            fat*= float64(i)
        }
        s+=(100-float64(i))/fat
    }
    
    fmt.Printf("Soma dos 20 primeiros termos = %.4f\n",s)
}