package main

import (
	"encoding/json"
	"fmt"
	"hello-go/tcp"
	"net"
)

func main() {
	var conn, err = net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	fmt.Println("connect success.")
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
		fmt.Println("connection closed.")
	}(conn)

	for {
		fmt.Println("please print a message:")
		var message string
		var _, _ = fmt.Scanf("%s \n", &message)
		if "exit" == message {
			fmt.Println("exit.")
			return
		}
		var messageBody = tcp.NewMessage("tom", message)

		//var b, _ = json.Marshal(&messageBody)
		//var _, err = conn.Write(b)

		var err = json.NewEncoder(conn).Encode(&messageBody)
		if err == nil {
			fmt.Println("send success", messageBody)
		}
	}
}
