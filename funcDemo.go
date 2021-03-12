package main

import "fmt"

func mySum(x int, y int) int {
	return x + y
}

func mySum2(x int, y int) {
	fmt.Println(x + y)
}

// 命名的返回值ret，ret int 表示声明变量
// 如果使用的命名返回值，return后边可以不写ret
func mySum3(x int, y int) (ret int) {
	ret = x + y
	return
}

// 多个返回值
func mySum4(x int, y int) (string, int) {
	return "hello", x + y
}

// 参数简写，连续多个参数类型相同，同类型的参数除最后一个外，其它的类型关键字可以省略
func mySum5(x, y, z int, i, j, k string) (string, int) {
	return i + j + k, x + y - z
}

// 可变长参数，y为slice类型，这里为[]int
// return: maixiaochai: []int, [1 2 3]
func mySumPlus(x string, y ...int) string {
	return fmt.Sprintf("%s: %T, %v", x, y, y)
}

// 判断是否是回文
func judgePlalindrome(words string) {
	fmt.Println(len(words)) // 27

	// 把字符串拿出来，放到一个[]rune中
	r := make([]rune, 0, len(words))
	for _, c := range words {
		r = append(r, c)
	}
	fmt.Println(r) // [19978 28023 33258 26469 27700 26469 33258 28023 19978]

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}

// defer
func show() {
	fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	// Out: 1 4 3 2
}

// 匿名函数
var fn = func(x, y int) {
	fmt.Println(x + y)

	// 如果只是调用一次的函数，还可以简写成立即执行函数
	func() {
		fmt.Println("立即执行函数")
	}()

	// 立即执行函数，带参数
	func(x string) {
		fmt.Printf("立即执行函数, %s\n", x)
	}("done.")

}

func main() {
	fn(1, 2)
}
