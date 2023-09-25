package main

import (
	"context"
	"example/pb"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	NOTIFICATION_SERVICE_PORT = ":8001"
	API_PORT                  = ":8000"
)

func main() {
	conn, err := grpc.Dial(
		NOTIFICATION_SERVICE_PORT,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to notification service: %s", err.Error())
	}

	notificationClient := pb.NewNotificationClient(conn)
	app := echo.New()

	app.POST("/register", func(c echo.Context) error {
		reqBody := RegisterBody{}
		if err := c.Bind(&reqBody); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		// repository success

		// send notification
		mailBody := pb.Mail{
			Email:   reqBody.Email,
			Subject: "register success confirmation",
			Body:    "Welcome to h8. Please activate your account!",
		}
		_, err := notificationClient.SendEmail(context.Background(), &mailBody)
		if err != nil {
			// flaging send mail tidak bisa dipanggil
			fmt.Println("failed to commmunicate with notification service")
		} else {
			fmt.Println("success send activation email")
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success register",
		})
	})

	log.Fatal(app.Start(API_PORT))
}
