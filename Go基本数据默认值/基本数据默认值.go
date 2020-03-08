/*
@Time : 2020/3/8
@Author : Doctor_Who丶Max
@File : 基本数据默认值
@Software: GoLand
*/

package main

import "fmt"

//在GO中，数据类型中都有一个默认值，当数据类型没有赋值时，就会保留默认值，在Go中默认值又叫零值
func main() {
	var a int     //默认值为:0
	var b float32 //默认值为:0
	var c float64 //默认值为:0
	var d bool    //默认值为:false
	var e string  //默认值为: ""
	fmt.Println("a=", a, "b=", b, "c=", c, "d=", d, "e=", e)
}
