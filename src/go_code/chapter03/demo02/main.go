package main
import "fmt"
func main(){
	//1、指定变量类型，声明后若不赋值，使用默认值
	var i int
	fmt.Println("i=",i)
	//2、根据值自动判断变量类型（类型推导）
	var j = 12
	fmt.Println("j=",j)
	//3、省略var，但是等号前必须使用:，写法:=,=左边的变量不应该是已经定义过的，否则编译会报错
	name := "tome"
	fmt.Println("name=",name)
}