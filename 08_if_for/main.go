package main

import "fmt"

func main(){

	//if
	//基本
	age := 18
	if age > 18{
		fmt.Println("老男孩")
	}else if age == 18{
		fmt.Println("刚好成年")
	}else{
		fmt.Println("小屁孩")
	}

	//局部
	if age := 20 ; age < 35 && age > 18{
		fmt.Println("还年轻")
	}else if age < 18{
		fmt.Println("small boy")
	}

	//for
	//基本
	for i := 0 ; i < 10 ; i++{
		fmt.Println(i)
	}

	//省略变量
	i := 5
	for ; i < 10 ; i++{
		fmt.Println(i)
	}

	//省略结束
	for ; i < 10 ;{
		fmt.Println(i)
		i++
	}

	//无限循环
	//for {
	//	fmt.Println(1)
	//}

	//for range
	s := "hello沙河"
	for i , v := range s{
		fmt.Printf("%d %c \n" , i , v)
	}
}
