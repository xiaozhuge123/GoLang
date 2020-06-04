package main
import "fmt"
func main(){
	//关系运算符
	var n1 = 9
	n2 := 10
	fmt.Print(n1 == n2, "\n")
	fmt.Print(n1 != n2 ,"\n")
	fmt.Print(n1 > n2, "\n")
	fmt.Print(n1 < n2 ,"\n")
	fmt.Print(n1 >= n2 ,"\n")
	fmt.Print(n1 <= n2 ,"\n")

	flag := n1 > n2
	fmt.Print(flag,"\n")

	//逻辑运算符 &&逻辑与  ||逻辑或  ！逻辑非

	flag2 := n2 > n1
	fmt.Print(flag && flag2,"\n")
	fmt.Print(flag || flag2,"\n")
	fmt.Print(!flag2,"\n")

}