package main

import "fmt"

func show1() {
	fmt.Println("maixiaochai")
}

func callShow(f func()) {
	f()
}

func main() {
	callShow(show1)
}
