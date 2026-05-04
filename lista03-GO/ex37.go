package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	decimal := 0
	exp := 1
	octal := n 
    soma:=0
	for octal > 0 {
		decimal=octal % 10          
		soma+=decimal*exp 
		exp *=8                     
		octal/=10                  
	}

	fmt.Printf("O numero %d na base 10 e %d\n", n, soma)
}
