package main

import "fmt"

func goPinter() {
	// &，取地址
	n := 666
	p := &n

	// *，根据地址取值
	m := *p

	fmt.Printf("值：%v，地址：%v，地址类型：%T \n", n, p, p) // 值：666，地址：0xc0000ac058，地址类型：*int
	fmt.Printf("值：%v\n", m)                      // 值：666

}

func pointerPro() {
	// 声明了一个int类型的指针
	var a1 *int     // 空指针
	fmt.Println(a1) // <nil>

	// new
	var a2 = new(int)
	fmt.Printf("%T\n", a2) // *int
	fmt.Println(*a2)       // 0

	*a2 = 100
	fmt.Println(a2)  // 0xc00000a0d0
	fmt.Println(*a2) // 100
}

func main() {
	pointerPro()
}
