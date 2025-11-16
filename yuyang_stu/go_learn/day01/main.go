package main

import "fmt"

func main() {
	/*var str string
	var num int
	var dou float64
	var flag bool
	var ch byte*/

	//运算符
	/*
		| 运算符 | 描述 | | :—-: | :–: | | + | 相加 | | - | 相减 | | * | 相乘 | | / | 相除 | | % | 求余 |

		注意： ++（自增）和--（自减）在Go语言中是 单独的语句，并不是运算符。

		关系运算符
		| 运算符 | 描述 | | :—-: | :———————————————————– | | == | 检查两个值是否相等，如果相等返回 True 否则返回 False。 | | != | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。 | | > | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。 | | >= | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | | < | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。 | | <= | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 |

		结果类型为 bool

		逻辑运算符
		运算符	描述
		&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。
		||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。
		!	逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
		具有短路特性，**&& 见假就停，	 	见真就停**

		00101100
		00101111
		00101100
	*/

	// +  1 + 1 = 2
	// -  1 - 1 = 0
	// *  1 * 1 = 1
	// /  1 / 1 = 1
	// %  10 % 5 = 0
	// ++ 和 --
	a := 10
	a++
	fmt.Printf("a:%v \n", a)
	b := 10
	b = b + b
	fmt.Printf("b:%v \n", b)
	c := 21
	d := c % 2
	fmt.Println(d)
	// ==  10 == 5 >> false
	fmt.Println(10 == 5)
	// !=   10 != 5 >> true
	fmt.Println(10 != 5)
	// >  10 > 5 >> true
	fmt.Println(10 > 5)
	// <  10 < 5 >> false
	fmt.Println(10 < 5)
	// >=  10 >= 10 >> true
	fmt.Println(10 >= 10)
	// <=  10 <= 10 true
	fmt.Println(10 <= 10)

	a1 := 2
	a2 := 4
	fmt.Println("====================")
	fmt.Println(a1 > a2 || a1 < a2)

	fmt.Println("====================")
	c1 := 11
	c1 += 10
	c1 = c1 + 10
	//-=
	// *=
	// /=
	// %=
}
