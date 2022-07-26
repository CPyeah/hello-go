package main

import (
	"encoding/json"
	"fmt"
	"hello-go/tcp"
	"io"
	"net"
)

func main() {
	var listen, err = net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("listener closed.")
	}(listen)

	for {
		fmt.Println("waiting...")
		var conn, err = listen.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("connection closed.")
	}(conn)
	fmt.Println(conn, "connection started.")
	for {
		//var buffer = make([]byte, 4096)
		//var size, err = conn.Read(buffer)
		//fmt.Println(string(buffer[:size]))
		//if err == io.EOF {
		//	fmt.Println("client exited.")
		//	return
		//}
		//if err != nil {
		//	fmt.Println("read message error.")
		//	return
		//}
		var messageBody = tcp.NewEmptyMessage()
		//var _ = json.Unmarshal(buffer[:size], &messageBody)
		//fmt.Println("accept message: ", messageBody)

		var err = json.NewDecoder(conn).Decode(&messageBody)
		if err == io.EOF {
			fmt.Println("client exited.")
			return
		}
		if err != nil {
			fmt.Println("read message error.")
			return
		}
		fmt.Println("accept message:", messageBody)
	}
}
