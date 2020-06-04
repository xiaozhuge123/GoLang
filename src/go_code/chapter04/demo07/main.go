package main
import "fmt"
func main(){
	//获取键盘输入的数据 fmt.Scanln()  fmt.Scanf()
	var name string
	var age int
	var sal float32
	var isPass bool
	// fmt.Println("请依次按提示输入信息：\n")
	// fmt.Println("请输入姓名：\n")
	// fmt.Scanln(&name)

	// fmt.Println("请输入年龄：\n")
	// fmt.Scanln(&age)

	// fmt.Println("请输入薪水：\n")
	// fmt.Scanln(&sal)

	// fmt.Println("请输入是否通过考试：\n")
	// fmt.Scanln(&isPass)

	// fmt.Printf("姓名=%v \n 年龄=%v \n 薪水=%v \n 是否通过=%v \n",name,age,sal,isPass)

	//第二种方式
	fmt.Println("请按照顺序输入：")
	fmt.Scanf("%s %d %f %t",&name,&age,&sal,&isPass)
	fmt.Printf("姓名=%v \n 年龄=%v \n 薪水=%v \n 是否通过=%v \n",name,age,sal,isPass)

}