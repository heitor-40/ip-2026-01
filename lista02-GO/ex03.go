package main

import "fmt"

func main(){

var n1, n2 int

fmt.Scan(&n1, &n2)

s1:=n1+n2

if(s1>20){
    s2:=s1+8
   fmt.Printf("%d\n", s2)
  }else{ s2:=s1-5
 fmt.Println(s2)
  }
}

