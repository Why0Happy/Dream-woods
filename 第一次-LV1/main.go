package main
import"fmt"

func convert(a string){
	for i:=len(a)-1;i>=0;i-- {
		fmt.Printf("%c",a[i])
	}
}
func main(){
	var str string
	fmt.Printf("请输入一个字符串：\n")
	fmt.Scanln(&str)
	fmt.Printf("字符串倒置结果为：\n")
	convert(str)
}