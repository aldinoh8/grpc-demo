package main

import (
	"context"
	"example/pb"
	"fmt"

	"google.golang.org/protobuf/types/known/emptypb"
)

type NotificationServer struct {
	pb.UnimplementedNotificationServer
	Mail Mailer
}

func (server NotificationServer) SendEmail(ctx context.Context, param *pb.Mail) (*pb.MailResponse, error) {
	// logic untuk kirim email
	// di
	// mailer
	err := server.Mail.GenerateMail(param.Email, param.Subject, param.Body)
	if err != nil {
		return &pb.MailResponse{}, nil
	}

	return &pb.MailResponse{Status: "SENT"}, nil
}

func (server NotificationServer) GetEmail(ctx context.Context, param *emptypb.Empty) (*pb.ListMail, error) {
	fmt.Println("notif server: GET email triggered")
	return &pb.ListMail{}, nil
}
