package main

func main() {
	//流程控制语句
	//if-else判断
	//if - else_if - else判断
	if condition {
		// 代码块
	} else if condition2 {
		// 代码块
	} else {
		// 代码块
	}

	//循环语句
	//for循环（Go只有for循环，没有while）：
	//传统for循环 for i := 0; i < 10; i++ { // 代码块 }
	//i := 0; i < 10; i++
	for i := 0; i < 10; i++ {
		// 代码块
	}
	// 类似while的用法 for condition { // 代码块 }
	//表达式
	condition := true
	for condition {
		//代码块
	}
	// 无限循环 for { // 代码块 }
	//注意这里for后面没有表达式，所以永远为true，因此每次循环判断表达式都不会退出
	for {
		// 代码块
	}

	//switch 等同于另一种if
	switch 表达式 {
	case value1:
		代码块
	case value2:
		代码块
	default:
		代码块
	}

	//select 等同于另一种if
	select {
	case ch1 <- data:
		// 代码块
	case data := <-ch2:
		// 代码块
	default:
		// 代码块
	}

}
