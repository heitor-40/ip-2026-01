package main

import "fmt"

func main() {
    // calcula o primeiro caso (passo=0) fora do laço
    preco     := 6.00
    ingressos := 130.0
    lucro     := ingressos*preco - 300.0

    // usa o primeiro caso como referência inicial
    maxLucro    := lucro
    melhorPreco := preco
    melhorQtd   := ingressos

    fmt.Printf("%-12s %-12s %-12s\n", "Preco(R$)", "Ingressos", "Lucro(R$)")
    fmt.Println("-------------------------------------")
    fmt.Printf("%-12.2f %-12.0f %-12.2f\n", preco, ingressos, lucro)

    // começa do passo 1 porque o passo 0 já foi calculado
    for passo := 1; passo <= 8; passo++ {
        precoInteiro := 600 - passo*60
        preco        = float64(precoInteiro) / 100.0
        ingressos    = 130.0 + float64(passo)*30
        lucro        = ingressos*preco - 300.0

        fmt.Printf("%-12.2f %-12.0f %-12.2f\n", preco, ingressos, lucro)

        if lucro > maxLucro {
            maxLucro    = lucro
            melhorPreco = preco
            melhorQtd   = ingressos
        }
    }

    fmt.Printf("\nLucro maximo: R$ %.2f\n", maxLucro)
    fmt.Printf("Preco ideal:  R$ %.2f\n", melhorPreco)
    fmt.Printf("Ingressos:    %.0f\n", melhorQtd)
}