package main

import "fmt"
func main() {

    var continuar  int
    var pesMenor40 int
    var qtd50      int
    var qtd10e20   int
    totPessoa := 0
    salturas  := 0.0
    mediaAlt  := 0.0
    for {
        var idade        int

        var altura, peso float64

        fmt.Println("Digite a idade:")

        fmt.Scan(&idade)

        fmt.Println("Digite a altura:")

        fmt.Scan(&altura)

        fmt.Println("Digite o peso:")

        fmt.Scan(&peso)

        totPessoa++

        if idade > 50 {
            qtd50++
        }
        if idade >= 10 && idade <= 20 {
            salturas += altura
            qtd10e20++
        }
        if peso < 40 {
            pesMenor40++
        }
        fmt.Println("Continuar? 1-Sim / outro-Não:")

        fmt.Scan(&continuar)

        if continuar != 1 {
            break
        }
    }
    if qtd10e20 > 0 {
        mediaAlt = salturas / float64(qtd10e20)
    }
    pcem := float64(pesMenor40) / float64(totPessoa) * 100
    fmt.Printf("idade > 50 anos: %d\n", qtd50)
    fmt.Printf("media altura entre 10-20 anos: %.2f\n", mediaAlt)
    fmt.Printf("porcentagem peso < 40kg: %.2f%%\n", pcem)
}