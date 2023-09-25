package main

import (
	"example/pb"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	PORT = ":8001"
)

func main() {
	mailer := NewMailer(
		"",
		"",
		"",
	)

	srv := grpc.NewServer()
	notificationServer := NotificationServer{
		Mail: mailer,
	}
	pb.RegisterNotificationServer(srv, notificationServer)

	fmt.Println("Starting RPC server at", PORT)
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("failed to start rpc server: %s", err.Error())
	}

	log.Fatal(srv.Serve(listener))
}
