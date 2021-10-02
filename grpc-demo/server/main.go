package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"grpc-demo/models"
)

const port = ":7609"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	models.RegisterGreeterServer(server, &greeterServer{})
	log.Printf("Listening at %v...", port)
	log.Fatalln(server.Serve(lis))
}
