package main

import "fmt"

func theMap() {
	// 声明
	var m1 map[string]int // 还没有初始化(没有在内存中开辟空间)

	// make 初始化，10：容量
	//可以自动扩容，但是最好在最开始就估算好要使用的容量，避免中途动态扩容导致的性能下降
	m1 = make(map[string]int, 10)

	m1["hello"] = 666
	m1["python"] = 777

	fmt.Println("m1: ", m1)                   // m1:  map[hello:666 python:777]
	fmt.Println("m1['hello']: ", m1["hello"]) // m1['hello']:  666

}

func main() {
	theMap()
}
