package main

import "fmt"

/*
数组
存放元素的容器
必须指定存放的元素类型和长度
*/
func main() {
	//数组定义
	var a1 [3]int
	var a2 [4]int
	fmt.Printf("%T %T \n", a1, a2)

	//数组初始化
	//如果不初始化，默认均为数据类型默认值
	fmt.Println(a1, a2)

	a1 = [3]int{1, 2, 3}
	fmt.Println(a1)

	//[...]根据初始值自动推断数组长度
	a10 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a10)
	fmt.Println(len(a10))

	a3 := [5]int{1, 2}
	fmt.Println(a3)

	//根据索引来初始化
	a4 := [5]int{0: 4, 4: 3}
	fmt.Println(a4)

	//数组遍历
	cities := [3]string{"北京", "成都", "广州"}
	//根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	//for range
	for i, v := range cities {
		fmt.Println(i, v)
	}

	//多维数组
	//[[1,2] , [3,4] , [5,6]]
	aa1 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(aa1)

	//多维数组遍历
	for _, v := range aa1 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}
	}

	//数组是值类型
	b1 := [3]int{1, 2, 3} //[1 2 3]
	b2 := b1              //[1 2 3] ctrl+c => ctrl+v
	b2[0] = 100           //b2:[100 2 3]
	fmt.Println(b1, b2)   //b1:?

	//数组值类型比较
	c1 := [3]int{1 ,2 ,3}
	c2 := [3]int{1 ,2 ,3}
	fmt.Println(c1 == c2)
}
