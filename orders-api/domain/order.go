package domain

type Order struct {
	Email  string `json:email`
	Amount string `json:amount`
}
