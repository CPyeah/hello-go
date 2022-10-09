package main

import (
	"fmt"
	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
)

func main() {
	// the language packages are available under golem/dicts
	// "en" is for english
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	word := lemmatizer.Lemma("revealing")
	fmt.Println(word)
}
