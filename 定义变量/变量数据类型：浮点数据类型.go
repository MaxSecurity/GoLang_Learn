/*
@Time : 2020/1/31
@Author : Max
@File : 变量数据类型：浮点数据类型
@Software: GoLand
*/
package main

import "fmt"

func main() {
	//小数类型就是存放小数的，比如1.1,2.6等等。
	var a float32 = 89.12
	fmt.Println(a)
	fmt.Printf("A的数据类型是: %T", a) //fmt.Printf 用于格式化输出案例如前
}
