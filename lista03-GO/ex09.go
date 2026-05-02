package main

import "fmt"

func main() {

    var n      int

    var taprov int

    var texam  int

    var treprov int

var somam float64

    fmt.Println("Quantos alunos?")

    fmt.Scan(&n)

    for i := 1; i <= n; i++ {

        var n1, n2 float64

        fmt.Println("Nota 1:")

        fmt.Scan(&n1)

        fmt.Println("Nota 2:")

        fmt.Scan(&n2)

        media := (n1 + n2) / 2

        somam += media

        if media >= 7 {

            fmt.Println("Aprovado")

            taprov++

        } else if media >= 3 {

            fmt.Println("Exame")

            texam++

        } else {

            fmt.Println("Reprovado")

            treprov++

        }

    }

    mclass := somam / float64(n)

    fmt.Printf("Total aprovados:  %d\n", taprov)

    fmt.Printf("Total exame:      %d\n", texam)

    fmt.Printf("Total reprovados: %d\n", treprov)

    fmt.Printf("Média da classe:  %.2f\n", mclass)

}