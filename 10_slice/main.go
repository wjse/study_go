package main

import (
	"fmt"
	"sort"
)

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

	//make
	sliceMake()

	//append
	sliceAppend()

	sliceCopy()

	sliceDel()

	train()
}

func sliceMake(){
	//s1 = [0 0 0 0 0] len(s1) = 5 cap(s1) = 10
	s1 := make([]int , 5 , 10)
	fmt.Printf("s1 = %v len(s1) = %d cap(s1) = %d\n" , s1 , len(s1) , cap(s1))

	//s0 = [0 0 0 0 0] len(s0) = 5 cap(s0) = 5
	s0 := make([]int , 5)
	fmt.Printf("s0 = %v len(s0) = %d cap(s0) = %d\n" , s0 , len(s0) , cap(s0))

	//s2 = [] len(s2) = 0 cap(s2) = 10
	s2 := make([]int , 0 , 10)
	fmt.Printf("s2 = %v len(s2) = %d cap(s2) = %d\n" , s2 , len(s2) , cap(s2))

	var s3 []int //len(s3) = 0 , cap(s3) = 0 , s3 == nil //true
	fmt.Printf("s3 = %v len(s3) = %d cap(s3) = %d s3 == nil %v\n" , s3 , len(s3) , cap(s3) , s3 == nil)

	s4 := []int{} //len(s4) = 0 , cap(s4) = 0 , s4 != nil //true
	fmt.Printf("s4 = %v len(s4) = %d cap(s4) = %d s4 != nil %v\n" , s4 , len(s4) , cap(s4) , s4 != nil)

	s5 := make([]int , 0) //len(s5) = 0 , cap(s5) = 0 , s5 != nil //true
	fmt.Printf("s5 = %v len(s5) = %d cap(s5) = %d s5 != nil %v\n" , s5 , len(s5) , cap(s5) , s5 != nil)

	//切片赋值
	s6 := []int{1 ,3 ,5}
	s7 := s6 //s6和s7都指向同一个底层数组
	fmt.Println(s6 , s7)

	s6[0] = 1000
	fmt.Println(s6 , s7)

	//切片的遍历
	for i := 0 ; i < len(s6) ; i++{
		fmt.Println(s6[i])
	}

	for i , v := range s6{
		fmt.Println(i , v)
	}
}

func sliceAppend(){
	s1 := []string{"北京" , "上海" , "深圳"}
	fmt.Printf("len(s1) = %d cap(s1) = %d %v\n" , len(s1) , cap(s1) , s1)

	//append追加元素，原来底层数组放不下的时候,将扩容
	s1 = append(s1, "成都")
	fmt.Printf("len(s1) = %d cap(s1) = %d %v\n" , len(s1) , cap(s1) , s1)

	s2 := []string{"广州" , "杭州"}
	s1 = append(s1 , s2...)//...表示拆开
	fmt.Printf("len(s1) = %d cap(s1) = %d %v\n" , len(s1) , cap(s1) , s1)

	//简原来的的容量小于1024时直接扩充至原来的2倍 ， 大于1024时每次扩充1/4
	s1 = append(s1, "重庆")
	fmt.Printf("len(s1) = %d cap(s1) = %d %v\n" , len(s1) , cap(s1) , s1)
}

func sliceCopy() {
	a1 := []int{1 ,3 , 5}
	a2 := a1 //赋值 ， 指向同一个内存
	var a3 = make([]int , 3 , 3)
	copy(a3 , a2) //copy
	fmt.Println(a1 , a2 , a3)
	a1[0] = 100
	fmt.Println(a1 , a2 , a3)
}

func sliceDel(){
	a := []int{1 , 2 , 3 , 4 , 5}

	//要删除索引为2的元素，a[:2]从0开始到2但不包含2[1 2] ， a[3:]...从3开始到最后[4 5]
	a = append(a[:2] , a[3:]...)
	fmt.Println(a)
	fmt.Printf("cap(a) = %d \n" , cap(a))

	a1 := [...]int{1 ,3 ,5} //数组
	s1 := a1[:] //切片
	fmt.Println(s1 , len(s1) , cap(s1))
	fmt.Printf("%p \n" , &a1)
	//1.切片不保存具体的值
	//2.切片对应一个底层的数组
	//3.底层数组占用一块联系的数组
	s1 = append(s1[:1] , s1[2:]...)
	fmt.Println(s1 , len(s1) , cap(s1))
	fmt.Println(a1)
	fmt.Printf("%p \n" , &a1)
}

//练习
func train(){
	a := make([]int , 5 , 10)
	for i := 0 ; i < 10 ; i++{
		//a = append(a , fmt.Sprintf("%v" , i))
		a = append(a , i)
	}

	//[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
	fmt.Println(a)
	fmt.Printf("len = %d cap = %d \n" , len(a) , cap(a))

	//切片排序
	a1 := []int{4 ,2 ,1 ,6 ,7,5}
	fmt.Println(a1)
	sort.Ints(a1)
	fmt.Println(a1)
}