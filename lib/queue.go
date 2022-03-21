package lib

type Queuer interface {
	Send(message string)
	Receive(message string)
}

type SQS struct {
	Queue Queuer
}
