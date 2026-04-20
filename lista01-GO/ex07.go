package main
import "fmt"

func main() {
    var fahr, pol,cels,mm float64

    fmt.Scan(&fahr)
    fmt.Scan(&pol)

    cels = (5*fahr - 160) / 9
    mm = pol * 25.4

    fmt.Printf("O VALOR EM CELSIUS = %.2f\n", cels)
    fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", mm)
}