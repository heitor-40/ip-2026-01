package main
import "fmt"
func main() {
    var n1,n2 int
    fmt.Scan(&n1)
    fmt.Scan(&n2)

    mmc:=n1
    for mmc%n2 != 0 {
        mmc+=n1
    }
    fmt.Printf("MMC= %d\n",mmc)
}