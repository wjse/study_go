package main

import "fmt"

func main(){

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
