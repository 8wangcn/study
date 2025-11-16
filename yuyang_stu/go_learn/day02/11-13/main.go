package main

import "fmt"

type Address struct {
	Street, City, Country string
	ZipCode               string
}

type User struct {
	ID       int
	username string
	password int
	email    string
	age      int
	address  Address
}

func main() {
	addr := Address{
		Street:  "123 Street",
		City:    "HSD",
		Country: "USA",
		ZipCode: "1234",
	}

	user := User{
		ID:       1,
		username: "jack",
		password: 123,
		email:    "qqwer@123.com",
		age:      25,
		address:  addr,
	}

	fmt.Printf("用户id：%d, 用户名：%s, 邮箱：%s，年龄：%d, 国家：%v，城市：%v", user.ID,
		user.username, user.email, user.age, user.address.Country, user.address.City)

}
