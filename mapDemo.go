package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func theMap() {
	// 声明
	var m1 map[string]int // 还没有初始化(没有在内存中开辟空间)

	// make 初始化，10：容量
	//可以自动扩容，但是最好在最开始就估算好要使用的容量，避免中途动态扩容导致的性能下降
	m1 = make(map[string]int, 10)

	// 键值方式去访问,这里是赋值
	m1["hello"] = 666
	m1["python"] = 777

	fmt.Println("m1: ", m1)                   // m1:  map[hello:666 python:777]
	fmt.Println("m1['hello']: ", m1["hello"]) // m1['hello']:  666

	// 这里v代表取到的值，ok是布尔类型（约定成俗，用ok接受布尔值），代表是否存在此值
	v, ok := m1["go"]

	if !ok {
		fmt.Println("没有这个key")
	} else {
		fmt.Println(v)
	}

}

func throughMap() {
	var m1 map[string]int
	m1 = make(map[string]int, 10)

	m1["hello"] = 666
	m1["python"] = 777

	// 遍历键值
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// 删除 hello的键值
	delete(m1, "hello")

	// 只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
}

func mapSortedFetch() {
	rand.Seed(time.Now().UnixNano()) // 初始化随机种子(这里用的纳秒)
	var scoreMap = make(map[string]int, 100)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("mai%02d", i) // 生成mai开头的字符串,fmt.Sprintf()返回字符
		value := rand.Intn(100)          // 生成一个0~99的随机数
		scoreMap[key] = value
	}

	// 取出所有key存入切片
	keys := make([]string, 0, 100) // 声明长度为0， 容量为100的切片
	for key := range scoreMap {
		keys = append(keys, key)
	}

	// 对切片进行排序
	sort.Strings(keys)

	// 按照排序好的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

// map 和 slice组合
func mapSlice() {
	// 1.元素为map类型的切片
	s1 := make([]map[string]int, 10, 10)

	// 对使用到的内部元素(这里是map)进行初始化
	s1[0] = make(map[string]int, 1)
	s1[0]["hello"] = 150
	fmt.Println(s1)
	// [map[hello:150] map[] map[] map[] map[] map[] map[] map[] map[] map[]]

	// 2.值为slice的map
	s2 := make(map[string][]int, 5)

	//s2["hello"] = make([]int, 1, 1) // 可以用make初始化
	s2["hello"] = []int{10, 20, 30} // 切片初始化的一种方式
	s2["hello"][0] = 100
	fmt.Println(s2) // map[hello:[100 20 30]]
}

// 统计单词出现次数
func countWords(words string) {
	s := strings.Split(words, " ")
	m := make(map[string]int, len(s))

	fmt.Printf("%v, %T, %v", m, m, len(m))

	for _, nbr := range s {
		if _, ok := m[nbr]; !ok {
			m[nbr] = 1
		} else {
			m[nbr]++
		}
	}
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func main() {
	//mapSlice()
	words := "hello python, python go"
	countWords(words)
}
