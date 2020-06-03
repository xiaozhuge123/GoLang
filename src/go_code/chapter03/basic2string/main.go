package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"strconv"
)
//基本数据类型转String
func main(){
	//第一种方式 fmt.Sprintf()
	var num1 int = 12
	var num2 float64 = 12.34567
	var b1 bool = true
	var c1 byte = 'a'
	var str string

	//int转String
	str = fmt.Sprintf("%d",num1)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	//float转String
	str = fmt.Sprintf("%f",num2)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	//布尔类型转成String
	str = fmt.Sprintf("%t",b1)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	//字符转字符串
	str = fmt.Sprintf("%c",c1)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	//基本数据类型转成字符串的第二种方式 strconv函数
	var num3 int = 1234
	var num4 float64 = 123.456789
	var b2 bool = true
	str = strconv.FormatInt(int64(num3),10)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	str = strconv.FormatFloat(num4,'f',10,64)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type is %T  str = %q \n",str,str)

	str = strconv.Itoa(num3)
	fmt.Printf("str type is %T  str = %q \n",str,str)


}
