package main

import "fmt"

func judgeFinger(number int) {
	switch number {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("非正常手！")
	}
}

func main() {
	judgeFinger(6)
}
