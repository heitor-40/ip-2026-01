package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    var s float64
    sinal:= 1.0
    numerador:=1000.0

    for i := 1; i <= n; i++ {
        s+=sinal*numerador/float64(i)
        sinal*=-1
        numerador-=3
    }
    fmt.Printf("S=%.2f\n",s)
}