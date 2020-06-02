package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
)

func main(){
	//字符串使用
	var str string = "中国 hello world"
	fmt.Println(str)
	//字符串一旦赋值了就不能修改
	Str2 := "hello"
	// Str2[0]='a'
	fmt.Println(Str2)

	//双引号  和 反引号
	str3 := "hello\nworld"
	fmt.Println(str3)
	str4 := `hello
	dnfdgjdfgj
	jsdfjgfdhghdf`
	fmt.Println(str4)

	//字符串拼接 "+"号必须在上一行的末尾
	str5 := "hello" +
	"world"
	fmt.Println(str5)

	//基本数据类型的默认值
	var a int
	var b float32
	var c float64
	var d byte
	var e string
	var f bool
	fmt.Println(a," ",b," ",c," ",d," ",e," ",f)

	//基本数据类型的相互转化，需要显示的转换  不能自动转换
	var i int32 = 123
	var j float32 = float32(i)
	var k int8 = int8(j)
	fmt.Printf("i=%v j=%v k=%v",i,j,k)

	//如果大转小，会按溢出处理
	var as int64 = 1245
	var sd int8 = int8(as)
	fmt.Printf("sd=%v",sd)

}
