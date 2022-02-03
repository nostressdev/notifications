package main

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

func main() {
	firebsaseCreds := "firebase_creds.json"
	creds := option.WithCredentialsFile(firebsaseCreds)
	app, err := firebase.NewApp(context.Background(), nil, creds)
	if err != nil {
		log.Println(err.Error())
	}
	client, err := app.Messaging(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}
	client.Send(context.Background(), &messaging.Message{
		Data: nil,
	})
}
