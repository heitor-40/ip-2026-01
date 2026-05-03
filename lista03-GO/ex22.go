package main
import "fmt"
func main() {
    var s float64
    for i := 1; i <= 37; i++ {
        a := 38 - i
        b := a + 1
        s += float64(a*b) / float64(i)
    }
    fmt.Printf("S = %.2f", s)
}