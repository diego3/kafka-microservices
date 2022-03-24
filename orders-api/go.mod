module github.com/diego3/kafka-microservices/orders-api

go 1.17

replace github.com/diego3/kafka-microservices/lib => ../lib

require (
	github.com/diego3/kafka-microservices/lib v0.0.0-20220323013629-0e1289ea7808
	github.com/go-stack/stack v1.8.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/klauspost/compress v1.15.1 // indirect
	github.com/segmentio/kafka-go v0.4.30
	github.com/youmark/pkcs8 v0.0.0-20201027041543-1326539a0a0a // indirect
	go.mongodb.org/mongo-driver v1.8.4
	golang.org/x/crypto v0.0.0-20220321153916-2c7772ba3064 // indirect
)
