package main

import "fmt"

// defer 经典理解题目
// 含有defer和return的函数的执行过程
// 	1. 返回值=x
// 	2. 运行defer
// 	3. RET指令

func f1() int {
	x := 5
	defer func() {
		x++ // 2. x = 6； 3.修改的是x，不是返回值，返回值=5
	}()
	return x // 1. 返回值 = x = 5
}

func f2() (x int) { // 1. 返回值 = x, 这个时候 返回值只是申请了内存，并没有放值
	defer func() {
		x++ // 2. 此时(在命名返回值的情况下)的x就是 "返回值retValue本身",x = x + 1; x = 5，则 x = x + 1 = 6; 返回值 = x = 6
	}()
	return 5
}

func f3() (y int) { // 1. 返回值 = y
	x := 5
	defer func() {
		x++ // 2.这里改变的是x，而 返回值已经申请了内存，并赋值，不影响 返回值； 3.返回值 = 5
	}()
	return x // 1.2 返回值 = y = x = 5
}

func f4() (x int) { // 1.返回值 = x,申请了内存，未赋值
	defer func(x int) {
		x++ // 2.这里的x是func内部自己的x，因为是传入的参数；3. 返回值=5
	}(x)
	return 5 // 1.2 返回值 = x = 5，内存上赋了值
}

// defer 案例2(包含main内部)
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//执行过程
// 1. 遇到第一个defer,先算出defer后跟的函数的内部的函数，使得defer后边的函数为一层函数, 10 1 2 3
// 2. a = 0, 同上理，20 0 2 2
// 3. b = 1, 执行倒数第一个defer后跟的语句, 2 0 2 2
// 4. 1 1 3 3
