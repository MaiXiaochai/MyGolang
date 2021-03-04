package main

import "fmt"

func arrayInit() {
	// 数组声明和初始化
	var arr1 [3]int
	var arr2 [4]int
	arr3 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	arr4 := [5]int{0, 1}
	arr5 := [5]int{0: 1, 4: 5}

	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", arr2)
	fmt.Printf("%T\n", arr3)
	fmt.Println("arr4: ", arr4)
	fmt.Println("arr5: ", arr5)
}

func throughArray() {
	// 数组遍历

	// 根据索引遍历
	heroes := [3]string{"超人", "蝙蝠侠", "钢铁侠"}
	for i := 0; i < len(heroes); i++ {
		fmt.Println(heroes[i])
	}

	// for range 遍历
	for i, v := range heroes {
		fmt.Println(i, v)
	}
}

func multiArray() {
	// 多维数组

	// 声明
	var arr1 [3][2]int

	// 初始化
	arr1 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(arr1)

	// 遍历
	for _, v := range arr1 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	// output:
	// [-1 2 3]
	// [1 2 3]
	arr2 := [3]int{1, 2, 3}
	arr3 := arr2
	arr2[0] = -1
	fmt.Println(arr2)
	fmt.Println(arr3)
}

func arrayCalc() {
	arr1 := [...]int{1, 3, 5, 7, 9}

	// for 循环法
	sum := 0
	for _, v := range arr1 {
		sum += v
	}
	fmt.Println("数组的和为：", sum)
}

func main() {
	arrayCalc()
}
