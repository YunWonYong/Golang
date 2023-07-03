package main

import "fmt"

type (
	User struct {
		Name string
		ID   string
		Age  int
	}

	VIPUser struct {
		UserInfo User
		VIPLevel int
		Price    int
	}
)

type (
	VIPUser2 struct {
		User
		VIPLevel int
		Price    int
	}
)

func main() {
	code1()
	code2()
}

func code1() {
	user := User{
		"wonyong.yun",
		"ywyi1992",
		32,
	}

	VIPUser := VIPUser{
		UserInfo: user,
		VIPLevel: 10,
		Price:    50000000,
	}

	userPrint(user)
	vipPrint(VIPUser)
}

func code2() {
	VIPUser := VIPUser2{
		User:     User{"wonyong.yun", "ywyi1992", 32},
		VIPLevel: 100,
		Price:    500000,
	}

	vip2Print(VIPUser)
}

func vipPrint(user VIPUser) {
	fmt.Printf(
		"VIPUser: %s, ID: %s, Age: %d, VIP Level: %d, VIP price: %d\n",
		user.UserInfo.Name,
		user.UserInfo.ID,
		user.UserInfo.Age,
		user.VIPLevel,
		user.Price,
	)
}

func vip2Print(user VIPUser2) {
	fmt.Printf(
		"VIPUser2: %s, ID: %s, Age: %d, VIP Level: %d, VIP price: %d\n",
		user.Name,
		user.ID,
		user.Age,
		user.VIPLevel,
		user.Price,
	)
}

func userPrint(user User) {
	fmt.Printf("User: %s, ID: %s, Age: %d\n", user.Name, user.ID, user.Age)
}
