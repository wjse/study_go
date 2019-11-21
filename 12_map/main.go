package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main(){
	var m1 map[string]int
	fmt.Println(m1)//还未初始化,没有在内存中开辟空间

	m1 = make(map[string]int , 10)//初始化时估算好容量，避免在程序运行期间再动态扩容
	fmt.Println(m1)

	m1["老吴"] = 18
	m1["老张"] = 19
	fmt.Println(m1)

	//约定成俗用ok接收返回的布尔值
	value , ok := m1["老李"]
	if !ok {
		fmt.Println("查无此key")
	}else{
		fmt.Println(value)
	}

	//如果不存在这个key拿到对应值类型的零值
	fmt.Println(m1["老杨"])

	//遍历 , 只想遍历key就只写一个变量接收 , 只写遍历value就需要用匿名变量_接收
	for k , v := range m1 {
		fmt.Println(k , v)
	}

	//遍历key
	for k := range m1{
		fmt.Println(k)
	}

	//遍历value
	for _ , v := range m1{
		fmt.Println(v)
	}

	//删除
	delete(m1, "老吴")
	fmt.Println(m1)

	//删除不存在的key
	delete(m1 , "老李")

	//按顺序遍历map
	getMapValueBySort()

	//元素为map的切片
	mapOfSlice()

	//值为切片的map
	sliceOfMap()
}

func getMapValueBySort(){
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int , 200)//初始化map

	//map放入元素
	for i := 0 ; i < 100 ; i++{
		key := fmt.Sprintf("stu%02d" , i) //生成stu开头的字符串
		value := rand.Intn(100) //生成0-99的随机整数
		scoreMap[key] = value
	}

	var keys = make([]string , 0 , 200) //初始化用来存放key的切片

	//将map的key放入keys切片中
	for key := range scoreMap{
		keys = append(keys , key)
	}

	//对keys进行排序
	sort.Strings(keys)

	//遍历排好序的keys切片，遍历切片取出map中对应的value
	for _ , key := range keys{
		fmt.Println(key , scoreMap[key])
	}
}

func mapOfSlice(){

	//切片在初始化时如果长度为0，则放入元素时会越界
	var s1 = make([]map[int]string , 10 , 10)

	//直接赋值时map未初始化
	//s1[0][1] = "A"

	s1[0] = make(map[int]string , 10)
	s1[0][1] = "A"
	fmt.Println(s1)
}

func sliceOfMap(){
	var m1 = make(map[int][]string , 10)
	var s1 = make([]string , 1 , 1)
	s1[0] = "A"
	m1[0] = s1
	fmt.Println(m1)

	m1[1] = []string{"B"}
	fmt.Println(m1)
}
