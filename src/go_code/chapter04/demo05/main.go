package main
import "fmt"
func main(){
	//赋值运算符 = += *= /= -= %=
	a := 12
	b := 13
	t := a
	a = b
	b = t
	fmt.Println(a,"\n",b,"\n")
	var c int = 12
	c += a
	fmt.Println(a,"\n",c,"\n")

	//交换ab不允许使用中间变量
	a = a + b
	b = a - b //a + b - b  = b = a
	a = a - b //a + b - a = b
}