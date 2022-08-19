package domain

type Order struct {
	Id     string `json:id`
	Email  string `json:email`
	Amount string `json:amount`
}
