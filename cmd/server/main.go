package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/yerrange/go_project_calc/docs"
	"github.com/yerrange/go_project_calc/internal/api"
	pb "github.com/yerrange/go_project_calc/proto"
	"google.golang.org/grpc"
)

func main() {
	// HTTP
	go func() {
		http.HandleFunc("/execute", api.HandleExecute)
		fmt.Println("HTTP server running on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// gRPC
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculatorServer(grpcServer, &api.GrpcServer{})

	fmt.Println("gRPC server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
