package main

import "fmt"

func arraysoma(lista []float64) float64{
    if len(lista)==0{
        return 0.0
    }else{
     return lista[0]+arraysoma(lista[1:])   
    }
    
}
func main(){
    var arr[10000] float64
    q:=0
    fmt.Println("quantidade de termos:")
    fmt.Scan(&q)
    for i:=0; i<q ; i++{
        fmt.Scan(&arr[i])
    }
    fmt.Printf("a soma da lista é:%.2f\n", arraysoma(arr[:q]))
}