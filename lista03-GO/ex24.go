package main
import "fmt"

func main() {
    fmt.Printf("%-10s %-15s\n", "A (rad)", "Sen(A)")
    fmt.Println("-------------------------")

    for a := 0.0; a <= 6.3; a += 0.1 {
        a2 := a * a
        a3 := a2 * a
        a5 := a3 * a2
        a7 := a5 * a2

        senA := a - a3/6.0 + a5/120.0 - a7/5040.0
        fmt.Printf("%-10.1f %-15.6f\n", a, senA)
    }
}