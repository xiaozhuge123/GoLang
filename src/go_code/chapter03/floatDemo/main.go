package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
)

func main(){
	//浮点型数 float32  float64
	var a float32 = 12.34567
	var b float64 = 375.233567
	fmt.Println("a=",a,"b=",b)

	//精度损失
	var c float32 = -12.000000234
	var d float64 = -12.000000234
	fmt.Println("c=",c,"d=",d)

	//默认float64位
	var num5 = 123
	fmt.Printf("num5 类型  %T \n",num5)

	//类型推导
	num6 := 2.12
	num7 := .123
	fmt.Println("num6=",num6,"num7=",num7)

	//科学计数法
	num8 := 2.134e2
	num9 := 2.134E2
	num10 := 2.134e-2
	fmt.Println("num8=",num8,"num9=",num9,"num10=",num10)
}
