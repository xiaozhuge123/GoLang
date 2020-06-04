package main
import "fmt"
func main(){
	//其他运算符
	var a = 12
	fmt.Println("a的地址为\n",&a)
	var ptr *int = &a
	fmt.Println("ptr指向的变量值为\n",*ptr)

	//求两个数中的最大值
	var num1 = 123
	var num2 = 234
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	fmt.Println("max = \n",max)

	var num3 = 1234

	//比较3个数的最大值，先求出两个数的最大值，然后比较第三个数的最大值
	if max > num3 {
		max = max
	} else {
		max = num3
	}
	fmt.Println("max = ",max)


}