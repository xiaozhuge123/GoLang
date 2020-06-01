package main
import "fmt"

func main(){
	//+号的运用，当两边的变量都是数值类型时，做加法计算；当两边的类型是字符串时，做拼接
	var i = 12
	var j = 25
	var k = i + j
	fmt.Println("k=",k)

	var str1 = "hello "
	var str2 = "go!"
	var str = str1 + str2
	fmt.Println("str="+str)
}