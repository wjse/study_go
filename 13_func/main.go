package main

import "fmt"

func main(){
	sum := sum(5 , 8)
	fmt.Println(sum)

	m , n := f4(10 , "hh")
	fmt.Println(m , n)

	ret := f5(100 , 2)
	fmt.Println(ret)

	xyz := f6(3 , 3 , 3 , "a" , "b" , true , false)
	fmt.Println(xyz)

	f7("1")
	f7("1" , 1)
}

//函数定义
func sum(a int , b int)(ret int){
	return a + b
}

//多参多返回
func f4(a int , b string)(int , string){
	c := a + 5
	d := b + "abc"
	return c , d
}

//参数可以命名可以不命名
//命名的返回值就相当于在函数中声明了一个变量
func f5(x int , y int)(ret int){
	ret = x + y
	return
}

//参数的类型简写 , 当参数的类型一致时，可以将前面的参数类型省略
func f6(x , y , z int , m , n string , i ,j bool) int {
	fmt.Println(m , n , i , j)
	return x * y / z
}

//可变长参数，必须放在参数列表最后
//y的类型是切片
func f7(x string , y ...int){
	fmt.Println(x , y)
}