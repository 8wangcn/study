package main

import (
	"fmt"
)

func main() {
	//1 var
	//2 str :=
	var str string
	var num int
	var dou float64
	var flag bool
	var ch byte
	str = "wang"
	num = 1
	dou = 1.2
	flag = true
	ch = 'c'
	fmt.Print()
	fmt.Println()
	fmt.Printf("str:%s  num:%d dou:%0.2f flag:%v ch:%v \n", str, num, dou, flag, ch)
	//%T 输出变量的类型
	fmt.Printf("这个变量的类型是: %T", str)

	//常量
	const Name string = "WangKaiBin"

}
