package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/reflection"
	"github.com/c12s/nebula/pkg/api"
)


func main() {
	s := grpc.NewServer()
	nebula := handlers.NewNebulaGrpcHandler()

	// s := grpc.NewServer(grpc.UnaryInterceptor(handlers.GetAuthInterceptor()))
	api.RegisterMeridianServer(s, nebula)
	reflection.Register(s)

	lis, err := net.Listen("tcp", os.Getenv("LISTEN_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Printf("server listening at %v", lis.Addr())
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	<-shutdown

	s.GracefulStop()
}
