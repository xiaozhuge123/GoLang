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
	fmt.Print(flag)
}