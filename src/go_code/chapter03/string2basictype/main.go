package main
// import "fmt"
// import "unsafe"

import (
	"fmt"
	"strconv"
)
//String转基本数据类型
func main(){
	//String转bool
	var str1 = "true"
	var b1 bool
	b1 , _ = strconv.ParseBool(str1)
	fmt.Printf("b1 type is %T b1 = %v \n",b1,b1)

	//String转int64  如果需要转成int  则需要用int强转
	var str2 = "123456"
	var num1 int64
	num1 , _ = strconv.ParseInt(str2,10,64)
	fmt.Printf("num1 type is %T  num1 = %v \n",num1,num1)
	var num2 int = int(num1)
	fmt.Printf("num2 type is %T  num2 = %v \n",num2,num2)

	//string 转float64
	var str3 = "12.34567"
	var flo float64
	var flo1 float32
	flo , _ = strconv.ParseFloat(str3,64)
	fmt.Printf("flo type is %T ,flo = %v \n",flo,flo)
	flo1 = float32(flo)
	fmt.Printf("flo1 type is %T,flo1 = %v \n",flo1,flo1)

	var str4 = "hello"
	num1 , _ = strconv.ParseInt(str4,10,64)
	fmt.Printf("num1 type is %T  num1 = %v \n",num1,num1)

	b1 , _ = strconv.ParseBool(str4)
	fmt.Printf("b1 type is %T b1 = %v \n",b1,b1)

	flo , _ = strconv.ParseFloat(str4,64)
	fmt.Printf("flo type is %T ,flo = %v \n",flo,flo)

}
