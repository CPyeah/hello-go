package main

import "fmt"

func main() {
	newUser()
	newUserByPointer()

	newSubStruct()
}

func newSubStruct() {
	var user = User{Name: "cp"}

	var address = Address{
		user: User{Name: "tom"},
		addr: "china",
	}
	fmt.Println(user)
	fmt.Println(address)

	var email = Email{
		user:  &user,
		email: "cp@gm.com",
	}
	fmt.Println(email.user.Name)
	user.Name = "cp1"
	fmt.Println(email.user.Name)
}

func newUser() {
	var u1 = User{
		Id:   1,
		Name: "tom",
	}
	fmt.Println(u1)
}

func newUserByPointer() {
	var u = &User{Name: "jerry"}
	fmt.Println(u.Name, (*u).Name)
}

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

type Address struct {
	user User
	addr string
}

type Email struct {
	user  *User
	email string
}
