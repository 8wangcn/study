package main

import "fmt"

const (
	ReadPermission    = 1 << iota // 0001
	WritePermission               // 0010
	ExecutePermission             // 0100
)

func main() {
	// 电商场景：计算折扣价格
	price := 299.9
	discount := 0.8
	finalPrice := price * discount

	// 权限校验场景
	isAdmin := true
	hasWritePermission := false
	canEdit := isAdmin && hasWritePermission

	// 年龄验证
	age := 18
	canVote := age >= 18

	fmt.Printf("原价: %.2f, 折扣价: %.2f\n", price, finalPrice)
	fmt.Printf("可以编辑吗? %v\n", canEdit)
	fmt.Printf("可以投票吗? %v\n", canVote)

	userPermission := ReadPermission | WritePermission // 0011

	// 检查是否有写权限
	if userPermission&WritePermission != 0 {
		fmt.Println("用户有写权限")
	}
}
