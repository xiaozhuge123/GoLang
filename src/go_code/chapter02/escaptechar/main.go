//转义字符 \t \n \\ \" \r
package main
import "fmt"
func main(){
	fmt.Println("tom\tjack")
	fmt.Println("tom\njack")
	fmt.Println("hello \"i love you\" shao")
	fmt.Println("d:\\user\\mysql")
	fmt.Println("hello\rwo")
	//用一句语句输出
	fmt.Println("姓名\t年龄\t籍贯\t住址\ntom\t12\t丑国\t白宫")
}