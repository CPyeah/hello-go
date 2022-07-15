package main

func main() {
	var user = User{id: 1, name: "cp"}
	println(user.getName())
	user.setName("tom")
	println(user.getName())
}

type User struct {
	id   int
	name string
}

func (u User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}
