package main

import (
	"fmt"
)

func sliceSome() {
	// 切片
	// 声明
	s1 := []int{1, 2, 3}
	s2 := []string{"缺氧", "城市天际线", "戴森球计划", "异星工厂"}
	var s3 []int

	fmt.Println(s1, s2)
	// nil代表空，切片是引用类型，不支持直接比较，只能和nil比较
	fmt.Println(s3 == nil)

	// 切片 长度和容量
	// 自己定义的切片
	fmt.Printf("s1 len: %d, cap: %d\n", len(s1), cap(s1))
	fmt.Printf("s3 len: %d, cap: %d\n", len(s3), cap(s3))

	// 由数组转换的切片
	arr1 := [...]int{1, 2, 3, 4, 5, 6}
	s4 := arr1[0:4]
	s5 := s4[1:2]

	fmt.Println(s4)             // [1, 2, 3, 4]
	fmt.Printf("%T\n", s4)      // []int
	fmt.Printf("%d\n", len(s4)) // 4
	fmt.Printf("%d\n", cap(s4)) // 6
	fmt.Println(s5)

	fmt.Println("原始数组: ", arr1, "长度：", len(arr1))     // 原始数组:  [1 2 3 4 5 6] 长度： 6
	fmt.Printf("S4切片长度：%d，容量：%d\n", len(s4), cap(s4)) // S4切片长度：4，容量：6
	fmt.Printf("S5切片长度：%d，容量：%d\n", len(s5), cap(s5)) // S5切片长度：1，容量：5
}

// 切片构造，make
func makeSlice() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1)) // s1: [0 0 0 0 0], len: 5, cap: 10

	// 只指定长度，容量值默认等于长度值
	s2 := make([]int, 5)
	fmt.Printf("s2: %v, len: %d, cap: %d\n", s2, len(s2), cap(s2)) // s2: [0 0 0 0 0], len: 5, cap: 5

	s3 := make([]int, 0, 10)
	fmt.Printf("s3: %v, len: %d, cap: %d\n", s3, len(s3), cap(s3)) // s3: [], len: 0, cap: 10
}

// 切片赋值
func sliceAssignment() {
	// 切片赋值
	s1 := []int{1, 2, 3}
	s2 := s1            // s2 和 s1都指向了同一个底层数组
	fmt.Println(s1, s2) // [1 2 3] [1 2 3]

	s2[0] = 100
	fmt.Println(s1, s2) // [100 2 3] [100 2 3]

	// 切片遍历
	// 1.for range遍历
	for i, v := range s2 {
		fmt.Println("NO.", i, ": ", v)
	}

	// 2.索引遍历
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}

}

func sliceAppend() {
	s1 := []string{"蜘蛛侠", "钢铁侠", "蝙蝠侠"}
	fmt.Printf("s1 before: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	// s1 before: [蜘蛛侠 钢铁侠 蝙蝠侠], len: 3, cap: 3

	// 切片使用append函数时，必须使用原来的切片接收返回值
	// 原来的底层数组放不下的时候，会把底层数组换一个
	s1 = append(s1, "闪电侠")
	fmt.Printf("s1 after: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	// s1 after: [蜘蛛侠 钢铁侠 蝙蝠侠 闪电侠], len: 4, cap: 6

	s2 := []string{"神奇女侠", "神奇四侠", "绿灯侠"}
	s1 = append(s1, s2...)
	fmt.Printf("s1 append s2: %v, len: %d, cap: %d\n", s1, len(s1), cap(s1))
	// s1 append s2: [蜘蛛侠 钢铁侠 蝙蝠侠 闪电侠 神奇女侠 神奇四侠 绿灯侠], len: 7, cap: 12
}

// 切片拷贝
func sliceCopy() {
	s1 := []int{1, 3, 5}
	s2 := s1 //赋值

	// var s3 []int 这样声明是nil，没有内存空间
	s3 := make([]int, 3, 3) // 做一个有内存的slice
	copy(s3, s1)            // copy

	s2[0] = 10

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

// 删除切片中的元素
func delSliceCell() {
	// 删除s1中索引为1的元素，用 append函数曲线救国
	a1 := [...]int{1, 2, 3, 4, 5}
	s1 := a1[:] // [1 2 3 4]

	//s1 = append(s1[:1], s1[2:3]...) // [1 3 3 4 5]
	s1 = append(s1[:1], s1[2:]...)

	fmt.Println("切片 after：", s1)  // [1 3 4 5]
	fmt.Println("数组：", a1)        // [1 3 4 5 5]
	fmt.Println("s1容量：", cap(s1)) // 4

}

func slicePractice() {
	a := make([]int, 5, 10)
	fmt.Println(a) // [0 0 0 0 0]

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}

	fmt.Println(a) // [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}

func main() {
	slicePractice()
}
