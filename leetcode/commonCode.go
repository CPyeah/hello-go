package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	stringFunc()
	stringBuilderFunc()
	intAndByteFunc()
	arrayFunc()
	mapAndStack()
}

//String
func stringFunc() {
	//	### string -> byte
	s := "h"
	b := s[0]
	fmt.Println(b)

	//	### byte -> string
	b = 'h'
	s = string(b)
	fmt.Println(s)

	//	### string -> int
	s = "7"
	i, _ := strconv.Atoi(s)
	fmt.Println(i)

	//	### int -> string
	i = 8
	s = strconv.Itoa(i)
	fmt.Println(s)

	//	### string -> array
	s = "hello"
	arr := strings.Split(s, "")
	fmt.Println(arr)

	// trim
	s = "  hello   "
	s = strings.Trim(s, " ")
	s = strings.TrimSpace(s)
	fmt.Println(s)
}

func stringBuilderFunc() {
	var sb strings.Builder
	sb.WriteString("a")
	sb.WriteString("b")
	sb.WriteString("c")
	s := sb.String()
	fmt.Println(s)
}

func intAndByteFunc() {
	//	### max(a, b)
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	fmt.Println(max(1, 2))
	//	### min(a, b)
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	fmt.Println(min(1, 2))

	//	### max(a, b, c)
	max3 := func(a int, b int, c int) int {
		m := a
		if b > m {
			m = b
		}
		if c > m {
			m = c
		}
		return m
	}
	fmt.Println(max3(1, 2, 3))
	//	### abs(a)
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	fmt.Println(abs(-3))
	//	### isNumber('7')
	s := "3c"
	r0 := unicode.IsDigit(rune(s[0]))
	r1 := unicode.IsDigit(rune(s[1]))
	fmt.Println(r0)
	fmt.Println(r1)

	//	### min and max int32
	maxInt := math.MaxInt32
	minInt := math.MinInt32
	fmt.Println(maxInt)
	fmt.Println(minInt)
}

func arrayFunc() {
	// array -> string
	arr := []string{"he", "llo", " ", "wor", "ld", "!"}
	s := strings.Join(arr, "")
	fmt.Println(s)

	// sort
	intArr := []int{3, 5, 2, 3, 1, 65, 4}
	sort.Slice(intArr, func(i, j int) bool {
		return intArr[i] < intArr[j]
	})
	fmt.Println(intArr)

	// sub array
	arr = []string{"he", "llo", " ", "wor", "ld", "!"}
	arr = arr[:2]
	fmt.Println(arr)

	// 2D array
	arr2 := make([][]int, 3)
	for i := 0; i < len(arr2); i++ {
		arr2[i] = make([]int, 4)
	}
	fmt.Println(arr2)
}

func mapAndStack() {
	// new map
	m := make(map[string]int)
	m["tom"] = 1
	m["jerry"] = 2
	a, ok := m["tom"]
	fmt.Println("tom", a, ok)
	delete(m, "tom")
	a, ok = m["tom"]
	fmt.Println("tom", a, ok)

	// iterate map
	for k, v := range m {
		fmt.Println(k, v)
	}

	// set
	set := make(map[string]bool)
	set["tom"] = true
	set["jerry"] = true
	delete(set, "jerry")
	containsTom := set["tom"]
	containsJerry := set["jerry"]
	fmt.Println(containsTom, containsJerry)

	// Stack
	stack := make([]int, 0)
	// add
	stack = append(stack, 1)
	stack = append(stack, 2)
	// pop
	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	fmt.Println(top) //2
	// peek
	top = stack[len(stack)-1]
	fmt.Println(top) //1
}

func bitOperation() {

}
