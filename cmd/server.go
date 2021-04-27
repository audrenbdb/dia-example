package main

import (
	"fmt"
	"log"
	"net"

	"github.com/audrenbdb/dia"
	"github.com/audrenbdb/dia/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = 50052

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	diaServer := dia.NewServer(dia.NewService(dia.NewInMemoryRepo()))
	reflection.Register(grpcServer)
	pb.RegisterSolutionsServer(grpcServer, diaServer)
	grpcServer.Serve(lis)
}
