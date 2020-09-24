package main

import (
	"fmt"
)

func soma(a int, b int) int {
	return a + b
}

func main() {
	fmt.Printf("A resultado da soma de 5 + 5 é %v \n", soma(5, 5))
}
