package main

import "fmt"

func main() {
    var I int

      fmt.Scan(&I)

    if I >= 0 && I <= 2 {
        fmt.Println("Recém-nascido")

    } else if I >=3 &&  I<= 11 {
        fmt.Println("Criança")

    } else if I>=12 && I<= 19 {
        fmt.Println("Adolescente")

    } else if I>=20 && I<= 55 {
        fmt.Println("Adulto")

    } else {
        fmt.Println("Idoso")
    }
}