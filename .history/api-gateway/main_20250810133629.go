package main

import "google.golang.org/grpc"

func main() {
	// create grpc connections
	userConn, err := grpc.Dial("ki")
}