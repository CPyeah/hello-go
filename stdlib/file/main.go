package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fileOperation()

	fileReadAndWriteSmallFile()

	fileReadAndWriteWithBuffer()

}

func fileReadAndWriteWithBuffer() {
	var path = "tempTest.txt"

	// open
	var file, _ = os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)

	// new writer
	var bufferWriter = bufio.NewWriter(file)

	// write
	var _, _ = bufferWriter.Write([]byte("this is first line\n"))
	var _, _ = bufferWriter.Write([]byte("this is second line\n"))
	var _, _ = bufferWriter.Write([]byte("this is third line\n"))
	var _, _ = bufferWriter.Write([]byte("this is fourth line\n"))

	var _ = bufferWriter.Flush()

	var _ = file.Close()

	// open read only file
	file, _ = os.OpenFile(path, os.O_RDONLY, 0666)

	// read line
	var bufferReader = bufio.NewReader(file)
	for {
		var r, _, _ = bufferReader.ReadLine()
		if r == nil {
			break
		}
		println(string(r))
	}

	var _ = file.Close()

	// remove file
	var _ = os.Remove(path)

}

func fileReadAndWriteSmallFile() {
	var path = "tempTest.txt"
	// open a file
	var file, _ = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0777)

	// defer close file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// write
	_ = os.WriteFile(path, []byte("hello word"), 0777)

	// read
	var data, _ = os.ReadFile(path)
	fmt.Println(string(data))

	// delete
	_ = os.Remove(path)
}

func fileOperation() {
	var path = "dir/hello.txt"
	exist(path)
	createDirAndFile(path)
	exist(path)
	path = rename(path)
	exist(path)
	readDir()
	deleteFile(path)
	exist(path)
}

func readDir() {
	var dirs, err = os.ReadDir("dir")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		fmt.Println("sub dir/file name is", dir.Name())
	}
}

func deleteFile(path string) {
	var index = strings.LastIndex(path, "/")
	var dir = path[:index]
	var err = os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}

func rename(path string) string {
	var newPath = "dir/helloWorld.txt"
	var err = os.Rename(path, newPath)
	if err != nil {
		panic(err)
	}
	return newPath
}

func createDirAndFile(path string) {
	var index = strings.LastIndex(path, "/")
	var dir = path[:index]
	var err = os.MkdirAll(dir, 0777)
	_, err = os.Create(path)
	if err != nil {
		panic(err)
	}
}

func exist(path string) bool {
	var s, err = os.Stat(path)
	if err != nil {
		fmt.Println(path, "not exist.")
		return os.IsExist(err)
	}
	fmt.Println(s.Name(), s.Mode(), s.ModTime(), s.IsDir())
	return true
}
