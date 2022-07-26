package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	newUser()
	jsonPackageUser()
}

func jsonPackageUser() {
	var jerry = NewUser("jerry", "12121")
	jerry.Age = 8
	jerry.Job = []string{"programing", "study", "watch TV"}
	jerry.Attr = map[string]string{
		"aa": "bb",
		"cc": "dd",
	}
	jerry.TagTest = "a tag"
	jerry.Ignore = "Ignored"
	fmt.Println(*jerry)

	var data, _ = json.Marshal(jerry)
	fmt.Println(string(data))
	data, _ = json.MarshalIndent(*jerry, "", "\t")
	fmt.Println(string(data))

	var u1 user
	var _ = json.Unmarshal(data, &u1)
	fmt.Println("u1 = ", u1)
}

func newUser() {
	var user = NewUser("tom", "111111")
	fmt.Println(*user)
	println(user.getName())
	user.setPassword("123456")
	fmt.Println(*user)
}
