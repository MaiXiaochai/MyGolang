package main

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

func main() {

}
