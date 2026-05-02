func main() {
    n int
    m bool

    fmt.Scan(&n)

    for i := 1; i * i * i <= n; i++ {
        if i * (i+1) * (i+2) == n {
            m = true
            break
        }
    }

    if m {
        fmt.Println("É triangular")
    } else {
        fmt.Println("Não é triangular")
    }
}