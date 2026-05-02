package main



import "fmt"



func main() {

    var n int



    fmt.Scan(&n)



    fat := 1



    for i := 1; i <= n; i++ {

        fat *= i

    }



    fmt.Printf("Fatorial de %d = %d\n", n, fat)

}