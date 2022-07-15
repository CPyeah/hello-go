package main

import (
	"fmt"
)

func main() {
	year := 2022
	increase(year)
	fmt.Println("year", year, &year)

	increase1(&year)
	fmt.Println("year: ", year, &year)

}

func increase(i int) {
	i++
	fmt.Println("i: ", i, &i)
}

func increase1(i *int) {
	*i++
	fmt.Println("i: ", i, *i, &i)
}
