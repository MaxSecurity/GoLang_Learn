/*
@Time : 2020/3/8
@Author : Doctor_Who丶Max
@File : 字符串拼接方式
@Software: GoLand
*/

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
}
