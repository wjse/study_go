package main

import "fmt"

func main(){
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片
	fmt.Println(s1 ,s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	s1 = []int{1 , 2 ,3}
	s2 = []string{"北京" , "成都" , "广州"}
	fmt.Println(s1 , s2)
	fmt.Println(len(s1) , len(s2))

	//长度和容量
	fmt.Printf("len(s1) : %d cap(s1) : %d \n" , len(s1) , cap(s1))
	fmt.Printf("len(s2) : %d cap(s2) : %d \n" , len(s2), cap(s2))

	//由数组得到切片
	a1 := [...]int{1 , 3 , 5 , 7 , 9 , 11 , 13}
	s3 := a1[0 : 4] //[1 3 5 7] 基于一个数组切割，左包含右不包含，左闭右开
	fmt.Println(s3)

	s4 := a1[ : 4] //=> [0 : 4] 从开始切割到指定位置
	fmt.Println(s4)

	s5 := a1[3 : ] //=> [3 : len(a1)] 从指定位置切割到最后
	fmt.Println(s5)

	s6 := a1[:]//=> [0 : len(a1)] 从开始切割到最后
	fmt.Println(s6)

	//切片再切片
	//首先看由数组切割而来的切片的长度与容量
	//长度为4 ， 容量为7
	fmt.Printf("len(s4) : %d cap(s4) : %d \n" , len(s4), cap(s4))

	//长度为4 ， 容量为4
	//底层数组从切片的第一个元素到最后的元素数量
	fmt.Printf("len(s5) : %d cap(s5) : %d \n" , len(s5), cap(s5))

	//再切割
	s7 := s5[3 : ] //[7 9 11 13]
	fmt.Printf("len(s7) : %d cap(s7) : %d \n" , len(s7), cap(s7))

	//引用体现，底层数组值改了，切片的值都会更改
	a1[6] = 1300
	fmt.Println(s5 , s6 ,s7)//[7 9 11 1300] [1 3 5 7 9 11 1300] [1300]
}
