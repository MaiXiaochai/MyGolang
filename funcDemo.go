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

func main() {
	//r1, r2 := mySum5(1, 2, 3, "mai", "xiao", "chai")
	result := mySumPlus("maixiaochai", 1, 2, 3)
	fmt.Println(result)
}
