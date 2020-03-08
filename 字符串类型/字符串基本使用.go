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
}
