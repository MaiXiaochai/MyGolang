package main

import "fmt"

// 闭包
// 应用场景：将函数fc2作为函数传给fc1，但是不能修改函数fc1
func fc1(f func()) {
	fmt.Println("I'm fc1")
	f()
}

func fc2(x, y int) {
	fmt.Println("I'm fc2")
	fmt.Println(x + y)
}

// 解决方法：将fc2写成闭包，返回一个复合fc1参数要求的函数
func fc3(f func(int, int), x, y int) func() {
	return func() {
		f(x, y)
	}
}

func main() {
	ret := fc3(fc2, 10, 12)
	fc1(ret)
}
