package main

import (
	"fmt"
	"strings"
)

const (
	a1, a2 = iota + 1, iota + 2
	a3, a4 = iota + 1, iota + 2
)

func main() {
	s1 := "D:\\data\\maixiaochai"
	s2 := "Hi, I'm Tom."
	s3 := `D:/data/maixiaochai
			Hi, I'm Tom.`
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	//
	//fmt.Println(s1 + s2)
	//result := fmt.Sprintf("%s%s", s1, s2)
	//fmt.Println(result)

	spl := strings.Split(s1, "\\")
	fmt.Println(spl)
	//fmt.Println("a1: ", a1)
	//fmt.Println("a2: ", a2)
	//fmt.Println("a3: ", a3)
	//fmt.Println("a4: ", a4)

	s4 := "Hello汤姆"
	fmt.Println(len(s4))

	//for i := 0; i < len(s4); i++ {
	//	//fmt.Println(s4[i])
	//	//fmt.Printf("%c\n", s4[i])
	//}

	for _, c := range s4 {
		fmt.Printf("%c\n", c)
	}

	// 字符串修改
	s5 := "水果"
	s51 := []rune(s5) // 把字符串强制转换成一个rune切片
	s51[0] = '苹'
	fmt.Println(string(s51)) // 把rune切片强制转换成字符串

	// 类型转换
	n1 := 10 // int
	var f float64
	f = float64(n1)
	fmt.Printf("%T", f)
}
