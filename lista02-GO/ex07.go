package main

import "fmt"

func main() {
    var A, B, C int
    var MENOR, INTER, MAIOR int

    fmt.Scan(&A, &B, &C)

    if A < B && A < C {
        MENOR = A

        if B < C {
            INTER = B
            MAIOR = C
        } else {
            INTER = C
            MAIOR = B
        }

    } else if B < A && B < C {
        MENOR = B

        if A < C {
            INTER = A
            MAIOR = C
        } else {
            INTER = C
            MAIOR = A
        }

    } else {
        MENOR = C

        if A < B {
            INTER = A
            MAIOR = B
        } else {
            INTER = B
            MAIOR = A
        }
    }

    fmt.Println(MENOR, INTER, MAIOR)
}