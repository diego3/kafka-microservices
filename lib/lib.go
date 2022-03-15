package lib

import "fmt"

// Shared lib beetwen microservices
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
