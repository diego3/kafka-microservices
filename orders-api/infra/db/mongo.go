package db

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Uri     string
	Timeout time.Duration
}

var ctx = context.TODO()

func (r *MongoDB) Connection() {
	// https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.8.4/mongo/options#ClientOptions
	clientOptions := options.Client().ApplyURI(r.Uri)
	clientOptions.SetServerSelectionTimeout(r.Timeout)
	clientOptions.SetMaxPoolSize(1) // Requests to a server will block if this maximum is reached.

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("MongoDB failed to connect:", err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("MongoDB failed to PING:", err)
	}
}
