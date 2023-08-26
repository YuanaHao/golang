package main

import "fmt"

func main() {
	//数组元素求和

	/*a := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range a {
		sum = sum + v
	}
	fmt.Println(sum)
	*/

	/*a1 := [...]int{1, 3, 5, 7, 8}
	for i := 0; i < len(a1); i++ {
		for j := i + 1; j < len(a1); j++ {
			if a1[i]+a1[j] == 8 {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}*/

	//由数组得到切片
	a1 := [...]int{1, 2, 3, 4, 5, 6}
	b1 := a1[0:4]
	fmt.Println(b1)

}
