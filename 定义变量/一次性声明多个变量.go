/*
@Time : 2020/3/7
@Author : Doctor_Who丶Max
@File : 一次性声明多个变量
@Software: GoLand
*/

package main

import "fmt"

func main() {
	//第一种方式：
	var n1, n2, n3 int
	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)
	//第二种方式:
	var n4, name, n5 = 100, "tom", 888
	fmt.Println("n4=", n4, "name=", name, "n5=", n5)
	//第三种:类型推导方式
	n6, name, n7 := 100, "tom233", 888
	fmt.Println("n6=", n6, "name=", name, "n7=", n7)
}
