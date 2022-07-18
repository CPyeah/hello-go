package main

import (
	"flag"
	"fmt"
	"os"
)

// go run stdlib/command/main.go -v=v0.0.1 -f=101 -isGolang -user cp noFlag
func main() {
	commandArgs()
}

func commandArgs() {
	var count = len(os.Args)
	fmt.Println("received", count, "args")
	// print all args
	for i, arg := range os.Args {
		fmt.Printf("the %v arg is %v \n", i, arg)
	}

	// get flag
	var isGolang = flag.Bool("isGolang", false, "isGolang")
	var version = flag.String("v", "", "version")
	var user = flag.String("user", "", "user")
	var empty = flag.String("empty", "default", "empty")
	flag.Func("f", "", func(s string) error {
		fmt.Println("auto invoke function, args is", s)
		return nil
	})
	flag.Parse()
	fmt.Println("isGolang:", *isGolang, "\nversion:", *version,
		"\nuser:", *user, "\nempty:", *empty)

	// print all noFlag args
	for i, arg := range flag.Args() {
		fmt.Printf("the %v noFlag arg is %v \n", i, arg)
	}
}
