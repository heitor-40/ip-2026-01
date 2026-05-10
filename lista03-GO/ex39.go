package main

import "fmt"

func main() {
    var id, peso, maxID, minID, maxPeso, minPeso int

    fmt.Print("Boi 1-ID:"); fmt.Scan(&id)
    fmt.Print("Boi 1-Peso:"); fmt.Scan(&peso)
    maxID,maxPeso,minID,minPeso = id,peso,id,peso

    for i := 2; i<=90; i++ {
        fmt.Printf("Boi %d - ID: ", i) 
        fmt.Scan(&id)
        fmt.Printf("Boi %d - Peso: ", i)
        fmt.Scan(&peso)
        if peso >maxPeso {
            maxPeso = peso; maxID = id 
        }
        if peso <minPeso {
            minPeso = peso; minID = id
        }
    }

    fmt.Printf("Mais gordo: ID %d, %d kg\n", maxID, maxPeso)
    fmt.Printf("Mais magro: ID %d, %d kg\n", minID, minPeso)
}