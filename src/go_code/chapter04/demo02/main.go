package main
import "fmt"
func main(){
	//97天放假  还有多少星期  多少天
	var day = 97
	var week = day / 7
	var day1 = day % 7
	fmt.Printf("week = %v,day1 = %v \n",week,day1)

	//华氏温度转摄氏温度  5/9*(摄氏温度-100)
	var hua = 120.34
	var she = 5.0 / 9 * (hua - 100)
	fmt.Printf("she = %v ",she)
}