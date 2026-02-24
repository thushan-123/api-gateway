package main

import "net"

func main() {
	println("Welcome to the API Gateway")

	// server listing incoming http request -> TCP

	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		panic(err)
	}
	defer lis.Close() // close program after exit

	for {
		connection, err := lis.Accept()

		if err != nil {
			println("Error -> Accepting connection", err)
			continue
		}

		println("req :", connection)
		responce := "hii"
		connection.Write([]byte(responce))

		connection.Close()

	}
}
