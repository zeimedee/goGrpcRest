package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/zeimedee/goGrpcRest/interceptors"
	"github.com/zeimedee/goGrpcRest/server"

	student "github.com/zeimedee/goGrpcRest/api/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func main() {
	addr := ":8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln("Failed to serve:..", err)
	}

	fmt.Println("Server is running on port:", addr)

	s := grpc.NewServer(grpc.UnaryInterceptor(interceptors.UnaryServerInterceptor()))

	student.RegisterStudentServiceServer(s, &server.StudentServer{})

	go func() {
		log.Fatal(s.Serve(lis))

	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("Failed to dial:..", err)
	}

	gmux := runtime.NewServeMux()

	err = student.RegisterStudentServiceHandler(context.Background(), gmux, conn)
	if err != nil {
		log.Fatalln("failed tp register service", err)
	}
	gServer := &http.Server{
		Addr:    ":8090",
		Handler: gmux,
	}

	log.Println("servers is listening on localhost:8090")
	log.Fatal(gServer.ListenAndServe())
}
