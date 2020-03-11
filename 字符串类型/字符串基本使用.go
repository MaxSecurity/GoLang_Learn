/*
@Time : 2020/3/8
@Author : Doctor_Who丶Max
@File : 字符串基本使用
@Software: GoLand
*/

package main

import "fmt"

//string类型基本使用
func main() {
	var address string = "北京 100 holle"
	fmt.Println(address)
	//字符串一旦赋值，字符串不能够修改：在Go中字符串无法修改
	var str = "holle"
	//str[0] = "A" //这里表示把holle的第一位修改成A这个是不行的（课外小知识：编程中位数是从0开始算的0123456）
	fmt.Println("str=", str)
	//在双引号里面换行
	str2 := "abc\nabc"
	fmt.Println(str2)
	//输出原始内容
	str3 := `
	package main
	
	import "fmt"
	
	//字符串拼接
	func main() {
		var str = "hello" + "world"
		str += "hhh" //如果helloworld 在想拼接时候可以这样写
		fmt.Println(str)
		//拼接太长怎么办:换行  换行时候要把+放在上面G
		var str1 = "hello" + "world" + "world" + "world" +
			"world" + "world" + "world" + "world" + "world" +
			"world" + "world"
		fmt.Println(str1)
	}`
	fmt.Println(str3)
}
