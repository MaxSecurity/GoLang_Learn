/*
@Time : 2020/3/7
@Author : Doctor_Who丶Max
@File : 变量的使用
@Software: GoLand
*/

package main

import "fmt"

func main() {
	//GO变量使用方法1,指定变量类型，声明不赋值，使用默认值
	var i int
	fmt.Println("i=", i)
	//第二种使用方式：根据值自行判定变量类型（类型推导）
	var test = 123.11
	fmt.Println("test=", test)
	//第三种使用方式：省下var，:=左侧变量不应该是声明过的，否则就会错误
	name := "tom"
	fmt.Println("name=", name)
}
