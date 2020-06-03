package main
import (
	"fmt"
)

func main(){
	//重点讲解 / %
	fmt.Println(10/4)
	fmt.Println(10/4.0)
	fmt.Println(10.000/4.0)

	//% a%b = a - a/b*b
	fmt.Println(10%4)
	fmt.Println(10%-4)
	fmt.Println(-10%4)
	fmt.Println(-10%-4)

	//算术运算符只能独立使用，不能赋值
	var i = 12
	i++
	fmt.Println(i)
	i--
	fmt.Println(i)

}