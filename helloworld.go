package main

import "fmt"

const (
	a1, a2 = iota + 1, iota + 2
	a3, a4 = iota + 1, iota + 2
)

func main() {
	fmt.Println("a1: ", a1)
	fmt.Println("a2: ", a2)
	fmt.Println("a3: ", a3)
	fmt.Println("a4: ", a4)
}
