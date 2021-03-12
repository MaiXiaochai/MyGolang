package main

import "fmt"

// 闭包
func f1(f func()) {
	fmt.Println("I'm f1")
	f()
}

func f2(x, y int) {
	fmt.Println("I'm f2")
	fmt.Println(x + y)
}

func main() {

}
