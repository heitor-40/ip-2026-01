func main() {
    var n1, n2 int
    fmt.Scan(&n1, &n2)
    if n1%n2 == 0 {
        fmt.Println("É divisível")
    } else {
        fmt.Println("Não é divisível")
    }
}