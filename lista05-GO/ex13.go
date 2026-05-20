package main

import "fmt"

func main() {
    ids:= make([]int, 100) 
    meses:= make([]int, 100)
    total:= 0

    for i:= 0;i < 100;i++ {
        fmt.Scan(&ids[i], &meses[i])
        if ids[i]== 0 && meses[i]== 0 { break }
        total++
    }

    for i:= 0;i < total-1;i++ {
        for j:= i+1; j < total;j++ {
            if meses[i] > meses[j] {
                meses[i], meses[j]= meses[j], meses[i]
                ids[i], ids[j]= ids[j], ids[i]
            }
        }
    }

    for i:= 0;i < 3 && i < total;i++ {
        fmt.Println(ids[i], meses[i])
    }
}