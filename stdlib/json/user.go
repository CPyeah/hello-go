package main

type user struct {
	name     string
	password string
}

func NewUser(name string, password string) *user {
	var user = user{
		name:     name,
		password: password,
	}
	return &user
}

func (u *user) getName() string {
	return u.name
}

func (u *user) setPassword(password string) {
	u.password = password
}
