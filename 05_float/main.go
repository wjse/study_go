package main

import (
	"fmt"
)

func main(){
	f1 := 1.23456 //float64 , Go中默认都为float64
	fmt.Printf("%T \n" , f1)

	f2 := float32(1.23456)
	fmt.Printf("%T \n" , f2)

	//f1 = f2  //float32类型的值不能直接赋值给float64类型
}
