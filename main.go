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
}
