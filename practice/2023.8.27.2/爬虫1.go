package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var s, e int
	fmt.Print("爬取起始页：")
	fmt.Scan(&s)
	fmt.Print("爬取终止页：")
	fmt.Scan(&e)

	working(s, e)
}

// 爬取页面
func working(s, e int) {
	fmt.Printf("正在读取%d到%d页\n", s, e)
	for i := s; i <= e; i++ {
		url := "https://tieba.baidu.com/f?kw=%E7%A6%8F%E5%B7%9E%E5%A4%A7%E5%AD%A6&ie=utf-8&pn=" + strconv.Itoa((i-1)*50)
		result, err := HttpGet(url)
		if err != nil {
			fmt.Println("Http err:", err)
			continue
		}
		//将读到的数据保存为一个文件
		f, err := os.Create("第" + strconv.Itoa(i) + "页" + ".html")
		if err != nil {
			fmt.Println("creat err:", err)
			continue
		}
		f.WriteString(result)
		f.Close() //保存好一个网站，就关闭一个文件，不用defer
	}
}

func HttpGet(url string) (result string, err error) {
	resp, err1 := http.Get(url)
	if err1 != nil {
		err = err1
		return
	}
	defer resp.Body.Close()

	//读取网络数据并传出给调用者
	buf := make([]byte, 4096)
	for {
		n, err2 := resp.Body.Read(buf)
		if n == 0 {
			fmt.Println("读取网络完成")
			break
		}
		if err2 != nil && err2 != io.EOF {
			err = err2
			return
		}
		//累加buf，存入result，一次性返回
		result += string(buf[:n])
	}
	return
}
