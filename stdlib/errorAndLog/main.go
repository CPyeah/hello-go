package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	errorOperation()

	logOperation()

	// remove file
	var _ = os.Remove(LogFilePath)
}

func logOperation() {
	defer func() {
		recover()
		var e = errors.New("this is a big error")
		ERROR.Fatalln(e)

	}()
	INFO.Println("this is a info")
	var p = errors.New("this is a panic")
	WARN.Panicln(p)

}

func errorOperation() {
	defer func() {
		var e = recover()
		fmt.Println("recover:", e)
	}()
	var lovelyError = errors.New("this is a lovely error")
	var cuteError = fmt.Errorf("this is a cute error")
	fmt.Println(lovelyError, ";", cuteError)

	panic(cuteError)
}
