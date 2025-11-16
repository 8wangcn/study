package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("仓库初始化成功！")
	//修改后保存 会显示哪个文件进行了修改
	fmt.Println("修改文件")
	//然后选择文件 提交到暂存区 git add .
	fmt.Println("提交文件")
	//提交到本地仓库
	fmt.Println("提交仓库")
	//最后推送到远程仓库 git push
	fmt.Println("推送仓库")
	fmt.Println("推送成功")
	fmt.Println("test push")
}
