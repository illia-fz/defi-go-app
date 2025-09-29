package main

import (
    "fmt"
)

func main() {
    deposit := 50.0
    interestRate := 0.05
    interest := deposit * interestRate
    fmt.Printf("Deposited %.2f tokens, interest earned %.2f\n", deposit, interest)
}
