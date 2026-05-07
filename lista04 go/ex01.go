
package main
import "fmt"
 func quad(x,n int)int{
     if (n==0){
         return 1
     }
     return x* quad(x,n-1)
 }
func main(){
    var x, n int
    
    fmt.Scan(&x, &n)
    
    fmt.Printf("%d^%d:%d\n",x,n,quad(x,n))
}
