package main

import (
	"fmt"
	"math"
	"strings"
	"unicode/utf8"
)

func main(){

	s1 := "adadad\rasdada"
	fmt.Println(s1)
	fmt.Println(len(s1))

	//多行字符串 ,``原样输出
	s2 := `
		第一行
		第二行
		第三行
          `
	fmt.Println(s2)

	s3 := `d:\file`
	fmt.Println(s3)

	//字符串相关操作
	//length
	fmt.Println(len(s3))

	//拼接
	name := "深田"
	word := "永美"
	fmt.Println(name + word)
	fmt.Printf("%s%s\n" , name , word)
	s4 := fmt.Sprintf("%s%s" , name , word)
	fmt.Println(s4)

	//分隔
	s5 := strings.Split(s3 , "\\")
	fmt.Println(s5)

	//包含
	fmt.Println(strings.Contains(s4 , "深田"))

	//前后缀判断
	fmt.Println(strings.HasPrefix(s4 , "永美"))
	fmt.Println(strings.HasSuffix(s4 , "永美"))

	//index
	fmt.Println(strings.Index(s4 , "田"))

	//join
	fmt.Println(strings.Join(s5 , "\\"))

	//非ASCII类型为rune类型
	str := "hello成都"
	//len()求的是byte字节数量
	n := len(str)
	fmt.Println(n)
	fmt.Println(utf8.RuneCountInString(str))

	//字符int值
	for i := 0 ; i < len(str) ; i++ {//byte
		fmt.Printf("%v(%c) \n" , str[i] , str[i])
	}

	//从字符传中拿出具体的字符
	for _, c := range str{//rune
		fmt.Printf("%c \n" , c)
	}

	//字符串修改，字符串是不能被修改的，只能重造
	changeString()

	//类型转换
	num := 20
	var f float64
	f = float64(num)
	fmt.Printf("%T %v \n" , f , f)
	sqrtDemo()
}

func changeString(){
	s1 := "big"
	byteS1 := []byte(s1)//强制转换为byte切片
	byteS1[0] = 'p'//设置第一个元素为新字符
	fmt.Println(string(byteS1))//将byte切片强制转换为字符串，new string

	s2 := "白萝卜"
	runeS2 := []rune(s2)//强制转换为rune切片
	runeS2[0] = '胡'//设置第一个元素为新字符
	fmt.Println(string(runeS2))//将rune切片强制转换为字符串，new string

}

func sqrtDemo(){
	var a, b = 3 , 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
