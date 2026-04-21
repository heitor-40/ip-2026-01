package main

import "fmt"

func main() {
    var id int
    var n1, n2, n3, mediaE float64
    var mf float64

    fmt.Println("Id aluno: ")
    fmt.Scan(&id)

    fmt.Println(“Notas aluno:”)
    fmt.Scan(&n1)
    fmt.Scan(&n2)
    fmt.Scan(&n3)

    fmt.Println("Media exercícios:")
    fmt.Scan(&mediaE)

    // Fórmula correta: (n1 + n2 + 2*n3)/7 + mediaE
    mf = (n1 + n2*2 +n3*3+mediaE)/7

    if mf >= 9 {
        fmt.Printf(" Id:%d Media final: %.2f conceito = A APROVADO\n",id ,mf)
    } else if mf >= 7.5 && mf < 9.0 {
        fmt.Printf("Id:%d Media final: %.2f conceito = B APROVADO\n",id, mf)
    } else if mf >= 6.0 && mf < 7.5 {
        fmt.Printf("Id:%d Media final: %.2f conceito = C APROVADO\n", id,mf)
    } else if mf >= 4.0 && mf < 6.0 {
        fmt.Printf("Id:%d Media final: %.2f conceito = D REPROVADO\n", id, mf)
    } else {
        fmt.Printf("Id:%d Media final: %.2f conceito = E REPROVADO\n", id, mf)
    }
}
