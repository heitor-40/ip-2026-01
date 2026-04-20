package main

import "fmt"

func main() {
    var alt, are float64

    fmt.Scan(&alt, &are)

    raiz3 := 1.7320508075688772935

    
    AB := (3.0 * raiz3 / 2.0) * (are * are)

    V := (1.0 / 3.0) * AB * alt

    fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n", V)
}