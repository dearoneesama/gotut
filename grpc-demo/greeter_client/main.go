package main

import (
	"context"
	"log"
	"os"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/models"
)

const address = "localhost:7609"

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Please provide your name as argument")
	}
	name := os.Args[1]

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	client := models.NewGreeterClient(conn)
	resp, err := client.SayHiToServer(context.TODO(), &models.SayHiToServerRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalln(err)
	}

	for _, m := range resp.Greetings {
		fmt.Printf("%v\n", m)
	}
	fmt.Printf("time: %v\nyour luck: %v\n", resp.Time.AsTime(), resp.Luck)
}
