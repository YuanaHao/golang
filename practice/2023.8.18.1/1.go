package main

import "fmt"

func main() {
	/*a1 := make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a1 = append(a1, i)
	}
	fmt.Println(a1)*/
	var m map[string]int
	m = make(map[string]int, 10) //初始化map
	m["理想"] = 18
	m["李华"] = 35
	// fmt.Println(m)
	// value, ok := m["张三"] //查找map中是否有key，约定用OK返回布尔值
	// if !ok {
	// 	fmt.Println("查无此人")
	// } else {
	// 	fmt.Println(value)
	// }
	for k, v := range m {
		fmt.Println(k, v)
	}
}
