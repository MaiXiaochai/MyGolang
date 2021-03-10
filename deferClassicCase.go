package main

func f1() int {
	x := 5
	defer func() {
		x++ // 修改的是x，不是返回值
	}()
	return x // 5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func main() {

}
