package main

import "fmt"

func main() {
    var matricula int
    var qhoras float64
    
   
    const sMinimo = 788.00
    const vExtra = 10.00
    
    fmt.Println("matrícula:")
    fmt.Scan(&matricula)
    
    fmt.Println("horas-extras: ")
    fmt.Scan(&qhoras)
    
    sExtra := qhoras * vExtra
    sBruto := 3*sMinimo + sExtra
    
    
    dINSS := float64(0)
    dIR:= float64(0)
    
    if sBruto > 1500 {
        dINSS = sBruto * 0.12
    }
    
    if sBruto > 2000 {
        dIR = sBruto * 0.20
    }
    
    sLiquido:= sBruto - dINSS - dIR
    
    fmt.Printf("Matrícula: %d\n", matricula)
    fmt.Printf(" Salário Bruto:%.2f\n" ,sBruto)
    fmt.Printf("Salário líquido: R$ %.2f\n", sLiquido)
}
    