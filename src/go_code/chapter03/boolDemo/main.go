package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main(){
	//bool类型
	var b = false
	fmt.Println("b=",b)
	fmt.Println("b占的字节数为：",unsafe.Sizeof(b))
}
