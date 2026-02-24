package main

import (
	"bufio"
	"net"
)

func main() {
	println("Welcome to the API Gateway\n\n TCP server Listing on 127.0.0.1:8080")

	// server listing incoming http request -> TCP

	lis, err := net.Listen("tcp", "127.0.0.1:8080")

	if err != nil {
		panic(err)
	}
	defer func(lis net.Listener) {
		err := lis.Close()
		if err != nil {
			println("ERROR : ", err)
		}
	}(lis)

	for {
		connection, err := lis.Accept()

		if err != nil {
			println("Error -> Accepting connection", err)
			continue
		}

		////println("req :", connection)
		//reader := bufio.NewReader(connection)
		//print(reader)
		//responce := "HTTP/1.1 200 OK\\r\\nContent-Length: 3\\r\\n\\r\\nhii"
		//connection.Write([]byte(responce))
		//
		//connection.Close()

		go handleConn(connection)

	}
}

func handleConn(connection net.Conn) {
	defer func(connection net.Conn) {
		err := connection.Close()
		if err != nil {
			println("Connection close error", err)
		}
	}(connection)

	reader := bufio.NewReader(connection)
	data, err := reader.ReadString('\n')

	if err != nil {
		println("Error occur read string", err)
		return
	}

}
