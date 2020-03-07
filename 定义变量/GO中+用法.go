/*
@Time : 2020/1/28
@Author : Max
@File : GO中+用法
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//当两边是数值时加作为加法运算
	var a = 1
	var b = 2
	var c = a + b
	fmt.Println(c)
	//当两边是字符串的时候加号表示拼接
	var name1 = "hello"
	var name2 = "world"
	var result = name1 + name2
	fmt.Println(result)
}
