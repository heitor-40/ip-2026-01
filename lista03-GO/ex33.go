package main
import "fmt"
func main() {
    var n1, n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)

    quociente:=0
    resto:=n1
    for resto>=n2{
        resto-=n2
        quociente++
    }

    fmt.Printf("Quociente %d:%d=%d\n",n1,n2,quociente)
    fmt.Printf("Resto%d:%d= %d\n",n1,n2,resto)
}