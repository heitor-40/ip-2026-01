package main

import "fmt"

func main() {
    var n, an,de int
    fmt.Scan(&an)
    fmt.Scan(&de)
    fmt.Scan(&n)

    fmt.Print(an,de)

    for i:=3;i<=n;i++{
        var pro int
        if i%2 ==0 {
            pro = an-de
        }else{
            pro=an+de
        fmt.Printf(" %d",pro)
        an=de
        de=pro
    }
}
}