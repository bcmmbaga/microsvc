package main

import (
	"context"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/bcmmbaga/microsvc/pkg/taskservice"

	"google.golang.org/grpc"

	"github.com/bcmmbaga/microsvc/pkg/taskpb"
)

var (
	ctx = context.Background()
	srv = taskservice.NewTaskService(ctx, "todosdb", "todos", "mongo-0.mongo", "mongo-1.mongo")
)

func main() {

	if err := runServer(); err != nil {
		log.Fatalf("Error occured: %#v", err.Error())
	}
}

func runServer() error {
	if srv == nil {
		return errors.New("Failed to initialize new Task service")
	}

	lis, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		return err
	}

	// register taskservice
	s := grpc.NewServer()
	taskpb.RegisterTaskServiceServer(s, srv)

	// gracefull shutdown server
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	go func() {
		// wait for ^C signal
		<-stop

		log.Println("Shutting down gRPC server...")

		s.GracefulStop()

		<-ctx.Done()

	}()

	return s.Serve(lis)
}
