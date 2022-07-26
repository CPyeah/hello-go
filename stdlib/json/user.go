package main

type user struct {
	name        string
	password    string
	Age         int
	Job         []string
	Attr        map[string]string
	TagTest     string `json:"tag"`
	EmptyValue  string `json:"empty_value,omitempty"`
	EmptyValue1 string
	Ignore      string `json:"-"`
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
