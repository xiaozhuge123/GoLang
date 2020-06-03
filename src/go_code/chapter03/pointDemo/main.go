package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	_"strconv"
)
//指针的基本介绍
func main(){
	var i int = 10
	fmt.Printf("i 的地址为%v \n",&i)

	//指针ptr
	var ptr *int = &i
	fmt.Printf("ptr 的地址为%v \n",ptr)

	//修改了i的值
	*ptr = 12
	//获取指针指向的值
	fmt.Printf("ptr指向的值%v \n",i)

	var _int = 1234
	fmt.Println("_int = ",_int)
}
