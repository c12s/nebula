package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc"
	"github.com/c12s/nebula/internal/handlers"
	"github.com/c12s/nebula/pkg/api"
	"github.com/c12s/nebula/internal/storage/etcd"
)


func main() {
	s := grpc.NewServer()

	_, storage := etcd.NewEtcdStorage()
	nebula := handlers.NewNebulaGrpcHandler(storage)

	// s := grpc.NewServer(grpc.UnaryInterceptor(handlers.GetAuthInterceptor()))
	api.RegisterNebulaServer(s, nebula)
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
