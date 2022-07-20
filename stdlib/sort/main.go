package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name string
	age  int
}

func main() {
	var intSlice = []int{3, 5, 2, 5, 1, 6, 3, 4, 7}
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	var stringSlice = []string{"A", "a", "AB", "Ab", "AB", "ABA"}
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)

	var personSlice = []Person{{"tom", 8}, {"jerry", 6}, {"chengpeng", 18}}
	sort.Slice(personSlice, func(i, j int) bool {
		return personSlice[i].age < personSlice[j].age
	})
	fmt.Println(personSlice)

	var i = sort.Search(len(personSlice), func(i int) bool {
		if personSlice[i].age >= 18 {
			return true
		}
		return false
	})
	fmt.Println(personSlice[i])

}
