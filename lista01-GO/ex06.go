package main
import "fmt"

func main() {
   var n int
   var Vfahr float64
   var Vcels float64

   fmt.Scan(&n)

   for i := 0; i < n; i++ {
       fmt.Scan(&Vfahr)
       Vcels = (5 * (Vfahr - 32)) / 9
       fmt.Printf("%.2f FAHRENHEIT EQUIVALE A %.2f CELSIUS\n", Vfahr, Vcels)
   }
}