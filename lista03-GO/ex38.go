package main

import "fmt"

func calcularDigito(soma int) int {
    resto := soma % 11
    if resto <2{
        return 0
    }
    return 11 -resto
}

func main() {
    var d int
    var ver1, ver2 int
    soma1 := 0
    soma2 := 0

    fmt.Print("Digite os 11 dígitos separados por espaço: ")

    for i := 0; i < 9; i++ {
        fmt.Scan(&d)
        soma1 += d* (10 - i)
        soma2 += d* (11 - i)
    }

    fmt.Scan(&ver1)
    soma2 +=ver1* 2

    fmt.Scan(&ver2)

    dig1 :=calcularDigito(soma1)
    dig2 :=calcularDigito(soma2)

    if dig1 ==ver1 && dig2 ==ver2 {
        fmt.Println("CPF VÁLIDO!")
    } else {
        fmt.Println("CPF INVÁLIDO!")
    }
}