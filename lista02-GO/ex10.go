package main

import "fmt"

func main() {
var D int
var I int
var V float64

fmt.Scan(&D)  

fmt.Scan(&I)  

if D == 1 {  
    if I == 1 {  
        V = 500  
    } else {  
        V = 900  
    }  

} else if D == 2 {  
    if I == 1 {  
        V = 350  
    } else {  
        V = 650  
    }  

} else if D == 3 {  
    if I == 1 {  
        V = 350  
    } else {  
        V = 600  
    }  

} else if D == 4 {  
    if I == 1 {  
        V = 300  
    } else {  
        V = 550  
    }  
}else{fmt.Println("Destino invalido")
return
}
fmt.Println(V)
}