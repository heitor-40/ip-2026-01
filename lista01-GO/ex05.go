package main

import "fmt"

func main() {

    var Ccliente int
    var Tconsumidor string
    var Cagua float32
    var valorc float32

    fmt.Scan(&Ccliente, &Tconsumidor, &Cagua)

    if Tconsumidor == "R" {
        valorc = 5 + 0.05*Cagua

    } else if (Tconsumidor == "C" && Cagua <= 80) {
        valorc = 500

    } else if (Tconsumidor == "C" && Cagua > 80) {
        valorc = 500 + (Cagua-80)*0.25

    } else if (Tconsumidor == "I" && Cagua <= 100) {
        valorc = 800

    } else if (Tconsumidor == "I" && Cagua > 100) {
        valorc = 800 + (Cagua-100)*0.04
    }

    fmt.Printf("CONTA: %d %s\n", Ccliente, Tconsumidor)
    fmt.Printf("VALOR DA CONTA: %.2f\n", valorc)
}