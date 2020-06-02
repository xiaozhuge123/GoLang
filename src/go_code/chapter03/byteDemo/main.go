package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
)

func main(){
	//字符数据
	var c1 byte = 'a'
	var d byte = '0'
	fmt.Println("c1:",c1,"d:",d)

	//格式化输出
	//fmt.Printf("c 类型 %T ",c,"d类型 %T ",d)
	fmt.Printf("c1=%c d=%c\n",c1,d)

	var c3 int = '被'
	fmt.Printf("c3=%c c3=%d",c3,c3)

}
