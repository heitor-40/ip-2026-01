package main
import "fmt"
func main(){
    for i:=1;i<=10;i++{
        for j:=0;j<=10;j++{
            if j>i{
                fmt.Printf("%d,%d  ", i, j)
            }
        }
    }
}