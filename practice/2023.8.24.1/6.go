package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	write()
}

func write() {
	file, err := os.OpenFile("D:\\biancheng\\code\\golang\\practice\\2023.8.24.1\\ninenine.txt", os.O_WRONLY, 0) //只写文件
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			_, err := file.WriteString(strconv.Itoa(i) + "*" + strconv.Itoa(j) + "=" + strconv.Itoa(i*j) + "\t")
			if err != nil {
				log.Fatal(err)
			}
		}
		file.WriteString("\n")
	}
}
