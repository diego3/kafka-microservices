package lib

import (
	"fmt"

	"github.com/google/uuid"
)

// Shared lib between microservices
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome! %v", name, uuid.String())
	return message
}
