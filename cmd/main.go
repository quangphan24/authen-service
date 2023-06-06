package main

import (
	grpc2 "authen-service/grpc"
	"authen-service/proto/authen"
	"google.golang.org/grpc"
	"log"
	"net"
)

const VERSION = "1.0.0"

// @title Example API
// @version 1.0

// @BasePath /api
// @schemes http http

// @securityDefinitions.apikey AuthToken
// @in header
// @name Authorization

// @description Transaction API.
func main() {
	port := ":8088"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("grpc server started on: [::]:8088")
	s := grpc.NewServer()
	authen.RegisterAuthenServiceServer(s, &grpc2.ServerGRPC{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
