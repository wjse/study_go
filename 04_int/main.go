package main

import "fmt"

func main(){
	//十进制
	var a1 = 101
	fmt.Printf("%d \n" , a1)

	//十进制 -> 八进制
	fmt.Printf("%o \n" , a1)

	//十进制 -> 十六进制
	fmt.Printf("%x \n" , a1)

	//十进制 -> 二进制
	fmt.Printf("%b \n" , a1)

	//八进制
	a2 := 077
	fmt.Printf("%d \n" , a2)

	//十六进制
	a3 := 0x1234567
	fmt.Printf("%d \n" , a3)

	//查看变量类型
	fmt.Printf("%T \n" , a3)

	a4 := int8(9) //明确指定类型，否则为int类型
	fmt.Printf("%T \n" , a4)
}
