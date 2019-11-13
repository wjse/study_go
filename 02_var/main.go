package main

import "fmt"

//普通声明
//var name string
//var age int
//var isMan bool

//批量声明
var(
	name string // ""
	age int // 0
	isMan bool //false
)
//全局变量必须以关键字开头如var ，const ， type等

func main(){
	name = "Jim"
	age = 18
	isMan = true
	//Go语言中变量必须使用
	fmt.Println(name , age , isMan)

	//局部变量声明必须使用否则编译不通过
	//var tmp string

	//声明并赋值
	var tmp string = "tmp"
	fmt.Printf(tmp)

	//类型推导声明
	var str = "str"
	fmt.Println(str)

	//简短变量声明,只能在函数里面使用
	num := 5
	fmt.Println(num)

	//匿名变量，比如某些函数返回多个返回值，而只需要使用一个，则其他的可用匿名变量接收
	//匿名变量不占用命名空间，不会分配内存，也不存在重复声明
	var fooValue , _ = foo()
	fmt.Println(fooValue)
}

func foo() (int , string){
	return 5 , "Ha"
}