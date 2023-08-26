package main

import "fmt"

func main() {
	var a [10]int
	var height int
	var sum int
	_, err := fmt.Scan(&a[0], &a[1], &a[2], &a[3], &a[4], &a[5], &a[6], &a[7], &a[8], &a[9])
	_, err = fmt.Scan(&height)
	if err == nil {
		for i := 0; i < len(a); i++ {
			if a[i] <= height {
				sum++
			} else {
				m := height + 30
				if a[i] <= m {
					sum++
				}
			}
		}
		fmt.Println(sum)
	}

}
