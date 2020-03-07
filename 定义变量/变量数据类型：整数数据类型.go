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
}
