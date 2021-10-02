package main

import (
	"fmt"
	"context"
	"log"
	"math/rand"
	"google.golang.org/protobuf/types/known/timestamppb"
	"grpc-demo/models"
)

// greeterServer is used to implement models.GreeterServer
type greeterServer struct {
	models.UnimplementedGreeterServer
}

var staticGreetings = []string {
	"Hope you have a nice day.",
	"Good morning!",
	"Good luck.",
	"I love you too.",
}

func (s *greeterServer) SayHiToServer(ctx context.Context, req *models.SayHiToServerRequest) (*models.SayHiToServerResponse, error) {
	log.Printf("SayHiToServer received: %v\n", req)
	greetings := [2]string { fmt.Sprintf("Hi, %v!", req.Name) }
	greetings[1] = staticGreetings[rand.Int() % 4]
	return &models.SayHiToServerResponse {
		Greetings: greetings[:],
		Time: timestamppb.Now(),
		Luck: int32(rand.Int() % 101),
	}, nil
}
