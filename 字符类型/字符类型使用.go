/*
@Time : 2020/3/8
@Author : Doctor_Who丶Max
@File : 字符类型使用
@Software: GoLand
*/

package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	//当直接输出byte值就是直接输出的对应的ASCLL码的值
	fmt.Println("c1=", c1)
	fmt.Println("c2=", c2)
	//如果希望输出对应的字符，需要使用格式化输出。
	fmt.Printf("c1=%c c2=%c", c1, c2)
	//字符型可以进行运算，相当于一个整数运行输出时按照ascll码值运行的
	var n1 = 10 + 'a'
	fmt.Println("\nn1 =", n1)
}
