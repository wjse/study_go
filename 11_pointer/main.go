package main

import "fmt"

//指针
func main(){
	//&取地址
	n := 19
	fmt.Println(&n)


	p := &n
	fmt.Println(p)
	fmt.Printf("%T \n" , p) //*int int类型的指针

	//*根据内存地址取值
	m := *p
	fmt.Println(m)
	fmt.Printf("%T \n" , m) //int

	//nil pointer
	//var a *int
	//*a = 100

	//new函数申请一个内存地址
	var a = new(int)
	fmt.Println(a , *a)
	*a = 100
	fmt.Println(*a)
}
