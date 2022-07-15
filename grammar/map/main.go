package main

import "fmt"

func main() {
	createMap()

	findDeleteRange()
}

func findDeleteRange() {
	var m1 = make(map[string]string)
	m1["morning"] = "eat breakfast"
	m1["noon"] = "have lunch"

	var v, ok = m1["noon"]
	fmt.Println(v, ok)

	delete(m1, "noon")

	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

func createMap() {
	var m1 = make(map[string]string)
	m1["morning"] = "eat breakfast"
	m1["noon"] = "have lunch"

	var m2 = map[string]string{
		"evening": "get dinner",
	}

	fmt.Println(m1, m2)

	fmt.Println(m1["noom"], m2["evening"])

}
