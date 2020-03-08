/*
@Time : 2020/1/30
@Author : Max
@File : 变量数据类型：整数数据类型
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//整数类型使用
	var i int = 10
	fmt.Println(i)
	//测试int8使用范围-128-127
	var a int8 = -128
	var b int8 = 127
	fmt.Println(a, b)

	//int uint rune ,byte 使用
	var c int = 8000
	fmt.Println("INT=", c)
	var d uint = 1
	fmt.Println("uint=", d)
	var e byte = 255
	fmt.Println("byte=", e)

	//查看某个变量数据类型
	//fmt.Printf 可以用于格式化输出
	var n1 = 100
	fmt.Printf("n1的数据类型=%T", n1)
}
