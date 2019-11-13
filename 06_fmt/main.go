package main

import "fmt"

func main(){
	//占位符,更多查看fmt包下doc.go
	n := 100
	fmt.Printf("%T \n" , n) //类型
	fmt.Printf("%v \n" , n) //值
	fmt.Printf("%b \n" , n) //二进制
	fmt.Printf("%d \n" , n) //十进制
	fmt.Printf("%o \n" , n) //八进制
	fmt.Printf("%x \n" , n) //十六进制

	s := "abc"
	fmt.Printf("%s \n" ,s) //字符
	fmt.Printf("%#v \n" ,s) //加#添加描述符比如字符串+""
	fmt.Printf("%T \n" , '沙')// 字符类型为int32
}
