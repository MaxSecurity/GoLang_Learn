/*
@Time : 2020/1/28
@Author : Max
@File : 全局变量声明
@Software: GoLand
*/
package main

import "fmt"

//定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

//上面的声明方式，也可以改成一次性声明
var (
	n3    = 300
	n4    = 900
	name2 = "mary"
)

func main() {

	//输出全局变量
	fmt.Println("n1=", n1, "name=", name, "n2=", n2)
	fmt.Println("n3=", n3, "name2=", name2, "n4=", n4)

}
