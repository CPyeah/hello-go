package main

import "fmt"

func main() {
	intType()
	floatType()
	numberConvert()
	charType()
	booleanType()
	stringType()
}

func intType() {
	const (
		i1       = 1
		i2       = 1000
		i3 int8  = 127
		i4 int16 = 128
		i5 uint  = 257
		i6 int   = 0b101
		i7       = 0o77
		i8       = 0xAF
	)
	fmt.Println(i1, i2, i3, i4, i5, i6, i7, i8)
}

func floatType() {
	var (
		f1         = 0.1
		f2 float64 = 0.0001
	)
	fmt.Println(f1, f2)
}

func numberConvert() {
	var n1 float32 = 0.0001
	var n2 float64 = 0.0002
	n3 := float64(n1)
	n4 := float32(n2)
	fmt.Printf("%v, %T; %v, %T; %v, %T; %v, %T \n", n1, n1, n2, n2, n3, n3, n4, n4)
}

func charType() {
	var (
		c1 byte
		c2 = 'c'
		c3 = '鹏'
	)
	fmt.Printf("\"%v, %T; %v, %T; %v, %T \n", c1, c1, c2, c2, c3, c3)

	c4 := 'a'
	c5 := 'A'
	c6 := 'x' + c5 - c4
	fmt.Println(c6, string(c6))
}

func booleanType() {
	var b1 bool = true
	var b2 = false
	b3 := b1 && b2
	fmt.Println(b1, b2, b3)
}

func stringType() {
	s1 := "hello"
	s2 := "世界"
	fmt.Println(s1 + s2)
	fmt.Println(len(s1), len(s2))
}
