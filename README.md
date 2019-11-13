# study_go

# 开始

## Go安装



下载 ：从 [Go 官网](https://golang.google.cn/)下载对应系统的Go

安装,选择不太复杂的路劲保存Go

安装完成后通过命令行执行

```shell
			go version
```

验证查看Go版本信息

配置GOPATH环境变量:

*GOPATH* 是一个环境变量，指向本机存放Go项目的目录（工作目录）

*GOPATH* 路劲最好只设置一个，所有的项目代码都放到*GOPATH*的src目录下

Linux环境编辑

```shell
sudo vi ~./bash_profile

​				export GOROOT='your go home'

​				export GOPATH='your dir'

​				export PATH=$PATH$:$GOROOT$/bin
```

保存后source ~./bash_profile使其生效

在GOPATH目录新建*src , pkg , bin*目录分别用来存放*源代码，包，执行文件*

执行

```shell
go env
```

将会看到如下系统环境信息

![image-20191113104837269](/Users/k66/Desktop/go/src/github.com/wjse/study_go/README.assets/image-20191113104837269.png)



## go install 交叉编译

windows => linux

```shell
SET CGO_ENABLE=0  //禁用CGO

SET GOOS=linux  //目标平台linux

SET GOARCH=amd64 //目标处理器架构为amd64
```



mac => linux or windows

```shell
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build

CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build
```



linux => mac or windows

```shell
CGO_ENABLE=0 GOOS=darwin GOARCH=amd64 go build

CGO_ENABLE=0 GOOS=windows GOARCH=amd64 go build
```



# Go语言文件的基本结构

```go
package main //包声明，如果为main，则生成可执行文件

//导入的包
import "fmt"

//main 函数，程序入口，如同c main函数或者java main
func main(){
   fmt.Println("Hello Go!")
}
```



## 变量和常量

### 标识符和关键字



#### 标识符

字母，数字和下划线"_"



####25个关键字



#### 37个保留字



### 变量

Go语言中必须先声明再使用变量

局部变量声明必须使用



```go
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
```



### 常量



```go
const Pi = 3.14
const Pi2 = 2 * Pi
```

