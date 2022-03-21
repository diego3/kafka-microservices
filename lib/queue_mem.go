package lib

type QueueMemory struct {
	Messages    []string
	Subscribers []string
}

// Sends a message to the queue
func (q *QueueMemory) Send(message string) {
	q.Messages = append(q.Messages, message)
}
