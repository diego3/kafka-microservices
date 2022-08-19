package main

import (
	"github.com/diego3/kafka-microservices/orders-api/application"
	//"github.com/diego3/kafka-microservices/orders-api/event"
)

// go get google.golang.org/grpc
// go get google.golang.org/protobuf
// go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

func init() {

}

func main() {
	var webApp application.WebApp
	webApp.Run()
}
