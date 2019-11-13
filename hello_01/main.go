package main //包声明，如果为main，则生成可执行文件

//导入的包
import "fmt"

//main 函数，程序入口，如同c main函数或者java main
func main(){
	fmt.Println("Hello Go!")
}

/**
函数外只能放标识符（变量\常量\函数\类型）的声明
不能放表达式如 a + 1 , fmt.Println("xxx")
 */