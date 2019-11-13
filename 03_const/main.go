package main

import "fmt"

const pi = 3.1415926

const(
	ok = 200
	notFound = 404
	serverError = 500
)

const(
	n1 = 100
	n2
	n3
)

const(
	a1 = iota * 10 // 0 * 10
	a2             // 1 * 10
	a3             // 2 * 10
	a4             // 3 * 10
)

//常见的iota运用
//_使用
const(
	b1 = iota // 0 当const出现时重置
	b2        // 1
	_         // 2 _匿名变量，被丢弃
	b3        // 3
)

//插队情况
const(
	c1 = iota // 0
	c2 = 100  // 100
	c3        // 100
	c4        // 100
	c5 = iota // 4 , 常量没增加一个声明，iota计数+1
)

//多个常量声明在一行
const(
	d1 , d2 = iota + 1 , iota + 2  // 1 , 2
	d3 , d4 = iota + 1 , iota + 2  // 2 , 3  每新增一行常量声明iota计数才+1
)

//定义数量级
const(
	_ = iota
	kb = 1 << (10 * iota)
	mb = 1 << (10 * iota)
	gb = 1 << (10 * iota)
	tb = 1 << (10 * iota)
	pb = 1 << (10 * iota)
)

func main(){
	fmt.Println(n1 , n2 , n3)

	fmt.Println(a1 , a2 , a3 , a4)

	fmt.Println(b1, b2 ,b3)

	fmt.Println(c1, c3 ,c4 , c5)

	fmt.Println(d1 , d2 , d3 , d4)

	fmt.Println(kb , mb , gb , tb , pb)
}