package main
import "fmt"

func main() {
 var n int 
 
 for 1==1{
     fmt.Scan(&n)
     if n<=0{
         break
     }
 }
 quad=false
 for i:=1;i<=n;i++{
     if i*i==n{
         quad=true
         break
     }
 }
 if quad{
     fmt.Println("Quadrado perfeito")
 }else{fmt.Println("Não é quadrado perfeito")
 
}
}
