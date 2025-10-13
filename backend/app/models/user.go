package models

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
	City string `json:"city"`
	State string `json:"state"`
	Zip string `json:"zip"`
}
