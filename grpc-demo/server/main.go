package main

import (
	"context"
	"grpc-demo/models"
	"log"
	"net"
	"net/http"
	"sync"

	gateway "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

const grpcPort = ":7609"
const httpPort = ":7610"

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func () {
		defer wg.Done()
		lis, err := net.Listen("tcp", grpcPort)
		if err != nil {
			log.Fatalln(err)
		}
		server := grpc.NewServer()
		models.RegisterGreeterServer(server, &greeterServer{})
		log.Printf("Grpc server listening at %v...", grpcPort)
		log.Fatalln(server.Serve(lis))
	}()

	go func () {
		defer wg.Done()
		mux := gateway.NewServeMux()
		err := models.RegisterGreeterHandlerFromEndpoint(
			context.TODO(), mux, "localhost" + grpcPort,
			[]grpc.DialOption{grpc.WithInsecure()},
		)
		if err != nil {
			log.Fatalln(err)
		}
		log.Printf("Http server listening at %v", httpPort)
		log.Fatal(http.ListenAndServe(httpPort, mux))
	}()

	wg.Wait()
}
