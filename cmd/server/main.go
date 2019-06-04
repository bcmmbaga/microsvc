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
	ctx      = context.Background()
	dbConn   = "mongodb://localhost:27017"
	dbName   = "todosdb"
	collName = "todos"
)

func main() {

	if err := runServer(); err != nil {
		log.Fatalf("Error occured: %#v", err.Error())
	}
}

func runServer() error {
	srv := taskservice.NewTaskService(ctx, dbConn, dbName, collName)
	if srv == nil {
		return errors.New("Failed to initialize new task service")
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
