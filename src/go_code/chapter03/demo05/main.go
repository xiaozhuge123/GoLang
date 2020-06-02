package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main(){
	//int8  有符号的整型，注意范围 -128~127
	var i int8
	i = 127
	fmt.Println("i=",i)
	//uint8  无符号的整型，注意范围 0~255
	var j uint8 = 255
	fmt.Println("j=",j)
	//int,uint,byte,rune
	var a int = 12345
	var b uint = 0
	var c byte = 255
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	fmt.Println("c = ",c)

	//显示定义的变量的数据类型
	var n1 = 8
	fmt.Printf("n1 默认的数据类型 %T \n",n1)
	//显示变量的数据类型以及占用的字节数的大小
	var n2 int64 = 1234
	fmt.Printf("n2 数据类型 %T  字节数 %d",n2,unsafe.Sizeof(n2))
}