package main
import "fmt"
//一次性声明多个全局变量
var n4, n5, n6 = 10,"tom",30.4

var (
	n7,n8,n9 = 12,"jerry",45.867
)

func main(){
	//该案例演示如何一次性声明多个变量
	// var n1, n2, n3 int
	// fmt.Println("n1=",n1,"n2=",n2,"n3=",n3)

	//一次性声明多个不同类型的变量
	// var n4, n5, n6 = 10,"tom",30.4
	// fmt.Println("n4=",n4,"n5=",n5,"n6=",n6)
	fmt.Println("n4=",n4,"n5=",n5,"n6=",n6)
	//一次性声明多个不同类型的变量，省略var的方式
	// n7,n8,n9 := 12,"jerry",45.867
	// fmt.Println("n7=",n7,"n8=",n8,"n9=",n9) 
	fmt.Println("n7=",n7,"n8=",n8,"n9=",n9)
}