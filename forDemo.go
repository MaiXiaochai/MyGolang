package main

import (
	"fmt"
	"unicode"
)

// 流程控制

// if的一般用法
func judgeAge(age int) {
	if age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("未成年")
	}
}

// if的特殊用法
func judgeAgePro(age int) {
	// 可在if表达式之前添加一个执行语句，再根据变量值进行判断
	if age := 17; age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("未成年")
	}
}

// 基本for循环
func forNormal(number int) {
	for i := 0; i < number; i++ {
		fmt.Println(i)
	}
}

// 变种for循环，省略结束语句
func forDemo2(number int) {
	var i = 5
	for i < number {
		fmt.Println(i)
	}
}

// for无限循环
func forInfinity() {
	for {
		fmt.Println("123")
	}
}

// for range
func forRange() {
	// 字符串
	s := "你好，maixiaochai"
	for i, v := range s {
		fmt.Println(i, v)
	}
}

// 九九乘法表
func printDoubleNineTable() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d %s %d %s %d \t", j, "X", i, "=", j*i)
		}
		fmt.Println()
	}
}

// goto，跳出多层循环案例
func printDoubleNineTableGoto() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d %s %d %s %d \t", j, "X", i, "=", j*i)
			if j == 2 {
				goto XX
			}
		}
		fmt.Println()
	}
XX:
	fmt.Println("跳出多层循环了")
}

// 判断字符串中的汉字个数
func judgeHan(hans string) int {
	counter := 0
	for _, c := range hans {
		// 打印每个字符串
		fmt.Printf("%c\n", c)
		if unicode.Is(unicode.Han, c) {
			counter++
		}
	}

	return counter
}

func main() {
	words := "maixiaochai,加油！"
	number := judgeHan(words)
	fmt.Println(number)
}
