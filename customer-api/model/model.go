package model

type Customer struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Validation struct {
	Status bool
	Error  string
}

type ErrorResponse struct {
	Message string `json:"message"`
}
